package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

type ProjectRepository interface {
	// CreateProject inserts a new project into the database
	CreateProject(ctx echo.Context, project *domain.Project) (*domain.Project, error)
	// // GetUserByID selects a user by id
	// GetUserByID(ctx context.Context, id uint64) (*domain.User, error)
	// // GetUserByEmail selects a user by email
	//GetUserByEmail(ctx echo.Context, email string) (*domain.User, error)
	// // ListProject selects a list of project with pagination
	FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error)
	// // UpdateUser updates a user
	// UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// // DeleteUser deletes a user
	// DeleteUser(ctx context.Context, id uint64) error
}

// UserService is an interface for interacting with user-related business logic
type ProjectService interface {
	// Register registers a new user
	CreateProject(ctx echo.Context, project *domain.ProjectRequest) (*domain.Project, error)
	// GetUser returns a user by id
	// GetUser(ctx context.Context, id uint64) (*domain.User, error)
	// // ListUsers returns a list of users with pagination
	FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error)
	// // UpdateUser updates a user
	// UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// // DeleteUser deletes a user
	// DeleteUser(ctx context.Context, id uint64) error
}
