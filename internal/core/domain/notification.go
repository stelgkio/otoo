package domain

// Notification represents a notification
type Notification struct {
	MongoBase
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Link        string `json:"link" bson:"link,omitempty"`
	IsActive    bool   `json:"is_active" bson:"is_active,omitempty"`
	IsRead      bool   `json:"is_read" bson:"is_read,omitempty"`
	IsDeleted   bool   `json:"is_deleted" bson:"is_deleted,omitempty"`
	UserID      string `json:"user_id" bson:"user_id,omitempty"`
	ProjectID   string `json:"project_id" bson:"project_id,omitempty"`
	Type        string `json:"type" bson:"type,omitempty"`
}
