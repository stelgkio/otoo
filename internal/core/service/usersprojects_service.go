package service

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	e "github.com/stelgkio/otoo/internal/core/util"
)

// UsersProjectsService implements port.UserService interface
type UsersProjectsService struct {
	repo port.UserProjectRepository
}

// NewUsersProjectsService creates a new user service instance
func NewUsersProjectsService(repo port.UserProjectRepository) *UsersProjectsService {
	return &UsersProjectsService{
		repo,
	}
}

// RemoveAllUsersFromProject removes all users from a project
func (up *UsersProjectsService) RemoveAllUsersFromProject(ctx echo.Context, projectID uuid.UUID) error {
	err := up.repo.RemoveAllUsersFromProject(ctx, projectID)
	if err != nil && err != e.ErrDataNotFound {
		return e.ErrInternal
	}
	return nil
}

// FindUsersByProjectID  retrieves all projects associated with a given user
func (up *UsersProjectsService) FindUsersByProjectID(ctx echo.Context, projectID uuid.UUID) ([]*domain.User, error) {
	user, err := up.repo.FindUsersByProjectID(ctx, projectID)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return user, nil
}

// FindProjectsByUserID  retrieves all projects associated with a given user
func (up *UsersProjectsService) FindProjectsByUserID(ctx echo.Context, userID uuid.UUID) ([]*domain.Project, error) {
	projects, err := up.repo.FindProjectsByUserID(ctx, userID)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return projects, nil
}

// RemoveUserFromProject deletes the association between a user and a project
func (up *UsersProjectsService) RemoveUserFromProject(ctx echo.Context, userID uuid.UUID, projectID uuid.UUID) error {
	err := up.repo.RemoveUserFromProject(ctx, userID, projectID)
	if err != nil && err != e.ErrDataNotFound {
		return e.ErrInternal
	}
	return nil
}

// AddUserToProject creates a new relationship between a user and a project
func (up *UsersProjectsService) AddUserToProject(ctx echo.Context, userID uuid.UUID, projectID uuid.UUID) error {
	err := up.repo.AddUserToProject(ctx, userID, projectID)
	if err != nil && err != e.ErrDataNotFound {
		return e.ErrInternal
	}
	return nil
}
