package domain

import (
	"time"

	"github.com/go-pg/pg/types"
	"github.com/stelgkio/otoo/internal/core/util"
)

type DtoUser struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"password"`
}

// UserRole is an enum for user's role
type UserRole string

// UserRole enum values
const (
	Admin  UserRole = "admin"
	Client UserRole = "client"
)

type User struct {
	Base
	Name        string
	Email       string
	Password    string
	Role        UserRole
	ValidatedAt types.NullTime
	LastLogin   types.NullTime
	IsActive    bool
}

// NewUser creates a instance of user with hashed password
func NewUser(email string, password string) (*User, error) {
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
	u.IsActive = true
	return u, nil
}
