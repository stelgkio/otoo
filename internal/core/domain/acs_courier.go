package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AcsCourier represents the structure for storing ACS Courier credentials in MongoDB
type AcsCourier struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`                            // MongoDB ObjectID
	ProjectID       string             `json:"project_id" bson:"project_id,omitempty"`   // Associated project ID
	CompanyID       string             `json:"company_id" bson:"company_id"`             // Company ID provided by ACS
	CompanyPassword string             `json:"company_password" bson:"company_password"` // Company password provided by ACS
	UserID          string             `json:"user_id" bson:"user_id"`                   // User ID provided by ACS
	UserPassword    string             `json:"user_password" bson:"user_password"`       // User password provided by ACS
	AcsAPIKey       string             `json:"acs_api_key" bson:"acs_api_key"`           // ACS API key
}
