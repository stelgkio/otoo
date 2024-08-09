package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Extention represents available extention
type Extention struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Code        string             `json:"code" bson:"code,omitempty"`
	Price       float64            `json:"price" bson:"price,omitempty"`
	CreatedAt   time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive    bool               `json:"is_active" bson:"is_active,omitempty"`
}

// ProjectExtention represents active project extention
type ProjectExtention struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Code        string             `json:"code" bson:"code,omitempty"`
	UserID      string             `json:"user_id" bson:"user_id,omitempty"`
	ProjectID   string             `json:"project_id" bson:"project_id,omitempty"`
	ExtentionID string             `json:"extention_id" bson:"extention_id,omitempty"`
	CreatedAt   time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive    bool               `json:"is_active" bson:"is_active,omitempty"`
}

// AcsCourierExtention represents active acs details for acs courier
type AcsCourierExtention struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Code        string             `json:"code" bson:"code,omitempty"`
	UserID      string             `json:"user_id" bson:"user_id,omitempty"`
	ProjectID   string             `json:"project_id" bson:"project_id,omitempty"`
	ExtentionID string             `json:"extention_id" bson:"extention_id,omitempty"`
	CreatedAt   time.Time          `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at"  bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `json:"deleted_at"  bson:"deleted_at,omitempty"`
	IsActive    bool               `json:"is_active" bson:"is_active,omitempty"`
}
