package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ProductRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID string             `bson:"projectId"`
	Event     string             `bson:"event"`
	Error     string             `bson:"error,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	ProductID int64              `bson:"productId,omitempty"`
	Product   woocommerce.Product  `bson:"product,omitempty"`
	CreatedAt time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt time.Time         `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive  bool               `json:"is_active" bson:"is_active,omitempty"`
	Orders    []int64           `bson:"orders,omitempty"`
	ParentId  int64              `bson:"parentId,omitempty"`
}

func NewProductRecord(projectID uuid.UUID, event string, productId int64, product woocommerce.Product,parentId int64) ProductRecord {
	return ProductRecord{
		ProjectID: projectID.String(),
		Event:     event,
		ProductID:   productId,
		Product:     product,
		IsActive:  true,       // Default value
		CreatedAt: time.Now(), // Initialize CreatedAt with the current time		
		Timestamp: time.Now(), // Initialize Timestamp with the current time
		Orders: []int64{},
		ParentId: parentId,
	}
}

func (o *ProductRecord) SoftDelete() {
	o.IsActive = false
	now := time.Now()
	o.DeletedAt = now
	o.Timestamp = now
}
