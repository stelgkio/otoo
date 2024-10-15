package domain

import (
	"time"

	"github.com/google/uuid"
)

// Base represents the base model for all models
type Base struct {
	Id        uuid.UUID  `json:"id" pg:"id,pk,type:uuid,default:gen_random_uuid()" bson:"_id,omitempty"`
	CreatedAt time.Time  `json:"created_at" pg:"created_at,default:now()" bson:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" pg:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" pg:"deleted_at" bson:"deleted_at,omitempty"`
	IsActive  bool       `pg:"is_active,default:true" bson:"is_active,omitempty"`
}

// MongoBase creates a new base model
type MongoBase struct {
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
	DeletedAt time.Time `bson:"deleted_at,omitempty"`
	IsActive  bool      `bson:"is_active,omitempty"`
}

// NewMongoBase creates a new base model
func NewMongoBase() MongoBase {
	return MongoBase{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		IsActive:  true,
	}
}

type Meta struct {
	TotalItems   int `json:"totalItems"`
	CurrentPage  int `json:"currentPage"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalPages   int `json:"totalPages"`
}
