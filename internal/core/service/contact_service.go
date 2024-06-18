package service

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
)

type ContactService struct {
	repo port.ContactRepository
	smtp port.SmtpService
}

// NewUserService creates a new user service instance
func NewContactService(repo port.ContactRepository, smtp port.SmtpService) *ContactService {
	return &ContactService{
		repo,
		smtp,
	}
}

func (c *ContactService) InsertContact(ctx echo.Context, req *domain.ContactRequest) error {

	contact, err := domain.NewContact(req)

	if err != nil {
		return errors.New("contact is not created")
	}

	if err := c.repo.InsertContact(ctx, contact); err != nil {
		return errors.New("contact is not created")
	}
	c.smtp.SendEmail(ctx)
	return nil
}
