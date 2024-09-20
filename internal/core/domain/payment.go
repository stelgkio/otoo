package domain

import "time"

// Payment represents a payment record
type Payment struct {
	UserID        string    `json:"user_id" bson:"user_id,omitempty"`
	ProjectID     string    `json:"project_id" bson:"project_id,omitempty"`
	ExtensionID   string    `json:"extension_id" bson:"extension_id,omitempty"`
	CreatedAt     time.Time `bson:"created_at,omitempty"`
	UpdatedAt     time.Time `bson:"updated_at,omitempty"`
	DeletedAt     time.Time `bson:"deleted_at,omitempty"`
	NextPaymentAt time.Time `bson:"next_payment_at,omitempty"`
	IsActive      bool      `bson:"is_active,omitempty"`
	IsPaid        bool      `bson:"is_paid,omitempty"`
	IsFail        bool      `bson:"fail,omitempty"`
}
