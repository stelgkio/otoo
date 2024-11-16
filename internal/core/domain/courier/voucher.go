package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Voucher represents a voucher
type Voucher struct {
	ID                   primitive.ObjectID     `json:"Id" bson:"_id,omitempty"`
	ProjectID            string                 `json:"projectId"  bson:"projectId,omitempty"`
	OrderID              *int64                 `json:"orderId"  bson:"orderId,omitempty"`
	CustomOrderID        *int64                 `json:"custom_orderId"  bson:"custom_orderId,omitempty"`
	VoucherID            string                 `json:"voucher_id"  bson:"voucher_id"`
	TotalAmount          string                 `json:"total_amount"  bson:"total_amount,omitempty"`
	PaymentMethod        string                 `json:"paymentmethod"  bson:"paymentmethod,omitempty"`
	Cod                  string                 `json:"cod"  bson:"cod,omitempty"`
	Status               VoucherStatus          `json:"status"  bson:"status,omitempty"`
	ShippingStatus       string                 `json:"shipping_status"  bson:"shipping_status,omitempty"`
	Note                 string                 `json:"note"  bson:"note,omitempty"`
	Shipping             *woocommerce.Shipping  `json:"shipping"  bson:"shipping,omitempty"`
	Billing              *woocommerce.Billing   `json:"billing"  bson:"billing,omitempty"`
	Products             []woocommerce.LineItem `bson:"products,omitempty" json:"products,omitempty"`
	AcsVoucherRequest    *AcsVoucherRequest     `json:"acs_courier"  bson:"acs_courier"`
	HermesVoucerRequest  *HermesVoucerRequest   `json:"hermes_courier"  bson:"hermes_courier"`
	CreatedAt            time.Time              `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt            time.Time              `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt            *time.Time             `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive             bool                   `json:"is_active" bson:"is_active,omitempty"`
	IsPrinted            bool                   `json:"is_printed" bson:"is_printed"`
	HasError             bool                   `json:"has_error" bson:"has_error"`
	Error                string                 `json:"error" bson:"error"`
	CourierProvider      string                 `json:"courier_provider" bson:"courier_provider"`
	TrackingStatus       string                 `json:"tracking_status"  bson:"tracking_status,omitempty"`
	HermesTrackingStages *TrackingResponse      `json:"hermes_tracking_stages"  bson:"hermes_tracking_stages,omitempty"`
}

// NewVoucher creates a new Voucher instance with the provided ProjectID, Cod, Note, and Shipping.
func NewVoucher(projectID, cod,
	note string,
	shipping *woocommerce.Shipping,
	billing *woocommerce.Billing,
	orderID int64,
	products []woocommerce.LineItem,
	paymentMethod string,
	totalAmount string) *Voucher {
	return &Voucher{
		ProjectID:           projectID,
		Cod:                 cod,
		Note:                note,
		Shipping:            shipping,
		Billing:             billing,
		Products:            products,
		OrderID:             &orderID,
		Status:              VoucherStatusNew, // Set a default status if needed
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		IsActive:            true, // Assuming the voucher is active upon creation
		IsPrinted:           false,
		AcsVoucherRequest:   nil,
		HermesVoucerRequest: nil,
		HasError:            false,
		Error:               "",
		TotalAmount:         totalAmount,
		PaymentMethod:       paymentMethod,
	}
}

// VoucherStatus enum
type VoucherStatus string

const (
	VoucherStatusAll        VoucherStatus = "all"
	VoucherStatusNew        VoucherStatus = "new"
	VoucherStatusProcessing VoucherStatus = "processing"
	VoucherStatusOnHold     VoucherStatus = "on-hold"
	VoucherStatusCompleted  VoucherStatus = "completed"
	VoucherStatusCancelled  VoucherStatus = "cancelled"
	VoucherStatusFailed     VoucherStatus = "failed"
)

// StringToVoucherStatus converts a string to a VoucherStatus
func StringToVoucherStatus(status string) (VoucherStatus, error) {
	switch status {
	case string(VoucherStatusAll):
		return VoucherStatusAll, nil
	case string(VoucherStatusNew):
		return VoucherStatusNew, nil
	case string(VoucherStatusProcessing):
		return VoucherStatusProcessing, nil
	case string(VoucherStatusOnHold):
		return VoucherStatusOnHold, nil
	case string(VoucherStatusCompleted):
		return VoucherStatusCompleted, nil
	case string(VoucherStatusCancelled):
		return VoucherStatusCancelled, nil
	case string(VoucherStatusFailed):
		return VoucherStatusFailed, nil
	default:
		return "", errors.New("invalid voucher status")
	}
}

// String returns the string representation of the ProductType
func (pt VoucherStatus) String() string {
	return string(pt)
}

// UpdateVoucherStatus updates the Voucher status and updates the UpdatedAt timestamp.
func (v *Voucher) UpdateVoucherStatus(status VoucherStatus) *Voucher {
	v.Status = status
	v.UpdatedAt = time.Now()
	return v
}

// UpdateVoucherIsPrinted updates the Voucher status and updates the UpdatedAt timestamp.
func (v *Voucher) UpdateVoucherIsPrinted(printed bool) *Voucher {
	v.IsPrinted = printed
	v.UpdatedAt = time.Now()
	return v
}
func (v *Voucher) SetVoucher(voucherId int64) *Voucher {
	v.VoucherID = fmt.Sprintf("%d", voucherId)

	return v
}

// SetCustomOrderID updates the Voucher status and updates the UpdatedAt timestamp.
func (v *Voucher) SetCustomOrderID(customorderID *int64, customOrderID bool) *Voucher {
	if customOrderID {
		v.OrderID = customorderID
	}

	return v
}

// UpdateVoucherisError updates the Voucher status and updates the UpdatedAt timestamp.
func (v *Voucher) UpdateVoucherError(errormsg string) *Voucher {
	if errormsg == "" {
		v.HasError = false
		v.Error = errormsg
		v.UpdatedAt = time.Now()
		return v
	}
	v.HasError = true
	v.Error = errormsg
	v.UpdatedAt = time.Now()
	return v

}

// UpdateVoucherisError updates the Voucher status and updates the UpdatedAt timestamp.
func (v *Voucher) UpdateVoucherProvider(provider string) *Voucher {
	v.CourierProvider = provider
	v.UpdatedAt = time.Now()
	return v
}

// UpdateHermerVoucherTracking updates the Voucher status and updates the UpdatedAt timestamp.
func (v *Voucher) UpdateHermerVoucherTracking(tracking *TrackingResponse) *Voucher {
	if tracking == nil || tracking.Data == nil {
		return v
	}

	v.TrackingStatus = tracking.Data[len(tracking.Data)-1].Status
	v.HermesTrackingStages = tracking
	return v
}

// UpdateVoucherAcs updates the Voucher status and updates the UpdatedAt timestamp.
func (v *Voucher) UpdateVoucherAcs(acsVoucherRequest *AcsVoucherRequest) *Voucher {
	v.AcsVoucherRequest = acsVoucherRequest
	v.UpdatedAt = time.Now()
	return v
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

// UpdateVoucherHermes updates the Voucher status and updates the UpdatedAt timestamp.
func (v *Voucher) UpdateVoucherHermes(hermesVoucherRequest *HermesVoucerRequest) *Voucher {
	firstName, lastName := SplitFullName(hermesVoucherRequest.ReceiverName)

	v.HermesVoucerRequest = hermesVoucherRequest
	v.TotalAmount = fmt.Sprintf("%.2f", hermesVoucherRequest.Cod)
	v.Shipping.Address1 = hermesVoucherRequest.ReceiverAddress
	v.Shipping.City = hermesVoucherRequest.ReceiverCity
	v.Shipping.PostCode = fmt.Sprintf("%d", hermesVoucherRequest.ReceiverPostal)
	v.Billing.Phone = hermesVoucherRequest.ReceiverTelephone
	v.Shipping.FirstName = firstName
	v.Shipping.LastName = lastName
	if hermesVoucherRequest.Notes != nil {
		v.Note = *hermesVoucherRequest.Notes
	} else {
		v.Note = ""
	}
	v.UpdatedAt = time.Now()
	return v
}

// UpdateVoucher updates the Voucher with the provided Cod, Note, and Shipping.
func (v *Voucher) UpdateVoucher(cod, note string, shipping *woocommerce.Shipping, billing *woocommerce.Billing, products []woocommerce.LineItem, paymentMethod string, totalAmount string) *Voucher {
	v.Cod = cod
	v.Note = note
	v.Shipping = shipping
	v.Billing = billing
	v.Products = products
	v.UpdatedAt = time.Now()
	v.DeletedAt = nil
	v.PaymentMethod = paymentMethod
	v.TotalAmount = totalAmount
	return v
}

// DeleteVoucher updates the Voucher with the provided Cod, Note, and Shipping.
func (v *Voucher) DeleteVoucher() *Voucher {
	v.IsActive = false
	v.UpdatedAt = time.Now()
	now := time.Now()
	v.DeletedAt = &now
	return v
}

// VoucherTableList represents an order table list
type VoucherTableList struct {
	ID                   primitive.ObjectID     `bson:"_id,omitempty" json:"Id,omitempty"`
	ProjectID            string                 `bson:"projectId" json:"projectId"`
	OrderID              int64                  `bson:"orderId,omitempty" json:"orderId,omitempty"`
	VoucherID            string                 `json:"voucherId" bson:"voucher_id,omitempty"`
	Status               VoucherStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Billing              woocommerce.Billing    `bson:"billing,omitempty" json:"billing,omitempty"`
	Shipping             woocommerce.Shipping   `bson:"shipping,omitempty" json:"shipping,omitempty"`
	Products             []woocommerce.LineItem `bson:"products,omitempty" json:"products,omitempty"`
	Cod                  string                 `bson:"cod,omitempty" json:"cod,omitempty"`
	CreateAt             time.Time              `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt            time.Time              `json:"updated_at"  bson:"updated_at,omitempty"`
	IsPrinted            bool                   `bson:"is_printed,omitempty" json:"is_printed,omitempty"`
	HasError             bool                   `json:"has_error" bson:"has_error"`
	Error                string                 `json:"error" bson:"error"`
	CourierProvider      string                 `json:"courier_provider" bson:"courier_provider"`
	TotalAmount          string                 `json:"total_amount"  bson:"total_amount,omitempty"`
	PaymentMethod        string                 `json:"paymentmethod"  bson:"paymentmethod,omitempty"`
	Note                 string                 `json:"note"  bson:"note,omitempty"`
	AcsVoucherRequest    *AcsVoucherRequest     `json:"acs_courier"  bson:"acs_courier"`
	HermesVoucerRequest  *HermesVoucerRequest   `json:"hermes_courier"  bson:"hermes_courier"`
	TrackingStatus       string                 `json:"tracking_status"  bson:"tracking_status,omitempty"`
	HermesTrackingStages *TrackingResponse      `json:"hermes_tracking_stages"  bson:"hermes_tracking_stages,omitempty"`
}

// VoucherTableResponde represents an order table response
type VoucherTableResponde struct {
	Data []VoucherTableList `json:"data"`
	Meta w.Meta             `json:"meta"`
}
