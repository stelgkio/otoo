package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
