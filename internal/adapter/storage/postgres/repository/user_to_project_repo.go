package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// UserProjectRepository handles the many-to-many relationships between users and projects
type UserProjectRepository struct {
	db *pg.DB
}

// NewUserProjectRepository returns a new instance of UserProjectRepository
func NewUserProjectRepository(db *pg.DB) *UserProjectRepository {
	return &UserProjectRepository{
		db: db,
	}
}

// AddUserToProject creates a new relationship between a user and a project
func (repo *UserProjectRepository) AddUserToProject(ctx echo.Context, userID uuid.UUID, projectID uuid.UUID) error {
	userProject := domain.NewUserProject(userID, projectID)

	_, err := repo.db.Model(userProject).Insert()
	if err != nil {
		return err
	}

	return nil
}

// RemoveUserFromProject deletes the association between a user and a project
func (repo *UserProjectRepository) RemoveUserFromProject(ctx echo.Context, userID uuid.UUID, projectID uuid.UUID) error {
	_, err := repo.db.Model((*domain.UserProject)(nil)).
		Where("user_id = ?", userID).
		Where("project_id = ?", projectID).
		Delete()

	if err != nil {
		return err
	}

	return nil
}

// FindProjectsByUserID retrieves all projects associated with a given user
func (repo *UserProjectRepository) FindProjectsByUserID(ctx echo.Context, userID uuid.UUID) ([]*domain.Project, error) {
	var projects []*domain.Project

	err := repo.db.Model(&projects).
		Join("JOIN user_projects ON user_projects.project_id = project.id").
		Where("user_projects.user_id = ?", userID).
		Where("project.is_active = true").
		Select()

	if err != nil {
		return nil, err
	}

	return projects, nil
}

// FindUsersByProjectID retrieves all users associated with a given project
func (repo *UserProjectRepository) FindUsersByProjectID(ctx echo.Context, projectID uuid.UUID) ([]*domain.User, error) {
	var users []*domain.User

	err := repo.db.Model(&users).
		Join("JOIN user_projects ON user_projects.user_id = user.id").
		Where("user_projects.project_id = ?", projectID).
		Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}

// RemoveAllUsersFromProject deletes all user associations for a given project
func (repo *UserProjectRepository) RemoveAllUsersFromProject(ctx echo.Context, projectID uuid.UUID) error {
	_, err := repo.db.Model((*domain.UserProject)(nil)).
		Where("project_id = ?", projectID).
		Delete()

	if err != nil {
		return err
	}

	return nil
}
