package domain

import (
	"time"

	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Voucher represents a voucher
type Voucher struct {
	ID                  primitive.ObjectID    `bson:"_id,omitempty"`
	ProjectID           string                `json:"projectId"  bson:"projectId,omitempty"`
	OrderID             int64                 `json:"orderId"  bson:"orderId,omitempty"`
	VoucherID           string                `json:"voucher_id"  bson:"voucher_id"`
	Cod                 string                `json:"cod"  bson:"cod,omitempty"`
	Status              VoucherStatus         `json:"status"  bson:"status,omitempty"`
	ShippingStatus      string                `json:"shipping_status"  bson:"shipping_status,omitempty"`
	Note                string                `json:"note"  bson:"note,omitempty"`
	Shipping            *woocommerce.Shipping `json:"shipping"  bson:"shipping,omitempty"`
	AcsVoucherRequest   *AcsVoucherRequest    `json:"acs_courier"  bson:"acs_courier"`
	HermesVoucerRequest *HermesVoucerRequest  `json:"hermes_courier"  bson:"hermes_courier"`
	CreatedAt           time.Time             `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt           time.Time             `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt           time.Time             `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive            bool                  `json:"is_active" bson:"is_active,omitempty"`
	IsPrinted           bool                  `json:"is_printed" bson:"is_printed"`
}

// NewVoucher creates a new Voucher instance with the provided ProjectID, Cod, Note, and Shipping.
func NewVoucher(projectID, cod, note string, shipping *woocommerce.Shipping, orderID int64) *Voucher {
	return &Voucher{
		ProjectID:           projectID,
		Cod:                 cod,
		Note:                note,
		Shipping:            shipping,
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
func (v *Voucher) UpdateVoucher(cod, note string, shipping *woocommerce.Shipping) *Voucher {
	v.Cod = cod
	v.Note = note
	v.Shipping = shipping
	v.UpdatedAt = time.Now()
	return v
}
