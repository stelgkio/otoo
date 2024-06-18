package domain

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	Id        uuid.UUID  `json:"id" pg:"id,pk,type:uuid,default:gen_random_uuid()" bson:"_id,omitempty"`
	CreatedAt time.Time  `json:"created_at" pg:"created_at,default:now()" bson:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" pg:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" pg:"deleted_at" bson:"deleted_at,omitempty"`
	IsActive  bool       `pg:"is_active,default:true" bson:"is_active,omitempty"`
}

type MongoBase struct {
	CreatedAt time.Time  `bson:"created_at,omitempty"`
	UpdatedAt time.Time  `bson:"updated_at,omitempty"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
	IsActive  bool       `bson:"is_active,omitempty"`
}

func NewMongoBase() MongoBase {
	return MongoBase{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
}
