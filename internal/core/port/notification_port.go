package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// NotificationRepository interface
type NotificationRepository interface {
	CreateNotification(ctx echo.Context, data *domain.Notification) error
	UpdateNotification(ctx echo.Context, data *domain.Notification) error
	FindNotification(projectID string, size, page int, sort, direction string) ([]*domain.Notification, error)
}

// NotificationService interface
type NotificationService interface {
	CreateNotification(ctx echo.Context, data *domain.Notification) error
	UpdateNotification(ctx echo.Context, data *domain.Notification) error
	FindNotification(projectID string, size, page int, sort, direction string) ([]*domain.Notification, error)
}
