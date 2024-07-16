package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID uuid.UUID          `bson:"projectId"`
	Event     string             `bson:"event"`
	Error     string             `bson:"error,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	OrderID   int64              `bson:"orderId,omitempty"`
	Order     woocommerce.Order  `bson:"order,omitempty"`
	CreatedAt time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt *time.Time         `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive  bool               `json:"is_active" bson:"is_active,omitempty"`
}

func NewOrderRecord(projectID uuid.UUID, event string, orderId int64, order woocommerce.Order) OrderRecord {
	return OrderRecord{
		ProjectID: projectID,
		Event:     event,
		OrderID:   orderId,
		Order:     order,
		IsActive:  true,       // Default value
		CreatedAt: time.Now(), // Initialize CreatedAt with the current time
		UpdatedAt: time.Now(), // Initialize UpdatedAt with the current time
		Timestamp: time.Now(), // Initialize Timestamp with the current time
	}
}

func (o *OrderRecord) SoftDelete() {
	o.IsActive = false
	now := time.Now()
	o.DeletedAt = &now
}
