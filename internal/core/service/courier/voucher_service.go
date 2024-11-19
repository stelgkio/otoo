package service

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	d "github.com/stelgkio/otoo/internal/core/domain"
	domain "github.com/stelgkio/otoo/internal/core/domain/courier"
	o "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/woocommerce"
)

// VoucherService defines the methods for interacting with the Voucher service
type VoucherService struct {
	repo            port.VoucherRepository
	hermesSvc       port.HermesService
	woocommerceRepo port.WoocommerceRepository
}

// NewVoucherService creates a new voucher service instance
func NewVoucherService(repo port.VoucherRepository, hermesSvc port.HermesService, woocommerceRepo port.WoocommerceRepository) *VoucherService {
	return &VoucherService{
		repo,
		hermesSvc,
		woocommerceRepo,
	}
}

// SplitFullName splits a full name into first and last names
func SplitFullName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	if len(parts) > 0 {
		firstName = parts[0]
	}
	if len(parts) > 1 {
		// Join the rest of the parts as the last name
		lastName = strings.Join(parts[1:], " ")
	}
	return
}

// CreateVoucher inserts a new Voucher into the database
func (vs *VoucherService) CreateVoucher(ctx echo.Context, OrderRecord *o.OrderRecord, projectID string) (*domain.Voucher, error) {
	voucher, err := vs.repo.GetVoucherByOrderIDAndProjectID(ctx, OrderRecord.Order.ID, projectID)
	if err != nil {
		return nil, err
	}

	if voucher == nil {
		newvoucher := domain.NewVoucher(projectID, OrderRecord.Order.ShippingTotal, OrderRecord.Order.CustomerNote, OrderRecord.Order.Shipping, OrderRecord.Order.Billing, OrderRecord.Order.ID, OrderRecord.Order.LineItems, OrderRecord.Order.PaymentMethod, OrderRecord.Order.Total)
		return vs.repo.CreateVoucher(ctx, newvoucher, projectID)
	}
	// Update the voucher status based on the order status
	switch OrderRecord.Order.Status {
	case "completed":
		if !voucher.IsPrinted {
			voucher.UpdateVoucherStatus(domain.VoucherStatusCompleted)
		}
	case "processing":
		if voucher.Status == domain.VoucherStatusNew {
			voucher.UpdateVoucherStatus(domain.VoucherStatusNew)
		} else {
			//TODO: to update a voucher we need to check in what state we have it in the courier provider
			if !voucher.IsPrinted {
				voucher.UpdateVoucherStatus(domain.VoucherStatusProcessing)
			}
		}
	case "cancelled":
		// If the voucher has not been printed, cancel it; otherwise, revert to processing
		if !voucher.IsPrinted {
			voucher.UpdateVoucherStatus(domain.VoucherStatusCancelled)
		}
	case "on-hold":
		if !voucher.IsPrinted {
			voucher.UpdateVoucherStatus(domain.VoucherStatusOnHold)
		}
	default:
	}

	voucher.UpdateVoucher(OrderRecord.Order.ShippingTotal, OrderRecord.Order.CustomerNote, OrderRecord.Order.Shipping, OrderRecord.Order.Billing, OrderRecord.Order.LineItems, OrderRecord.Order.PaymentMethod, OrderRecord.Order.Total)

	return vs.repo.UpdateVoucher(ctx, voucher, projectID, voucher.VoucherID, OrderRecord.OrderID)
}

// CreateHermesVoucher inserts a new Voucher into the database
func (vs *VoucherService) CreateHermesVoucher(ctx echo.Context, voucher *domain.HermesVoucerRequest, projectID string) (*domain.Voucher, error) {
	fullname := strings.Split(voucher.ReceiverName, " ")
	var firstName, lastName string
	if len(fullname) > 0 {
		firstName = fullname[0]
	}
	if len(fullname) > 1 {
		// Join the rest of the parts as the last name
		lastName = strings.Join(fullname[1:], " ")
	}

	shipping := &woocommerce.Shipping{
		FirstName: firstName,
		Address1:  voucher.ReceiverAddress,
		LastName:  lastName,
		City:      voucher.ReceiverCity,
		PostCode:  fmt.Sprintf("%d", voucher.ReceiverPostal),
	}

	billing := &woocommerce.Billing{
		Email:     voucher.ReceiverEmail,
		Phone:     voucher.ReceiverTelephone,
		FirstName: firstName,
		LastName:  lastName,
		Address1:  voucher.ReceiverAddress,
		City:      voucher.ReceiverCity,
		PostCode:  fmt.Sprintf("%d", voucher.ReceiverPostal),
	}
	orderID, err := strconv.ParseInt(voucher.OrderID, 10, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error converting string to int64:", err)
		return nil, err
	}
	newvoucher := domain.NewVoucher(projectID, fmt.Sprintf("%.2f", voucher.Cod), *voucher.Notes, shipping, billing, orderID, nil, "cod", fmt.Sprintf("%.2f", voucher.Cod))
	return vs.repo.CreateVoucher(ctx, newvoucher, projectID)

}

// GetVoucherByVoucherID retrieves a Voucher by its ID
func (vs *VoucherService) GetVoucherByVoucherID(ctx echo.Context, voucherID string) (*domain.Voucher, error) {
	return vs.repo.GetVoucherByVoucherID(ctx, voucherID)
}

