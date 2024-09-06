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
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Link        string             `json:"link" bson:"link,omitempty"`
	IsRead      bool               `json:"is_read" bson:"is_read,omitempty"`
	CreatedAt   time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive    bool               `json:"is_active" bson:"is_active,omitempty"`
	UserID      string             `json:"user_id" bson:"user_id,omitempty"`
	ProjectID   string             `json:"project_id" bson:"project_id,omitempty"`
	Type        string             `json:"type" bson:"type,omitempty"`
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
