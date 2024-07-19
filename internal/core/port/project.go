package port

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

type ProjectRepository interface {
	// CreateProject inserts a new project into the database
	CreateProject(ctx echo.Context, project *domain.Project) (*domain.Project, error)
	// // GetProjectByID selects a project by id
	GetProjectByID(ctx echo.Context, id string) (*domain.Project, error)
	// // GetProjectByDomain selects a project by domain
	GetProjectByDomain(ctx echo.Context, id string) (*domain.Project, error)
	// // ListProject selects a list of project with pagination
	FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error)
	// // UpdateUser updates a user
	// UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// // DeleteUser deletes a user
	// DeleteUser(ctx context.Context, id uint64) error
	DeleteProjectsByUserId(ctx echo.Context, userId uuid.UUID) error
}

// UserService is an interface for interacting with user-related business logic
type ProjectService interface {
	// Register registers a new user
	CreateProject(ctx echo.Context, project *domain.ProjectRequest) (*domain.Project, error)
	// GetProject returns a project by id
	GetProjectByID(ctx echo.Context, id string) (*domain.Project, error)
	// // ListUsers returns a list of users with pagination
	FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error)
	// // UpdateUser updates a user
	// UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// // DeleteUser deletes a user
	// DeleteUser(ctx context.Context, id uint64) error
	SoftDeleteProjects(ctx echo.Context, userId uuid.UUID) error
}
