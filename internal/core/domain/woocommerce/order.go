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

// String returns the string representation of the ProductType
func (pt OrderStatus) String() string {
	return string(pt)
}

// OrderRecord represents an order record in MongoDB
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

// NewOrderRecord creates a new OrderRecord
func NewOrderRecord(projectID uuid.UUID, event string, orderID int64, order woocommerce.Order) OrderRecord {
	return OrderRecord{
		ProjectID: projectID.String(),
		Event:     event,
		OrderID:   orderID,
		Order:     order,
		IsActive:  true,             // Default value
		CreatedAt: time.Now().UTC(), // Initialize CreatedAt with the current time
		UpdatedAt: time.Now().UTC(), // Initialize UpdatedAt with the current time
		Timestamp: time.Now().UTC(), // Initialize Timestamp with the current time

	}
}

// SoftDelete updates the order record
func (o *OrderRecord) SoftDelete() {
	o.IsActive = false
	now := time.Now().UTC()
	o.DeletedAt = now
	o.Timestamp = now
}

// StringToOrderStatus converts a string to an OrderStatus
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

// OrderTableList represents an order table list
type OrderTableList struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	ProjectID   string               `bson:"projectId" json:"projectId"`
	Timestamp   time.Time            `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	OrderID     int64                `bson:"orderId,omitempty" json:"orderId,omitempty"`
	TotalAmount string               `bson:"total_amount,omitempty" json:"total_amount,omitempty"`
	Status      OrderStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Billing     woocommerce.Billing  `bson:"billing,omitempty" json:"billing,omitempty"`
	Shipping    woocommerce.Shipping `bson:"shipping,omitempty" json:"shipping,omitempty"`
}

// OrderTableResponde represents an order table response
type OrderTableResponde struct {
	Data []OrderTableList `json:"data"`
	Meta Meta             `json:"meta"`
}

// BulkActionRequest represents a bulk action request
type BulkActionRequest struct {
	Status string   `json:"status"`
	Orders []string `json:"orders"`
}
