package service

import (
	"github.com/labstack/echo/v4"
	domain "github.com/stelgkio/otoo/internal/core/domain/courier"
	"github.com/stelgkio/otoo/internal/core/port"
)

// VoucherService defines the methods for interacting with the Voucher service
type VoucherService struct {
	repo port.VoucherRepository
}

// NewVoucherService creates a new voucher service instance
func NewVoucherService(repo port.VoucherRepository) *VoucherService {
	return &VoucherService{
		repo,
	}
}

// CreateVoucher inserts a new Voucher into the database
func (vs *VoucherService) CreateVoucher(ctx echo.Context, voucher *domain.Voucher, projectID string) (*domain.Voucher, error) {
	return vs.repo.CreateVoucher(ctx, voucher, projectID)
}

// GetVoucherByVoucherID retrieves a Voucher by its ID
func (vs *VoucherService) GetVoucherByVoucherID(ctx echo.Context, voucherID string) (*domain.Voucher, error) {
	return vs.repo.GetVoucherByVoucherID(ctx, voucherID)
}

// FindVoucherByProjectID retrieves vouchers by project ID
func (vs *VoucherService) FindVoucherByProjectID(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus) ([]*domain.Voucher, error) {
	return vs.repo.FindVoucherByProjectID(projectID, size, page, sort, direction, voucherStatus)
}

// GetAllVouchers retrieves all vouchers for a project
func (vs *VoucherService) GetAllVouchers(ctx echo.Context, projectID string) ([]*domain.Voucher, error) {
	return vs.repo.GetAllVouchers(ctx, projectID)
}

// UpdateVoucher updates a Voucher
func (vs *VoucherService) UpdateVoucher(ctx echo.Context, voucher *domain.Voucher, projectID string, voucherID string) (*domain.Voucher, error) {
	return vs.repo.UpdateVoucher(ctx, voucher, projectID, voucherID)
}

// DeleteVouchersByID deletes a voucher by its ID
func (vs *VoucherService) DeleteVouchersByID(ctx echo.Context, voucherID string) error {
	return vs.repo.DeleteVouchersByID(ctx, voucherID)
}

// GetVoucherCount retrieves the count of Vouchers for a given project ID
func (vs *VoucherService) GetVoucherCount(projectID string, voucherStatus domain.VoucherStatus) (int64, error) {
	return vs.repo.GetVoucherCount(projectID, voucherStatus)
}

// GetVoucherCountAsync retrieves the count of vouchers asynchronously
func (vs *VoucherService) GetVoucherCountAsync(projectID string, voucherStatus domain.VoucherStatus, results chan<- int64, errors chan<- error) {
	count, err := vs.repo.GetVoucherCount(projectID, voucherStatus)
	if err != nil {
		errors <- err
	} else {
		results <- count
	}
}

// FindVoucherByProjectIDAsync retrieves vouchers for a given project ID asynchronously
func (vs *VoucherService) FindVoucherByProjectIDAsync(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus, results chan<- []*domain.Voucher, errors chan<- error) {
	vouchers, err := vs.repo.FindVoucherByProjectID(projectID, size, page, sort, direction, voucherStatus)
	if err != nil {
		errors <- err
	} else {
		results <- vouchers
	}
}
