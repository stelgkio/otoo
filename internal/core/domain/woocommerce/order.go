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
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID    string             `bson:"projectId"`
	Event        string             `bson:"event"`
	Error        string             `bson:"error"`
	Timestamp    time.Time          `bson:"timestamp,omitempty"`
	OrderID      int64              `bson:"orderId,omitempty"`
	Order        woocommerce.Order  `bson:"order,omitempty"`
	CreatedAt    time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	OrderCreated time.Time          `json:"order_created"  bson:"order_created,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt    time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive     bool               `json:"is_active" bson:"is_active,omitempty"`
	Status       OrderStatus        `bson:"status,omitempty"`
}

// NewOrderRecord creates a new OrderRecord
func NewOrderRecord(projectID uuid.UUID, event string, orderID int64, order woocommerce.Order) OrderRecord {

	_, orderData, _ := ConvertDateString(order.DateCreatedGmt)
	return OrderRecord{
		ProjectID:    projectID.String(),
		Event:        event,
		OrderID:      orderID,
		Order:        order,
		IsActive:     true,             // Default value
		CreatedAt:    time.Now().UTC(), // Initialize CreatedAt with the current time
		UpdatedAt:    time.Now().UTC(), // Initialize UpdatedAt with the current time
		Timestamp:    time.Now().UTC(), // Initialize Timestamp with the current time
		OrderCreated: orderData,
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
	ID             primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"`
	ProjectID      string                 `bson:"projectId" json:"projectId"`
	Timestamp      time.Time              `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	OrderCreated   time.Time              `json:"order_created"  bson:"order_created,omitempty"`
	OrderID        int64                  `bson:"orderId,omitempty" json:"orderId,omitempty"`
	TotalAmount    string                 `bson:"total_amount,omitempty" json:"total_amount,omitempty"`
	Status         OrderStatus            `bson:"status,omitempty" json:"status,omitempty"`
	Billing        woocommerce.Billing    `bson:"billing,omitempty" json:"billing,omitempty"`
	Shipping       woocommerce.Shipping   `bson:"shipping,omitempty" json:"shipping,omitempty"`
	Products       []woocommerce.LineItem `bson:"products,omitempty" json:"products,omitempty"`
	CurrencySymbol string                 `bson:"currency_symbol,omitempty" json:"currency_symbol,omitempty"`
	PaymentMethod  string                 `bson:"payment_method,omitempty" json:"payment_method,omitempty"`
	CustomerNote   string                 `bson:"customer_note,omitempty" json:"customer_note,omitempty"`
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

// ConvertDateString  Function to convert and format the date
func ConvertDateString(dateStr string) (string, time.Time, error) {
	// Define the layout for parsing the original string
	layoutInput := "2006-01-02T15:04:05"

	// Parse the input date string into a time.Time object
	parsedTime, err := time.Parse(layoutInput, dateStr)
	if err != nil {
		return "", parsedTime, err
	}

	// Change the date to 2024-09-25 while keeping the same time
	newDate := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, parsedTime.Location())

	// Define the layout for formatting the output string
	layoutOutput := "2006-01-02 15:04:05"

	// Format the new date into the desired string format
	formattedTime := newDate.Format(layoutOutput)

	return formattedTime, newDate, nil
}
