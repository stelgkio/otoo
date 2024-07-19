package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type CustomerRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID string             `bson:"projectId"`
	Event     string             `bson:"event"`
	Error     string             `bson:"error,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	CustomerID   int64              `bson:"customerId,omitempty"`
	Customer     woocommerce.Customer  `bson:"customer,omitempty"`
	CreatedAt time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt time.Time         `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive  bool               `json:"is_active" bson:"is_active,omitempty"`
	Orders    []int64           `bson:"orders,omitempty"`
}

func NewCustomerRecord(projectID uuid.UUID, event string, customerId int64, customer woocommerce.Customer) CustomerRecord {
	return CustomerRecord{
		ProjectID: projectID.String(),
		Event:     event,
		CustomerID:   customerId,
		Customer:     customer,
		IsActive:  true,       // Default value
		CreatedAt: time.Now(), // Initialize CreatedAt with the current time		
		Timestamp: time.Now(), // Initialize Timestamp with the current time
		Orders: []int64{},
	}
}

func (o *CustomerRecord) SoftDelete() {
	o.IsActive = false
	now := time.Now()
	o.DeletedAt = now
	o.Timestamp = now
}
