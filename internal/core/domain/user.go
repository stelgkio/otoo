package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/stelgkio/otoo/internal/core/util"
)

// UserRole is an enum for user's role
type UserRole string

// UserRole enum values
const (
	Admin      UserRole = "admin"
	Client     UserRole = "client"
	ClientUser UserRole = "client_user"
)

// User ...
type User struct {
	tableName struct{} `pg:"user,alias:user"`
	Base
	Name               string `json:"name" pg:"name,notnull"`
	Email              string `json:"email" pg:"email,unique,notnull"`
	Password           string `json:"password" pg:"password,notnull"`
	Role               UserRole
	ValidatedAt        time.Time
	LastLogin          time.Time
	LastName           string     `json:"last_name" pg:"last_name,notnull"`
	Projects           []*Project `pg:"many2many:user_projects"`
	ReseveNotification bool       `json:"reseve_notification" pg:"reseve_notification"`
}

// NewUser creates a instance of user with hashed password
func NewUser(email string, password string, name string, lastName string) (*User, error) {
	var err error
	u := new(User)
	var hash util.Hash
	u.Password, err = hash.Generate(password)
	if err != nil {
		return u, err
	}

	now := time.Now().UTC()
	u.Role = Client
	u.Email = email
	u.CreatedAt = now
	u.UpdatedAt = now
	u.Name = name
	u.LastName = lastName
	u.IsActive = true
	u.ReseveNotification = true
	return u, nil
}

// NewClientUser creates a instance of user with hashed password
func NewClientUser(email string, password string, name string, lastName string, role UserRole) (*User, error) {
	var err error
	u := new(User)
	var hash util.Hash
	u.Password, err = hash.Generate(password)
	if err != nil {
		return u, err
	}

	now := time.Now().UTC()
	u.Role = role
	u.Email = email
	u.CreatedAt = now
	u.UpdatedAt = now
	u.Name = name
	u.LastName = lastName
	u.IsActive = true
	u.ReseveNotification = true
	return u, nil
}

// ValidateEmail validates user's email and password
func (u *User) ValidateEmail(email string) error {

	return nil
}

// AddProject to user
func (u *User) AddProject(projectID uuid.UUID) {

}
func (pt UserRole) String() string {
	return string(pt)
}
