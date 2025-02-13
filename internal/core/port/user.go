package port

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// UserRepository ...
type UserRepository interface {
	// CreateUser inserts a new user into the database
	CreateUser(ctx echo.Context, user *domain.User) (*domain.User, error)
	//GetUserByID selects a user by id
	GetUserById(ctx echo.Context, id uuid.UUID) (*domain.User, error)
	// GetUser returns a user by projectid
	GetAdminUserByProjectId(ctx echo.Context, id uuid.UUID) ([]*domain.User, error)
	// // GetUserByEmail selects a user by email
	GetUserByEmail(ctx echo.Context, email string) (*domain.User, error)
	// // ListUsers selects a list of users with pagination
	// ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error)
	// // UpdateUser updates a user
	UpdateUser(ctx echo.Context, user *domain.User) (*domain.User, error)
	// // DeleteUser deletes a user
	DeleteUser(ctx echo.Context, id uuid.UUID) error
	FindUsersByProjectId(ctx echo.Context, id uuid.UUID) ([]*domain.User, error)
	FindProjectsByUserId(ctx echo.Context, userId uuid.UUID) ([]*domain.Project, error)
}

// UserService is an interface for interacting with user-related business logic
type UserService interface {
	// Register registers a new user
	CreateUser(ctx echo.Context, user *domain.User) (*domain.User, error)
	// GetUser returns a user by id
	GetUserById(ctx echo.Context, id uuid.UUID) (*domain.User, error)
	// GetUser returns a user by projectid
	GetAdminUserByProjectId(ctx echo.Context, id uuid.UUID) ([]*domain.User, error)
	// GetUser returns a user by id
	GetUserByEmail(ctx echo.Context, email string) (*domain.User, error)
	// // ListUsers returns a list of users with pagination
	// ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error)
	// // UpdateUser updates a user
	UpdateUser(ctx echo.Context, user *domain.User) (*domain.User, error)
	// // DeleteUser deletes a user
	DeleteUser(ctx echo.Context, id uuid.UUID) error
	FindUsersByProjectId(ctx echo.Context, id uuid.UUID) ([]*domain.User, error)
	FindProjectsByUserId(ctx echo.Context, userId uuid.UUID) ([]*domain.Project, error)
}
