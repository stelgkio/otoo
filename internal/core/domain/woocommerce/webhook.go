package domain

import (
	"time"

	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebhookResult struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	ProjectID string              `bson:"projectId"`
	Event     string              `bson:"event"`
	WebhookID int64               `bson:"webhookId,omitempty"`
	Error     string              `bson:"error,omitempty"`
	Timestamp time.Time           `bson:"timestamp,omitempty"`
	Webhook   woocommerce.Webhook `bson:"webhook,omitempty"`
	CreatedAt time.Time           `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt *time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive  bool                `json:"is_active" bson:"is_active,omitempty"`
}
