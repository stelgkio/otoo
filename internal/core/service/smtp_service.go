package service

import (
	"bytes"
	"html/template"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/adapter/config"
	"github.com/stelgkio/otoo/internal/core/domain"
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
	m.SetHeader("To", email, "support@otoo.gr")
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
	tmpl, err := template.ParseFiles("/project/otoo/build/assets/template/forgot_password_template.html")
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

	subject := "Otoo Reset Your Password"

	go s.SendEmail(ctx, "administrator@otoo.gr", email, body, subject, true)
	return nil
}
func (s *SmtpService) SendContactEmail(ctx echo.Context, req *domain.ContactRequest) error {

	// Load and parse the HTML template
	tmpl, err := template.ParseFiles("assets/template/contact_verification.html")
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

	subject := "Otoo: Your Message Has Been Received"

	go s.SendEmail(ctx, "hello@otoo.gr", req.Email, body, subject, true)
	return nil

}
