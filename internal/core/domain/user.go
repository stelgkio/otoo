package domain

import (
	"time"

	"github.com/go-pg/pg/types"
	"github.com/stelgkio/otoo/internal/core/util"
)

// UserRole is an enum for user's role
type UserRole string

// UserRole enum values
const (
	Admin  UserRole = "admin"
	Client UserRole = "client"
)

type User struct {
	Base
	Name        string `json:"name" pg:"name,notnull"`
	Email       string `json:"email" pg:"email,unique,notnull"`
	Password    string `json:"password" pg:"password,notnull"`
	Role        UserRole
	ValidatedAt types.NullTime
	LastLogin   types.NullTime
	IsActive    bool   `json:"is_active" pg:"is_active"`
	LastName    string `json:"last_name" pg:"last_name,notnull"`
}

// NewUser creates a instance of user with hashed password
func NewUser(email string, password string, name string, last_name string) (*User, error) {
	var err error
	u := new(User)
	var hash util.Hash
	u.Password, err = hash.Generate(password)
	if err != nil {
		return u, err
	}

	now := time.Now()
	u.Role = Client
	u.Email = email
	u.CreatedAt = now
	u.UpdatedAt = now
	u.Name = name
	u.LastName = last_name
	u.IsActive = false
	return u, nil
}

// Validate validates user's email and password
func (u *User) ValidateEmail(email string) error {

	return nil
}
