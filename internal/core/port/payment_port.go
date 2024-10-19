package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// PaymentRepository interface
type PaymentRepository interface {
	CreatePayment(ctx echo.Context, data *domain.Payment) error
	UpdatePayment(ctx echo.Context, data *domain.Payment) error
	DeletePayment(ctx echo.Context, projectID, PaymentID string) error
	FindPayment(projectID string, size, page int, sort, direction string, filters bool) ([]*domain.Payment, error)
	PaymentCount(projectID string) (int64, error)
}

// PaymentService interface
type PaymentService interface {
	CreatePayment(ctx echo.Context, data *domain.Payment) error
	UpdatePayment(ctx echo.Context, data *domain.Payment) error
	DeletePayment(ctx echo.Context, projectID, PaymentID string) error
	FindPayment(projectID string, size, page int, sort, direction string, filters bool) ([]*domain.Payment, error)
	FindPaymentByProjectIDAsync(ctx echo.Context, projectID string, size, page int, sort, direction string, results chan<- []*domain.Payment, errors chan<- error)
	GetPaymentCountAsync(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error)
}
