package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRecord struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	ProjectID string              `bson:"projectId"`
	Event     string              `bson:"event"`
	Error     string              `bson:"error,omitempty"`
	Timestamp time.Time           `bson:"timestamp,omitempty"`
	ProductID int64               `bson:"productId,omitempty"`
	Product   woocommerce.Product `bson:"product,omitempty"`
	CreatedAt time.Time           `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt time.Time           `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive  bool                `json:"is_active" bson:"is_active,omitempty"`
	Orders    []int64             `bson:"orders,omitempty"`
	ParentId  int64               `bson:"parentId,omitempty"`
}

func NewProductRecord(projectID uuid.UUID, event string, productId int64, product woocommerce.Product, parentId int64) ProductRecord {
	return ProductRecord{
		ProjectID: projectID.String(),
		Event:     event,
		ProductID: productId,
		Product:   product,
		IsActive:  true,       // Default value
		CreatedAt: time.Now(), // Initialize CreatedAt with the current time
		Timestamp: time.Now(), // Initialize Timestamp with the current time
		Orders:    []int64{},
		ParentId:  parentId,
	}
}

// UpdateProductRecord updates the product record
func (o *ProductRecord) SoftDelete() {
	o.IsActive = false
	now := time.Now()
	o.DeletedAt = now
	o.Timestamp = now
}

// ProductTableList represents a product table list
type ProductTableList struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProjectID        string             `bson:"projectId" json:"projectId"`
	ProductImageURL  string             `bson:"product_image_url,omitempty" json:"product_image_url,omitempty"`
	Timestamp        time.Time          `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	ProductID        int64              `bson:"productId,omitempty" json:"productId,omitempty"`
	TotalAmountSpend string             `bson:"total_amount,omitempty" json:"total_amount,omitempty"`
	TotalOrders      int                `bson:"total_orders,omitempty" json:"total_orders,omitempty"`
	SKU              string             `bson:"sku,omitempty" json:"sku,omitempty"`
	Name             string             `bson:"name,omitempty" json:"name,omitempty"`
	Price            string             `bson:"price,omitempty" json:"price,omitempty"`
	Category         string             `bson:"category,omitempty" json:"category,omitempty"`
	ProductType      string             `bson:"product_type,omitempty" json:"product_type,omitempty"`
}

// ProductTableResponde represents a product table response
type ProductTableResponde struct {
	Data []ProductTableList `json:"data"`
	Meta Meta               `json:"meta"`
}

// ProductType is a custom type for WooCommerce product types
type ProductType string

// Define constants for each product type
const (
	Simple       ProductType = "simple"
	Grouped      ProductType = "grouped"
	External     ProductType = "external"
	Variable     ProductType = "variable"
	Variation    ProductType = "variation" // Adding Variation type for Variable Products
	Virtual      ProductType = "virtual"
	Downloadable ProductType = "downloadable"
	Subscription ProductType = "subscription"
	Composite    ProductType = "composite"
	Bundle       ProductType = "bundle"
)

// String returns the string representation of the ProductType
func (pt ProductType) String() string {
	return string(pt)
}
