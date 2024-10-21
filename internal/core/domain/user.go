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
	Email              string `json:"email" pg:"email,notnull"`
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
func NewClientUser(email string, password string, name string, lastName string, role string, reseveNotification bool) (*User, error) {
	var err error
	u := new(User)
	var hash util.Hash
	u.Password, err = hash.Generate(password)
	if err != nil {
		return u, err
	}

	now := time.Now().UTC()
	u.Role = ReturnUserRoleFromWeb(role)
	u.Email = email
	u.CreatedAt = now
	u.UpdatedAt = now
	u.Name = name
	u.LastName = lastName
	u.IsActive = true
	u.ReseveNotification = reseveNotification
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

func ReturnUserRoleFromWeb(role string) UserRole {
	switch role {
	case "Admin":
		return Client
	case "User":
		return ClientUser
	case "client":
		return Client
	case "client_user":
		return ClientUser
	default:
		return ClientUser
	}
}

// ContainsUserID Function to check if any ProjectExtension contains the given ExtensionID
func ContainsUserID(users []*User, userID uuid.UUID) bool {
	for _, user := range users {
		if user.Id == userID {
			return true // Found a matching ExtensionID
		}
	}
	return false // No matching ExtensionID found
}

// UpdatePasswordRequest represents the request body for updating a user's password.
type UpdatePasswordRequest struct {
	CurrentPassword      string `form:"current-password" validate:"required"`
	Password             string `form:"password" validate:"required"`
	ConfirmationPassword string `form:"confirmation-password" validate:"required"`
}

// Validate validates the request body
func (p *UpdatePasswordRequest) Validate() map[string](string) {

	errors := make(map[string]string)

	if p.CurrentPassword == "" {
		errors["currentPassword"] = "CurrentPassword is required"
	}
	if p.Password == "" {
		errors["password"] = "Password is required"
	}
	if p.ConfirmationPassword == "" {
		errors["confirmationPassword"] = "ConfirmationPassword is required"
	}
	if p.Password != "" && p.ConfirmationPassword != "" {

		if p.Password != p.ConfirmationPassword {
			errors["password"] = "Password is not matching"
			errors["confirmationPassword"] = "Password is not matching"
		}
	}
	return errors
}
