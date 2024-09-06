package service

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
)

// NotificationService defines the methods for interacting with the Notification service
type NotificationService struct {
	repo port.NotificationRepository
	smtp port.SmtpService
}

// NewNotificationService creates a new notification service instance
func NewNotificationService(repo port.NotificationRepository, smtp port.SmtpService) *NotificationService {
	return &NotificationService{
		repo,
		smtp,
	}
}

// CreateNotification inserts a new notification
func (ns *NotificationService) CreateNotification(ctx echo.Context, data *domain.Notification) error {
	return nil
}

// UpdateNotification update a new notification
func (ns *NotificationService) UpdateNotification(ctx echo.Context, data *domain.Notification) error {
	return nil
}

// FindNotification find notification not read
func (ns *NotificationService) FindNotification(projectID string, size, page int, sort, direction string) ([]*domain.Notification, error) {
	return nil, nil
}
