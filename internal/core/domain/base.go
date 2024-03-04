package domain

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	Id        uuid.UUID  `json:"id" pg:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"created_at" pg:"created_at,default:now()"`
	UpdatedAt time.Time  `json:"updated_at" pg:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" pg:"deleted_at"`
}
