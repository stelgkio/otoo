package service

import (
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

func (s *SmtpService) SendEmail(ctx echo.Context) error {
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "administrator@otoo.gr")
	m.SetHeader("To", "gkiostyl13@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "This is the plain text body of the email.")
	d := gomail.NewDialer(config.SMTP.Host, 587, config.SMTP.User, config.SMTP.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}
