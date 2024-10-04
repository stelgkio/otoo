package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WebhookRecord represents a webhook record in MongoDB
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

// NewWebhookRecord creates a new WebhookRecord
func NewWebhookRecord(projectID uuid.UUID, event string, webhookID int64, webhook woocommerce.Webhook) WebhookRecord {
	return WebhookRecord{
		ProjectID: projectID.String(),
		Event:     event,
		WebhookID: webhookID,
		Webhook:   webhook,
		IsActive:  true,             // Default value
		CreatedAt: time.Now().UTC(), // Initialize CreatedAt with the current time
		UpdatedAt: time.Now().UTC(), // Initialize UpdatedAt with the current time
	}
}

// SoftDelete marks the record as deleted
func (wr *WebhookRecord) SoftDelete() {
	wr.IsActive = false
	now := time.Now().UTC()
	wr.DeletedAt = &now
}

// WebhookTableList represents an order table list
type WebhookTableList struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProjectID string             `bson:"projectId" json:"projectId,omitempty"`
	Event     string             `bson:"event" json:"event,omitempty"`
	WebhookID int64              `bson:"webhookId,omitempty" json:"webhookId,omitempty"`
	Status    WebhookStatus      `bson:"status,omitempty" json:"status,omitempty"`
}

// WebhookTableResponde represents an order table response
type WebhookTableResponde struct {
	Data []WebhookTableList `json:"data"`
	Meta Meta               `json:"meta"`
}

// WebhookBulkActionRequest represents a bulk action request
type WebhookBulkActionRequest struct {
	Status   string   `json:"status"`
	Webhooks []string `json:"webhooks"`
}

// WebhookStatus enum
type WebhookStatus string

const (
	// WebhookStatusActive represents an active webhook
	WebhookStatusActive WebhookStatus = "active"
	// WebhookStatusPaused represents a paused webhook
	WebhookStatusPaused WebhookStatus = "paused"
	// WebhookStatusDisabled represents a disabled webhook
	WebhookStatusDisabled WebhookStatus = "disabled"
)

// String returns the string representation of the ProductType
func (ws WebhookStatus) String() string {
	return string(ws)
}
