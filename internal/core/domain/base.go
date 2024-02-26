package domain

import (
	"time"
)

type Base struct {
	ID        uint64     `json:"id" pg:"id,pk"`
	CreatedAt time.Time  `json:"created_at" pg:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" pg:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" pg:"deleted_at"`
}
