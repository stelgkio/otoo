package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Extension here we store all the available extensions we can add to the project
type Extension struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Title              string             `json:"title" bson:"title,omitempty"`
	Description        string             `json:"description" bson:"description,omitempty"`
	Code               string             `json:"code" bson:"code,omitempty"`
	Price              float64            `json:"price" bson:"price,omitempty"`
	PriceID            string             `json:"price_Id" bson:"price_Id,omitempty"`
	CreatedAt          time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt          time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt          time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive           bool               `json:"is_active" bson:"is_active,omitempty"`
	SubscriptionPeriod int                `json:"subscription_period" bson:"subscription_period,omitempty"`
}

// ProjectExtension represents active project Extension
type ProjectExtension struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Code        string             `json:"code" bson:"code,omitempty"`
	UserID      string             `json:"user_id" bson:"user_id,omitempty"`
	ProjectID   string             `json:"project_id" bson:"project_id,omitempty"`
	ExtensionID string             `json:"Extension_id" bson:"extension_id,omitempty"`
	CreatedAt   time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive    bool               `json:"is_active" bson:"is_active,omitempty"`
}

// AcsCourierExtension represents active acs details for acs courier
type AcsCourierExtension struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	ProjectID   string    `json:"project_id" bson:"project_id,omitempty"`
	ExtensionID string    `json:"extension_id" bson:"extension_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive    bool      `json:"is_active" bson:"is_active,omitempty"`

	CompanyID       string `json:"company_id" validate:"required" bson:"company_id,omitempty" form:"company_id"`
	CompanyPassword string `json:"company_password" validate:"required" bson:"company_password,omitempty" form:"company_password"`
	UserID          string `json:"user_id" validate:"required"  bson:"user_id,omitempty" form:"user_id"`
	UserPassword    string `json:"user_password" validate:"required" bson:"user_password,omitempty" form:"user_password"`
	AcsAPIKey       string `json:"acs_api_key" bson:"acs_api_key" form:"acs_api_key"`
	BillingCode     string `json:"billing_code" bson:"billing_code" form:"billing_code"`
}

// Validate validates the request body
func (p *AcsCourierExtension) Validate() map[string](string) {

	errors := make(map[string]string)

	if p.CompanyID == "" {
		errors["company_id"] = "Company Id is required"
	}
	if p.CompanyPassword == "" {
		errors["company_password"] = "Company Password is required"
	}
	if p.UserID == "" {
		errors["user_id"] = "User Id is required"
	}
	if p.UserPassword == "" {
		errors["user_password"] = "User Password is required"
	}
	if p.AcsAPIKey == "" {
		errors["acs_api_key"] = "AcsAPIKey is required"

	}
	if p.BillingCode == "" {
		errors["billing_code"] = "BillingCode is required"

	}

	return errors
}
