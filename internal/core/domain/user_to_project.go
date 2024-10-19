package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserProject struct {
	tableName struct{}   `pg:"user_projects,alias:user_projects"`
	ID        uuid.UUID  `json:"id" pg:"id,pk,type:uuid,default:gen_random_uuid()" bson:"_id,omitempty"`
	CreatedAt time.Time  `json:"created_at" pg:"created_at,default:now()" bson:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at" pg:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" pg:"deleted_at" bson:"deleted_at,omitempty"`
	IsActive  bool       `pg:"is_active,default:true" bson:"is_active,omitempty"`

	// UserID and ProjectID are not nullable
	UserID    uuid.UUID `pg:"user_id,notnull,type:uuid," json:"user_id"`
	ProjectID uuid.UUID `pg:"project_id,notnull,type:uuid," json:"project_id"`

	// Adding a unique constraint on the combination of UserID and ProjectID
	_ struct{} `pg:",unique:user_project_unique" bson:"-"` // Unique constraint on user_id and project_id
}

// NewUserProject creates and returns a new UserProject instance
func NewUserProject(userID, projectID uuid.UUID) *UserProject {
	return &UserProject{
		ID:        uuid.New(),       // Generate a new UUID for the ID
		CreatedAt: time.Now().UTC(), // Set the current timestamp for CreatedAt
		UpdatedAt: time.Now().UTC(), // Set the current timestamp for UpdatedAt
		IsActive:  true,             // Set IsActive to true by default
		UserID:    userID,           // Initialize with provided UserID
		ProjectID: projectID,        // Initialize with provided ProjectID
	}
}

// -- Create user_projects table
// CREATE TABLE user_projects (
//     id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
//     created_at timestamptz DEFAULT now(),
//     updated_at timestamptz,
//     deleted_at timestamptz,
//     is_active boolean DEFAULT true,
//     user_id uuid NOT NULL,
//     project_id uuid NOT NULL
// );

// -- Add unique constraint to user_id and project_id
// ALTER TABLE user_projects
// ADD CONSTRAINT unique_user_project UNIQUE (user_id, project_id);
