package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/adapter/config"
	"github.com/stelgkio/otoo/internal/core/domain"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"gopkg.in/gomail.v2"
)

// SmtpService represents the SMTP service
type SmtpService struct {
}

// NewSmtpService creates a new instance of SmtpService
func NewSmtpService() *SmtpService {
	return &SmtpService{}
}

// SendEmail sends an email to the user
func (s *SmtpService) SendEmail(ctx echo.Context, sender, email, template, subject string, isHtml bool) error {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", email, "support@konektorx.com")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", template)
	d := gomail.NewDialer(config.SMTP.Host, 587, config.SMTP.User, config.SMTP.Password)
	if err := d.DialAndSend(m); err != nil {
		slog.Error("error DialAndSend", "error", err)
		return err
	}
	return nil
}

// SendForgetPasswordEmail sends a forget password email to the user
func (s *SmtpService) SendForgetPasswordEmail(ctx echo.Context, email, firstName, lastName, resetPasswordLink string) error {
	type ForgotPasswordData struct {
		FirstName         string
		LastName          string
		ResetPasswordLink string
	}
	emailData := ForgotPasswordData{
		FirstName:         firstName,
		LastName:          lastName,
		ResetPasswordLink: resetPasswordLink,
	}

	// Load and parse the HTML template
	tmpl, err := template.ParseFiles(fmt.Sprintf("%sassets/template/forgot_password_template.html", os.Getenv("SERVER_TEMPLATE_PATH")))
	if err != nil {
		slog.Error("Error", "Error loading template: %v", err)
	}

	// Create a buffer to store the parsed template
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, emailData); err != nil {
		slog.Error("Error", "Error executing template: %v", err)
	}

	// Convert the buffer to a string
	body := tpl.String()

	subject := "KonektorX Reset Your Password"

	go s.SendEmail(ctx, "administrator@konektorx.com", email, body, subject, true)
	return nil
}

// SendContactEmail sends a forget password email to the user
func (s *SmtpService) SendContactEmail(ctx echo.Context, req *domain.ContactRequest) error {

	// Load and parse the HTML template
	tmpl, err := template.ParseFiles(fmt.Sprintf("%sassets/template/contact_verification.html", os.Getenv("SERVER_TEMPLATE_PATH")))
	if err != nil {
		slog.Error("Error", "Error loading template: %v", err)
	}

	// Create a buffer to store the parsed template
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, req); err != nil {
		slog.Error("Error", "Error executing template: %v", err)
	}

	// Convert the buffer to a string
	body := tpl.String()

	subject := "KonektorX: Your Message Has Been Received"

	go s.SendEmail(ctx, "hello@konektorx.com", req.Email, body, subject, true)
	return nil

}

type WeeklyData struct {
	FullName        string
	StartDate       time.Time
	EndDate         time.Time
	TotalOrders     int64
	ActiveOrders    int64
	TotalRevenue    float64
	ActiveOrderRate float64
}

// SendWeeklyBalanceEmail sends a forget password email to the user
func (s *SmtpService) SendWeeklyBalanceEmail(ctx echo.Context, req *w.WeeklyAnalytics, email, fullname string) error {
	// Assume you have a struct that holds the email content.

	emailData := WeeklyData{
		FullName:        fullname,
		StartDate:       req.StartDate,
		EndDate:         req.EndDate,
		TotalOrders:     req.TotalOrders,
		ActiveOrders:    req.ActiveOrders,
		TotalRevenue:    req.TotalRevenue,
		ActiveOrderRate: req.ActiveOrderRate,
	}
	// Load and parse the HTML template
	tmpl, err := template.ParseFiles(fmt.Sprintf("%sassets/template/weekly_balance_template.html", os.Getenv("SERVER_TEMPLATE_PATH")))
	if err != nil {
		slog.Error("Error", "Error loading template: %v", err)
	}

	// Create a buffer to store the parsed template
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, emailData); err != nil {
		slog.Error("Error", "Error executing template: %v", err)
	}

	// Convert the buffer to a string
	body := tpl.String()

	subject := "KonektorX: Your Message Has Been Received"

	go s.SendEmail(ctx, "hello@konektorx.com", email, body, subject, true)
	return nil

}

// SendMonthlyOrdersEmail sends a forget password email to the user
func (s *SmtpService) SendMonthlyOrdersEmail(ctx echo.Context, req *w.MonthlyOrderCountAnalytics, email, fullname string) error {

	// Load and parse the HTML template
	tmpl, err := template.ParseFiles(fmt.Sprintf("%sassets/template/contact_verification.html", os.Getenv("SERVER_TEMPLATE_PATH")))
	if err != nil {
		slog.Error("Error", "Error loading template: %v", err)
	}

	// Create a buffer to store the parsed template
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, req); err != nil {
		slog.Error("Error", "Error executing template: %v", err)
	}

	// Convert the buffer to a string
	body := tpl.String()

	subject := "KonektorX: Your Message Has Been Received"

	go s.SendEmail(ctx, "hello@konektorx.com", email, body, subject, true)
	return nil

}
