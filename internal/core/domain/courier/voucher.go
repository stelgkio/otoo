package domain

import (
	"errors"
	"time"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Voucher represents a voucher
type Voucher struct {
	ID                  primitive.ObjectID     `json:"Id" bson:"_id,omitempty"`
	ProjectID           string                 `json:"projectId"  bson:"projectId,omitempty"`
	OrderID             int64                  `json:"orderId"  bson:"orderId,omitempty"`
	VoucherID           string                 `json:"voucher_id"  bson:"voucher_id"`
	Cod                 string                 `json:"cod"  bson:"cod,omitempty"`
	Status              VoucherStatus          `json:"status"  bson:"status,omitempty"`
	ShippingStatus      string                 `json:"shipping_status"  bson:"shipping_status,omitempty"`
	Note                string                 `json:"note"  bson:"note,omitempty"`
	Shipping            *woocommerce.Shipping  `json:"shipping"  bson:"shipping,omitempty"`
	Billing             *woocommerce.Billing   `json:"billing"  bson:"billing,omitempty"`
	Products            []woocommerce.LineItem `bson:"products,omitempty" json:"products,omitempty"`
	AcsVoucherRequest   *AcsVoucherRequest     `json:"acs_courier"  bson:"acs_courier"`
	HermesVoucerRequest *HermesVoucerRequest   `json:"hermes_courier"  bson:"hermes_courier"`
	CreatedAt           time.Time              `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt           time.Time              `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt           *time.Time             `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive            bool                   `json:"is_active" bson:"is_active,omitempty"`
	IsPrinted           bool                   `json:"is_printed" bson:"is_printed"`
}

// NewVoucher creates a new Voucher instance with the provided ProjectID, Cod, Note, and Shipping.
func NewVoucher(projectID, cod, note string, shipping *woocommerce.Shipping, billing *woocommerce.Billing, orderID int64, products []woocommerce.LineItem) *Voucher {
	return &Voucher{
		ProjectID:           projectID,
		Cod:                 cod,
		Note:                note,
		Shipping:            shipping,
		Billing:             billing,
		Products:            products,
		OrderID:             orderID,
		Status:              VoucherStatusNew, // Set a default status if needed
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		IsActive:            true, // Assuming the voucher is active upon creation
		IsPrinted:           false,
		AcsVoucherRequest:   nil,
		HermesVoucerRequest: nil,
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

// UpdateVoucher updates the Voucher with the provided Cod, Note, and Shipping.
func (v *Voucher) UpdateVoucher(cod, note string, shipping *woocommerce.Shipping, billing *woocommerce.Billing, products []woocommerce.LineItem) *Voucher {
	v.Cod = cod
	v.Note = note
	v.Shipping = shipping
	v.Billing = billing
	v.Products = products
	v.UpdatedAt = time.Now()
	v.DeletedAt = nil
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
	ID        primitive.ObjectID     `bson:"_id,omitempty" json:"Id,omitempty"`
	ProjectID string                 `bson:"projectId" json:"projectId"`
	OrderID   int64                  `bson:"orderId,omitempty" json:"orderId,omitempty"`
	VoucherID string                 `bson:"voucher_id,omitempty" json:"voucher_id,omitempty"`
	Status    VoucherStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Billing   woocommerce.Billing    `bson:"billing,omitempty" json:"billing,omitempty"`
	Shipping  woocommerce.Shipping   `bson:"shipping,omitempty" json:"shipping,omitempty"`
	Products  []woocommerce.LineItem `bson:"products,omitempty" json:"products,omitempty"`
	Cod       string                 `bson:"cod,omitempty" json:"cod,omitempty"`
	CreateAt  time.Time              `bson:"created_at,omitempty" json:"created_at,omitempty"`
	IsPrinted bool                   `bson:"is_printed,omitempty" json:"is_printed,omitempty"`
}

// VoucherTableResponde represents an order table response
type VoucherTableResponde struct {
	Data []VoucherTableList `json:"data"`
	Meta w.Meta             `json:"meta"`
}
