package service

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
)

// PaymentService defines the methods for interacting with the Payment service
type PaymentService struct {
	repo port.PaymentRepository
}

// NewPaymentService creates a new Payment service instance
func NewPaymentService(repo port.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo,
	}
}

// CreatePayment inserts a new Payment
func (ns *PaymentService) CreatePayment(ctx echo.Context, data *domain.Payment) error {
	err := ns.repo.CreatePayment(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePayment update a new Payment
func (ns *PaymentService) UpdatePayment(ctx echo.Context, data *domain.Payment) error {
	return nil
}

// DeletePayment delete a  Payment
func (ns *PaymentService) DeletePayment(ctx echo.Context, projectID, PaymentID string) error {
	err := ns.repo.DeletePayment(ctx, projectID, PaymentID)
	if err != nil {
		return err
	}
	return nil
}

// FindPayment find Payment not read
func (ns *PaymentService) FindPayment(projectID string, size, page int, sort, direction string, filters bool) ([]*domain.Payment, error) {
	Payment, err := ns.repo.FindPayment(projectID, size, page, sort, direction, filters)
	if err != nil {
		return nil, err
	}
	return Payment, nil
}

// GetPaymentCountAsync count of payments
func (ns *PaymentService) GetPaymentCountAsync(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error) {
	webhookCount, err := ns.repo.PaymentCount(projectID)
	if err != nil {
		errors <- err
	} else {
		results <- webhookCount
	}
}

// FindPaymentByProjectIDAsync return list of payment
func (ns *PaymentService) FindPaymentByProjectIDAsync(ctx echo.Context, projectID string, size, page int, sort, direction string, results chan<- []*domain.Payment, errors chan<- error) {
	payments, err := ns.repo.FindPayment(projectID, size, page, sort, direction, true)
	if err != nil {
		errors <- err
	} else {
		results <- payments
	}
}
