package service

import (
	"bytes"
	"html/template"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/adapter/config"
	"gopkg.in/gomail.v2"
)

type SmtpService struct {
}

// NewUserService creates a new user service instance
func NewSmtpService() *SmtpService {
	return &SmtpService{}
}

func (s *SmtpService) SendEmail(ctx echo.Context, email, template, subject string, isHtml bool) error {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "administrator@otoo.gr")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", template)
	d := gomail.NewDialer(config.SMTP.Host, 587, config.SMTP.User, config.SMTP.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}

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
	tmpl, err := template.ParseFiles("assets/template/forgot_password_template.html")
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

	s.SendEmail(ctx, email, body, subject, true)
	return nil
}
func (s *SmtpService) SendContactEmail(ctx echo.Context) error {
	return nil
}
