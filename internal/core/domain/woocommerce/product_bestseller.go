package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ProductBestSellerRecord struct {
	ID        		primitive.ObjectID  `bson:"_id,omitempty"`
	ProjectID 		string              `bson:"projectId"`	
	Error     		string              `bson:"error,omitempty"`
	Timestamp 		time.Time           `bson:"timestamp,omitempty"`
	ProductID 		int64               `bson:"productId,omitempty"`	
	TotalOrders		int64               `bson:"total_orders,omitempty"`
	TotalMoneyMade	float64             `bson:"total_money_spend,omitempty"`
	StartDate 		time.Time 		    `bson:"start_date" json:"start_date"`
	EndDate   		time.Time 		    `bson:"end_date" json:"end_date"`
		
}

func NewProductBestSellerRecord(projectID string, productID int64, totalOrders int64, startDate, endDate time.Time, totalMoneyMade float64) ProductBestSellerRecord {
	return ProductBestSellerRecord{
		ProjectID: projectID,
		ProductID: productID,
		TotalOrders: totalOrders,
		StartDate: startDate,
		EndDate: endDate,
		TotalMoneyMade: totalMoneyMade,
	}
}