package service

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

type SmtpService struct {
}

// NewUserService creates a new user service instance
func NewSmtpService() *SmtpService {
	return &SmtpService{}
}

func (s *SmtpService) SendEmail(ctx echo.Context) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "administrator@otoo.gr")
	m.SetHeader("To", "gkiostyl13@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "This is the plain text body of the email.")
	d := gomail.NewDialer("mail.otoo.gr", 587, "administrator@otoo.gr", "Liwizt&R1Gg&oUWp")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}
