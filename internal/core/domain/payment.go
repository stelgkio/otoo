package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Payment represents a payment record
type Payment struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID          string             `json:"projectId" bson:"projectId,omitempty"`
	ProjectExtensionID string             `json:"projectextensionId" bson:"projectextensionId,omitempty"`
	ExtensionName      string             `json:"extension_name" bson:"extension_name,omitempty"`
	CreatedAt          time.Time          `bson:"created_at,omitempty"`
	UpdatedAt          time.Time          `bson:"updated_at,omitempty"`
	DeletedAt          time.Time          `bson:"deleted_at,omitempty"`
	NextPaymentAt      time.Time          `bson:"next_payment_at,omitempty"`
	Amount             int64              `json:"amount" bson:"amount,omitempty"`
	IsActive           bool               `bson:"is_active,omitempty"`
	IsPaid             bool               `bson:"is_paid,omitempty"`
	IsFail             bool               `bson:"fail,omitempty"`
}

// NewPaymentSuccess new success payment
func NewPaymentSuccess(projectID, projectExtensionID, extensionName string, amount int64, nextPaymentAt time.Time) *Payment {
	return &Payment{
		ID:                 primitive.NewObjectID(), // Automatically generate a new ObjectID
		ProjectID:          projectID,
		ProjectExtensionID: projectExtensionID,
		ExtensionName:      extensionName,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		NextPaymentAt:      nextPaymentAt,
		Amount:             amount,
		IsActive:           true,  // Payment is active by default
		IsPaid:             true,  // Not paid initially
		IsFail:             false, // Not failed initially
	}
}

// NewPaymentFail new fail payment
func NewPaymentFail(projectID, projectExtensionID, extensionName string, amount int64, nextPaymentAt time.Time) *Payment {
	return &Payment{
		ID:                 primitive.NewObjectID(), // Automatically generate a new ObjectID
		ProjectID:          projectID,
		ProjectExtensionID: projectExtensionID,
		ExtensionName:      extensionName,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		NextPaymentAt:      nextPaymentAt,
		Amount:             amount,
		IsActive:           true,  // Payment is active by default
		IsPaid:             false, // Not paid initially
		IsFail:             true,  // Not failed initially
	}
}

// PaymentTableList represents an payment table list
type PaymentTableList struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProjectID     string             `bson:"projectId" json:"projectId,omitempty"`
	ExtensionName string             `bson:"extension_name,omitempty" json:"extension_name"`
	Amount        float64            `bson:"amount,omitempty" json:"amount"`
	IsPaid        bool               `bson:"is_paid,omitempty" json:"is_paid"`
	IsFail        bool               `bson:"fail,omitempty" json:"fail"`
	NextPaymentAt time.Time          `bson:"next_payment_at,omitempty" json:"next_payment_at"`
	CreatedAt     time.Time          `bson:"created_at,omitempty" json:"created_at"`
}

// PaymentTableResponde represents an payment table response
type PaymentTableResponde struct {
	Data []PaymentTableList `json:"data"`
	Meta Meta               `json:"meta"`
}
