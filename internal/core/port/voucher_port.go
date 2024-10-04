package port

import (
	"github.com/labstack/echo/v4"
	domain "github.com/stelgkio/otoo/internal/core/domain/courier"
)

// VoucherRepository 	defines the methods for interacting with the Voucher repository
type VoucherRepository interface {
	// CreateVoucher inserts a new Voucher into the database
	CreateVoucher(ctx echo.Context, Voucher *domain.Voucher, projectID string) (*domain.Voucher, error)
	// // GetVoucherByID selects a Voucher by id
	GetVoucherByVoucherID(ctx echo.Context, voucherID string) (*domain.Voucher, error)
	// FindVoucherByProjectIDAsync finds a Voucher by projectID
	FindVoucherByProjectID(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus) ([]*domain.Voucher, error)
	// 	GetAllVouchers() ([]*domain.Voucher, error)
	GetAllVouchers(ctx echo.Context, projectID string) ([]*domain.Voucher, error)
	// // UpdateVoucher updates a Voucher
	UpdateVoucher(ctx echo.Context, user *domain.Voucher, projectID string, voucherID string) (*domain.Voucher, error)
	// DeleteUser(ctx context.Context, id uint64) error
	DeleteVouchersByID(ctx echo.Context, voucherID string) error
	// GetVoucherCount retrieves the count of Vouchers for a given project ID
	GetVoucherCount(projectID string, voucherStatus domain.VoucherStatus) (int64, error)
}

// VoucherService 	defines the methods for interacting with the Voucher repository
type VoucherService interface {
	// CreateVoucher inserts a new Voucher into the database
	CreateVoucher(ctx echo.Context, Voucher *domain.Voucher, projectID string) (*domain.Voucher, error)
	// // GetVoucherByID selects a Voucher by id
	GetVoucherByVoucherID(ctx echo.Context, voucherID string) (*domain.Voucher, error)
	// FindVoucherByProjectIDAsync finds a Voucher by projectID
	FindVoucherByProjectIDAsync(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus) ([]*domain.Voucher, error)
	// 	GetAllVouchers() ([]*domain.Voucher, error)
	GetAllVouchers(ctx echo.Context, projectID string) ([]*domain.Voucher, error)
	// // UpdateVoucher updates a Voucher
	UpdateVoucher(ctx echo.Context, user *domain.Voucher, projectID string, voucherID string) (*domain.Voucher, error)
	// // DeleteUser deletes a user
	// DeleteUser(ctx context.Context, id uint64) error
	DeleteVouchersByID(ctx echo.Context, voucherID string) error
	/// GetVoucherCount retrieves the count of Vouchers for a given project ID
	GetVoucherCountAsync(projectID string, voucherStatus domain.VoucherStatus) (int64, error)
}