// GetVoucherByOrderIDAndProjectID retrieves a Voucher by its ID
func (vs *VoucherService) GetVoucherByOrderIDAndProjectID(ctx echo.Context, orderID int64, projectID string) (*domain.Voucher, error) {
	return vs.repo.GetVoucherByOrderIDAndProjectID(ctx, orderID, projectID)
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
func (vs *VoucherService) UpdateVoucher(ctx echo.Context, order *o.OrderRecord, projectID string) (*domain.Voucher, error) {
	voucher, err := vs.repo.GetVoucherByOrderIDAndProjectID(ctx, order.Order.ID, projectID)
	if err != nil {
		return nil, err
	}

	if voucher == nil {
		slog.Error("Voucher not found")
		return nil, errors.New("voucher not found")
	}
	// Update the voucher status based on the order status
	switch order.Order.Status {
	case "completed":
		if !voucher.IsPrinted {
			voucher.UpdateVoucherStatus(domain.VoucherStatusCompleted)
		}
	case "processing":
		if voucher.Status == domain.VoucherStatusNew {
			voucher.UpdateVoucherStatus(domain.VoucherStatusNew)
		} else {
			//TODO: to update a voucher we need to check in what state we have it in the courier provider
			if !voucher.IsPrinted {
				voucher.UpdateVoucherStatus(domain.VoucherStatusProcessing)
			}
		}
	case "cancelled":
		// If the voucher has not been printed, cancel it; otherwise, revert to processing
		if !voucher.IsPrinted {
			voucher.UpdateVoucherStatus(domain.VoucherStatusCancelled)
		}
	case "on-hold":
		if !voucher.IsPrinted {
			voucher.UpdateVoucherStatus(domain.VoucherStatusOnHold)
		}
	default:
	}

	voucher.UpdateVoucher(order.Order.ShippingTotal, order.Order.CustomerNote, order.Order.Shipping, order.Order.Billing, order.Order.LineItems, order.Order.PaymentMethod, order.Order.Total)

	return vs.repo.UpdateVoucher(ctx, voucher, projectID, voucher.VoucherID, order.OrderID)
}

// DeleteVouchersByOrderIdandProjectID deletes a voucher by its ID
func (vs *VoucherService) DeleteVouchersByOrderIdandProjectID(ctx echo.Context, projectID string, orderID int64) error {
	voucher, err := vs.repo.GetVoucherByOrderIDAndProjectID(ctx, orderID, projectID)
	if err != nil {
		return err
	}
	if voucher == nil {
		slog.Error("Voucher not found")
		return errors.New("voucher not found")
	}
	if !voucher.IsPrinted {
		voucher.DeleteVoucher()

		_, err = vs.repo.UpdateVoucher(ctx, voucher, projectID, voucher.VoucherID, orderID)
		return err
	}
	return nil
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

// UpdateVoucherNewDetails retrieves a Voucher by its ID asynchronously
func (vs *VoucherService) UpdateVoucherNewDetails(ctx echo.Context, voucher *domain.Voucher, projectID string, courier4u *d.Courier4uExtension, redcourier *d.RedCourierExtension) (*domain.Voucher, error) {
	if !voucher.HasError {
		vID, err := strconv.ParseInt(voucher.VoucherID, 10, 64)
		if err != nil {
			return nil, err
		}
		tracking, err := vs.hermesSvc.TrackingHermerVoucherStatus(ctx, courier4u, redcourier, vID)
		if err != nil {
			return nil, err
		}
		if redcourier != nil || courier4u != nil {
			voucher.UpdateHermerVoucherTracking(tracking)
		}
	}

	return vs.repo.UpdateVoucherNewDetails(ctx, voucher, projectID)
}

// UpdateVoucherTracking updates a Voucher
func (vs *VoucherService) UpdateVoucherTracking(ctx echo.Context, voucher *domain.Voucher, projectID string, client *woocommerce.Client) (*domain.Voucher, error) {

	order, err := vs.woocommerceRepo.GetOrderByID(projectID, *voucher.OrderID)
	if err != nil {

	}

	for _, lineItem := range voucher.HermesTrackingStages.Data {
		if lineItem.Status == "Παραδόθηκε" {
			order.Order.Status = o.OrderStatusCompleted.String()
			order.Status = o.OrderStatusCompleted
			voucher.UpdateVoucherStatus(domain.VoucherStatusCompleted)
			if order != nil {
				go func() {
					_, err = client.Order.Update(&order.Order)
				}()

			}
		}
		if lineItem.Status == "Επιστροφή στον αποστολέα" {
			order.Order.Status = o.OrderStatusCancelled.String()
			order.Status = o.OrderStatusCancelled
			voucher.UpdateVoucherStatus(domain.VoucherStatusCancelled)
			if order != nil {
				go func() {
					_, err = client.Order.Update(&order.Order)
				}()
			}
		}
		if lineItem.Status == "Ακύρωση Αποστολής" {
			order.Order.Status = o.OrderStatusCancelled.String()
			order.Status = o.OrderStatusCancelled
			voucher.UpdateVoucherStatus(domain.VoucherStatusCancelled)
			if order != nil {
				go func() {
					_, err = client.Order.Update(&order.Order)
				}()
			}
		}

	}

	return vs.repo.UpdateVoucher(ctx, voucher, projectID, voucher.VoucherID, *voucher.OrderID)
}
