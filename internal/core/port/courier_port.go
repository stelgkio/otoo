package port

import (
	"github.com/labstack/echo/v4"
	domain "github.com/stelgkio/otoo/internal/core/domain"
	domain_courier "github.com/stelgkio/otoo/internal/core/domain/courier"
)

// HermesService 	defines the methods for interacting with the Voucher repository
type HermesService interface {
	// CreateVoucher inserts a new Voucher into the database
	PrintVoucher(ctx echo.Context, courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension, voucherId string, projectID, printType string) ([]byte, error)
	// // GetVoucherByID selects a Voucher by id
	CreateVoucher(ctx echo.Context, courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension, hermesVoucerRequest *domain_courier.HermesVoucerRequest, projectID string) (*domain_courier.VoucherResponse, error)
}
