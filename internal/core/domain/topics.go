package domain

import (
	"github.com/google/uuid"
)

type WooCommerce_Topic struct {
	Base
	TopicName   string    `json:"name" pg:"topic_name,notnull"`
	Description string    `json:"description" pg:"description,notnull"`
	ProjectId   uuid.UUID `json:"project_id" pg:"fk:project_id,type:uuid"`
	Project     *Project  `pg:"rel:has-one"`
	UserId      uuid.UUID `pg:"fk:user_id,type:uuid"`
	User        *User     `pg:"rel:has-one"`
}

type Shopify_Topic struct {
	Base
	TopicName   string    `json:"name" pg:"topic_name,notnull"`
	Description string    `json:"description" pg:"description,notnull"`
	ProjectId   uuid.UUID `json:"project_id" pg:"fk:project_id,type:uuid"`
	Project     *Project  `pg:"rel:has-one"`
	UserId      uuid.UUID `pg:"fk:user_id,type:uuid"`
	User        *User     `pg:"rel:has-one"`
}
