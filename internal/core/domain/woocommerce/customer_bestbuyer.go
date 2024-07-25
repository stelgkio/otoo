package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type CustomerBestBuyerRecord struct {
	ID        		primitive.ObjectID  `bson:"_id,omitempty"`
	ProjectID 		string              `bson:"projectId"`	
	Error     		string              `bson:"error,omitempty"`
	Timestamp 		time.Time           `bson:"timestamp,omitempty"`
	Email     		string              `bson:"email,omitempty"`		
	TotalOrders		int64               `bson:"total_orders,omitempty"`
	TotalMoneySpend	float64             `bson:"total_money_spend,omitempty"`
	StartDate 		time.Time 		    `bson:"start_date" json:"start_date"`
	EndDate   		time.Time 		    `bson:"end_date" json:"end_date"`
		
}

func NewCustomerBestBuyerRecordRecord(projectID string, Email string, totalOrders int64, startDate, endDate time.Time,totalMoneySpend float64) CustomerBestBuyerRecord {
	return CustomerBestBuyerRecord{
		ProjectID: projectID,
		Email: Email,
		TotalOrders: totalOrders,
		StartDate: startDate,
		EndDate: endDate,
		TotalMoneySpend: totalMoneySpend,
	}
}