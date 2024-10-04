package domain

import (
	"time"

	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Voucher represents a voucher
type Voucher struct {
	ID                  primitive.ObjectID   `bson:"_id,omitempty"`
	ProjectID           string               `json:"projectId"  bson:"projectId,omitempty"`
	VoucherID           string               `json:"voucher_id"  bson:"voucher_id"`
	Cod                 string               `json:"cod"  bson:"cod,omitempty"`
	Status              VoucherStatus        `json:"status"  bson:"status,omitempty"`
	ShippingStatus      string               `json:"shipping_status"  bson:"shipping_status,omitempty"`
	Note                string               `json:"note"  bson:"note,omitempty"`
	Shipping            woocommerce.Shipping `json:"shipping"  bson:"shipping,omitempty"`
	AcsVoucherRequest   AcsVoucherRequest    `json:"acs_courier"  bson:"acs_courier"`
	HermesVoucerRequest HermesVoucerRequest  `json:"hermes_courier"  bson:"hermes_courier"`
	CreatedAt           time.Time            `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt           time.Time            `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt           time.Time            `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive            bool                 `json:"is_active" bson:"is_active,omitempty"`
}

// NewVoucher creates a new Voucher instance with the provided ProjectID, Cod, Note, and Shipping.
func NewVoucher(projectID, cod, note string, shipping woocommerce.Shipping) *Voucher {
	return &Voucher{
		ID:        primitive.NewObjectID(), // Generate a new ObjectID for the voucher
		ProjectID: projectID,
		Cod:       cod,
		Note:      note,
		Shipping:  shipping,
		Status:    VoucherStatusNew, // Set a default status if needed
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true, // Assuming the voucher is active upon creation
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
