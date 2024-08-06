package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderStatus enum
type OrderStatus string

const (
	OrderStatusAll           OrderStatus = "all"
	OrderStatusPending       OrderStatus = "pending"
	OrderStatusProcessing    OrderStatus = "processing"
	OrderStatusOnHold        OrderStatus = "on-hold"
	OrderStatusCompleted     OrderStatus = "completed"
	OrderStatusCancelled     OrderStatus = "cancelled"
	OrderStatusFailed        OrderStatus = "failed"
	OrderStatusCheckoutDraft OrderStatus = "checkout-draft"
)

type OrderRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID string             `bson:"projectId"`
	Event     string             `bson:"event"`
	Error     string             `bson:"error,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	OrderID   int64              `bson:"orderId,omitempty"`
	Order     woocommerce.Order  `bson:"order,omitempty"`
	CreatedAt time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive  bool               `json:"is_active" bson:"is_active,omitempty"`
	Status    OrderStatus        `bson:"status,omitempty"`
}

func NewOrderRecord(projectID uuid.UUID, event string, orderId int64, order woocommerce.Order) OrderRecord {
	return OrderRecord{
		ProjectID: projectID.String(),
		Event:     event,
		OrderID:   orderId,
		Order:     order,
		IsActive:  true,             // Default value
		CreatedAt: time.Now().UTC(), // Initialize CreatedAt with the current time
		UpdatedAt: time.Now().UTC(), // Initialize UpdatedAt with the current time
		Timestamp: time.Now().UTC(), // Initialize Timestamp with the current time

	}
}

func (o *OrderRecord) SoftDelete() {
	o.IsActive = false
	now := time.Now().UTC()
	o.DeletedAt = now
	o.Timestamp = now
}

func StringToOrderStatus(status string) (OrderStatus, error) {
	switch status {
	case string(OrderStatusAll):
		return OrderStatusAll, nil
	case string(OrderStatusPending):
		return OrderStatusPending, nil
	case string(OrderStatusProcessing):
		return OrderStatusProcessing, nil
	case string(OrderStatusOnHold):
		return OrderStatusOnHold, nil
	case string(OrderStatusCompleted):
		return OrderStatusCompleted, nil
	case string(OrderStatusCancelled):
		return OrderStatusCancelled, nil
	case string(OrderStatusFailed):
		return OrderStatusFailed, nil
	case string(OrderStatusCheckoutDraft):
		return OrderStatusCheckoutDraft, nil
	default:
		return "", errors.New("invalid order status")
	}
}

type OrderTableList struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProjectID   string             `bson:"projectId" json:"projectId"`
	Timestamp   time.Time          `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	OrderID     int64              `bson:"orderId,omitempty" json:"orderId,omitempty"`
	TotalAmount string             `bson:"total_amount,omitempty" json:"total_amount,omitempty"`
	Status      OrderStatus        `bson:"status,omitempty" json:"status,omitempty"`
}

type OrderTableResponde struct {
	Data []OrderTableList `json:"data"`
	Meta Meta             `json:"meta"`
}
