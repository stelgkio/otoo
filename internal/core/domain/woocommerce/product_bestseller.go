package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductBestSellerRecord represents a product record
type ProductBestSellerRecord struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID       string             `bson:"projectId"`
	Error           string             `bson:"error,omitempty"`
	Timestamp       time.Time          `bson:"timestamp,omitempty"`
	ProductName     string             `bson:"product_name,omitempty"`
	ProductID       int64              `bson:"productId,omitempty"`
	TotalOrders     int64              `bson:"total_orders,omitempty"`
	TotalOrdersRate float64            `bson:"total_orders_rate,omitempty"`
	TotalMoneyMade  float64            `bson:"total_money_spend,omitempty"`
	StartDate       time.Time          `bson:"start_date" json:"start_date"`
	EndDate         time.Time          `bson:"end_date" json:"end_date"`
	CreatedAt       time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt       time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive        bool               `json:"is_active" bson:"is_active,omitempty"`
}

// NewProductBestSellerRecord creates a new ProductBestSellerRecord instance
func NewProductBestSellerRecord(projectID, productName string, productID int64, totalOrders int64, startDate, endDate time.Time, totalMoneyMade float64) ProductBestSellerRecord {
	return ProductBestSellerRecord{
		ProjectID:      projectID,
		ProductID:      productID,
		TotalOrders:    totalOrders,
		StartDate:      startDate,
		EndDate:        endDate,
		TotalMoneyMade: totalMoneyMade,
		ProductName:    productName,
		IsActive:       true,
		Timestamp:      time.Now(),
		CreatedAt:      time.Time{},
		DeletedAt:      time.Time{},
	}
}

// CalculatePercentages calculates the percentage of orders in the period
func (a *ProductBestSellerRecord) CalculatePercentages(totatOrdersCount int64) {
	if a.TotalOrders > 0 {
		a.TotalOrdersRate = (float64(a.TotalOrders) / float64(totatOrdersCount)) * 100
	} else {
		a.TotalOrdersRate = 0
	}
}
