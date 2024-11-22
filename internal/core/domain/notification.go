package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NotificationType represents a notification type
type NotificationType string

const (
	NotificationTypeInfo     = "info"
	NotificationTypeWarning  = "warning"
	NotificationTypeAlert    = "alert"
	NotificationTypeSuccess  = "success"
	NotificationTypeError    = "error"
	NotificationTypeReminder = "reminder"
	NotificationTypeMessage  = "message"
)

// Notification represents a notification
type Notification struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Link        string             `bson:"link,omitempty"`
	IsRead      bool               `bson:"is_read"`
	Timestamp   time.Time          `bson:"timestamp,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `bson:"deleted_at,omitempty"`
	IsActive    bool               `bson:"is_active,omitempty"`
	UserID      string             `bson:"user_id,omitempty"`
	ProjectID   string             `bson:"project_id,omitempty"`
	Type        string             `bson:"type,omitempty"`
}

// String returns the string representation of the NotificationType
func (pt NotificationType) String() string {
	return string(pt)
}

// StringToNotificationType converts a string to a NotificationType
func StringToNotificationType(notificationType string) (NotificationType, error) {
	switch notificationType {
	case string(NotificationTypeInfo):
		return NotificationTypeInfo, nil
	case string(NotificationTypeWarning):
		return NotificationTypeWarning, nil
	case string(NotificationTypeAlert):
		return NotificationTypeAlert, nil
	case string(NotificationTypeSuccess):
		return NotificationTypeSuccess, nil
	case string(NotificationTypeError):
		return NotificationTypeError, nil
	case string(NotificationTypeReminder):
		return NotificationTypeReminder, nil
	case string(NotificationTypeMessage):
		return NotificationTypeMessage, nil
	default:
		return "", errors.New("invalid notification type")
	}
}

// CreateDataSynchronizerNotification creates a new analytics notification
func CreateDataSynchronizerNotification(userID, projectID string) *Notification {

	// Create a new Notification instance
	newNotification := &Notification{
		ID:          primitive.NewObjectID(), // Generate a new ObjectID
		Title:       "Data Synchronizer",
		Description: "Keep your eShop and Otoo perfectly in sync. Effortlessly transfer product details, customer information, and order data.",
		Link:        "",
		IsRead:      false,      // Default value
		Timestamp:   time.Now(), // Current timestamp
		CreatedAt:   time.Now(), // Current time for creation
		UpdatedAt:   time.Now(), // Set it to current time initially
		IsActive:    true,       // Default to active
		UserID:      userID,
		ProjectID:   projectID,
		Type:        NotificationTypeInfo, // Set the type
	}

	return newNotification
}

// CreateAnalyticsNotification creates a new analytics notification
func CreateAnalyticsNotification(userID, projectID string) *Notification {
	// Create a new Notification instance
	return &Notification{
		ID:          primitive.NewObjectID(), // Generate a new ObjectID
		Title:       "Analytics",
		Description: "Every night, we perform essential calculations to keep your analytics up-to-date. This includes identifying best-selling products, tracking inventory balances, and analyzing customer behavior. Stay informed with the latest insights and ensure that your eShop and Otoo are perfectly aligned. Our automated processes will help you make informed decisions, optimize your inventory, and enhance customer satisfaction.",
		Link:        "",
		IsRead:      false,      // Default value
		Timestamp:   time.Now(), // Current timestamp
		CreatedAt:   time.Now(), // Current time for creation
		UpdatedAt:   time.Now(), // Set it to current time initially
		IsActive:    true,       // Default to active
		UserID:      userID,
		ProjectID:   projectID,
		Type:        NotificationTypeInfo, // Set the type
	}

}
