package port

import (
	"github.com/labstack/echo/v4"
	d "github.com/stelgkio/otoo/internal/core/domain"
	domain "github.com/stelgkio/otoo/internal/core/domain/courier"
	o "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	woo "github.com/stelgkio/woocommerce"
)

// VoucherRepository 	defines the methods for interacting with the Voucher repository
type VoucherRepository interface {
	// CreateVoucher inserts a new Voucher into the database
	CreateVoucher(ctx echo.Context, Voucher *domain.Voucher, projectID string) (*domain.Voucher, error)
	// // GetVoucherByID selects a Voucher by id
	GetVoucherByVoucherID(ctx echo.Context, voucherID string) (*domain.Voucher, error)
	/// GetVoucherByOrderIDAndProjectID retrieves a Voucher by orderID and projectID
	GetVoucherByOrderIDAndProjectID(ctx echo.Context, orderID int64, projectID string) (*domain.Voucher, error)
	// FindVoucherByProjectIDAsync finds a Voucher by projectID
	FindVoucherByProjectID(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus) ([]*domain.Voucher, error)
	// FindVoucherByProjecAndCourierProvidertID finds a Voucher by projectID
	FindVoucherByProjecAndCourierProvidertID(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus, provider string) ([]*domain.Voucher, error)
	// 	GetAllVouchers() ([]*domain.Voucher, error)
	GetAllVouchers(ctx echo.Context, projectID string) ([]*domain.Voucher, error)
	// // UpdateVoucher updates a Voucher
	UpdateVoucher(ctx echo.Context, voucher *domain.Voucher, projectID string, voucherID string, orderID int64) (*domain.Voucher, error)
	// UpdateVoucherNewDetails updates a Voucher
	UpdateVoucherNewDetails(ctx echo.Context, voucher *domain.Voucher, projectID string) (*domain.Voucher, error)
	// DeleteUser(ctx context.Context, id uint64) error
	DeleteVouchersByID(ctx echo.Context, voucherID string) error
	// GetVoucherCount retrieves the count of Vouchers for a given project ID
	GetVoucherCount(projectID string, voucherStatus domain.VoucherStatus) (int64, error)
	GetVoucherCountByProvider(projectID string, voucherStatus domain.VoucherStatus, provider string) (int64, error)
}

// VoucherService 	defines the methods for interacting with the Voucher repository
type VoucherService interface {
	// CreateVoucher inserts a new Voucher into the database
	CreateVoucher(ctx echo.Context, OrderRecord *o.OrderRecord, projectID string) (*domain.Voucher, error)
	CreateHermesVoucher(ctx echo.Context, voucher *domain.HermesVoucerRequest, projectID string) (*domain.Voucher, error)
	// // GetVoucherByID selects a Voucher by id
	GetVoucherByVoucherID(ctx echo.Context, voucherID string) (*domain.Voucher, error)
	// FindVoucherByProjectIDAsync finds a Voucher by projectID
	FindVoucherByProjectIDAsync(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus, results chan<- []*domain.Voucher, errors chan<- error)
	//FindVoucherByProjectIDAndCourierProviderAsyn finds a Voucher by projectID and courier provider
	FindVoucherByProjectIDAndCourierProviderAsync(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus, provider string, results chan<- []*domain.Voucher, errors chan<- error)
	// 	GetAllVouchers() ([]*domain.Voucher, error)
	GetAllVouchers(ctx echo.Context, projectID string) ([]*domain.Voucher, error)
	// // UpdateVoucher updates a Voucher
	UpdateVoucher(ctx echo.Context, voucher *o.OrderRecord, projectID string) (*domain.Voucher, error)
	// UpdateVoucherWithDetails updates a Voucher with details
	UpdateVoucherNewDetails(ctx echo.Context, voucher *domain.Voucher, projectID string, courier4u *d.Courier4uExtension, redcourier *d.RedCourierExtension) (*domain.Voucher, error)
	// // DeleteUser deletes a user
	// DeleteUser(ctx context.Context, id uint64) error
	DeleteVouchersByID(ctx echo.Context, voucherID string) error
	/// GetVoucherCount retrieves the count of Vouchers for a given project ID
	GetVoucherCountAsync(projectID string, voucherStatus domain.VoucherStatus, results chan<- int64, errors chan<- error)
	GetVoucherCountByProviderAsync(projectID string, voucherStatus domain.VoucherStatus, provider string, results chan<- int64, errors chan<- error)
	DeleteVouchersByOrderIdandProjectID(ctx echo.Context, projectID string, orderID int64) error
	GetVoucherByOrderIDAndProjectID(ctx echo.Context, orderID int64, projectID string) (*domain.Voucher, error)
	UpdateVoucherTracking(ctx echo.Context, vch *domain.Voucher, projectID string, client *woo.Client) (*domain.Voucher, error)
}
