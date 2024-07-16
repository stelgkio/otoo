package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebhookRecord struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	ProjectID string              `bson:"projectId"`
	Event     string              `bson:"event"`
	WebhookID int64               `bson:"webhookId,omitempty"`
	Error     string              `bson:"error,omitempty"`
	Webhook   woocommerce.Webhook `bson:"webhook,omitempty"`
	CreatedAt time.Time           `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt *time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive  bool                `json:"is_active" bson:"is_active,omitempty"`
}

func NewWebhookRecord(projectID uuid.UUID, event string, webhookID int64, webhook woocommerce.Webhook) WebhookRecord {
	return WebhookRecord{
		ProjectID: projectID.String(),
		Event:     event,
		WebhookID: webhookID,
		Webhook:   webhook,
		IsActive:  true,       // Default value
		CreatedAt: time.Now(), // Initialize CreatedAt with the current time
		UpdatedAt: time.Now(), // Initialize UpdatedAt with the current time
	}
}

func (wr *WebhookRecord) SoftDelete() {
	wr.IsActive = false
	now := time.Now()
	wr.DeletedAt = &now
}
