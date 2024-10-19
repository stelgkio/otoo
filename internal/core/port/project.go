package port

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// ProjectRepository is the interface for the project repository
type ProjectRepository interface {
	// CreateProject inserts a new project into the database
	CreateProject(ctx echo.Context, project *domain.Project) (*domain.Project, error)
	// // GetProjectByID selects a project by id
	GetProjectByID(ctx echo.Context, id string) (*domain.Project, error)
	// // GetProjectByDomain selects a project by domain
	GetProjectByDomain(ctx echo.Context, id string) (*domain.Project, error)
	// // ListProject selects a list of project with pagination
	FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error)
	SearchByDomain(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error)
	// // UpdateProject updates a project
	UpdateProject(ctx echo.Context, user *domain.Project) (*domain.Project, error)
	// // DeleteUser deletes a user
	// DeleteUser(ctx context.Context, id uint64) error
	DeleteProjectsByUserID(ctx echo.Context, userID uuid.UUID) error
	// 	GetAllProjects() ([]*domain.Project, error)
	GetAllProjects() ([]*domain.Project, error)
}

// ProjectService is an interface for interacting with user-related business logic
type ProjectService interface {
	// Register registers a new user
	CreateProject(ctx echo.Context, project *domain.ProjectRequest) (*domain.Project, error)
	// GetProject returns a project by id
	GetProjectByID(ctx echo.Context, id string) (*domain.Project, error)
	// // ListUsers returns a list of users with pagination
	FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error)
	SearchByDomain(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error)
	// // UpdateProject updates a project
	UpdateProject(ctx echo.Context, user *domain.Project) (*domain.Project, error)
	// // DeleteUser deletes a user
	// DeleteUser(ctx context.Context, id uint64) error
	SoftDeleteProjects(ctx echo.Context, userID uuid.UUID) error
	// GetAllProjects returns all projects
	GetAllProjects() ([]*domain.Project, error)
}

// UserProjectRepository interface
type UserProjectRepository interface {
	RemoveAllUsersFromProject(ctx echo.Context, projectID uuid.UUID) error
	FindUsersByProjectID(ctx echo.Context, projectID uuid.UUID) ([]*domain.User, error)
	FindProjectsByUserID(ctx echo.Context, userID uuid.UUID) ([]*domain.Project, error)
	RemoveUserFromProject(ctx echo.Context, userID uuid.UUID, projectID uuid.UUID) error
	AddUserToProject(ctx echo.Context, userID uuid.UUID, projectID uuid.UUID) error
}

// UserProjectService interface
type UserProjectService interface {
	RemoveAllUsersFromProject(ctx echo.Context, projectID uuid.UUID) error
	FindUsersByProjectID(ctx echo.Context, projectID uuid.UUID) ([]*domain.User, error)
	FindProjectsByUserID(ctx echo.Context, userID uuid.UUID) ([]*domain.Project, error)
	RemoveUserFromProject(ctx echo.Context, userID uuid.UUID, projectID uuid.UUID) error
	AddUserToProject(ctx echo.Context, userID uuid.UUID, projectID uuid.UUID) error
}
