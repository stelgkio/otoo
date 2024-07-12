package repository

import (
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
)

type ProjectRepository struct {
	db *pg.DB
}

func NewProjectRepository(db *pg.DB) *ProjectRepository {
	return &ProjectRepository{
		db,
	}
}

// CreatProject creates a newProject in the database
func (repo *ProjectRepository) CreateProject(ctx echo.Context, project *domain.Project) (*domain.Project, error) {

	_, err := repo.db.Model(project).Insert()
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (repo *ProjectRepository) FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error) {

	var projects []*domain.Project

	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, errors.New("user is not found")
	}
	offset := (skip - 1) * limit

	query := repo.db.Model(&projects)
	if filters.Name != "" {
		query = query.Where("name = ?", filters.Name)
	}
	if filters.Domain != "" {
		query = query.WhereOr("shopify_domain =?", filters.Domain)
		query = query.WhereOr("woocommerce_domain =?", filters.Domain)
	}
	query = query.
		Where("user_id =?", userId).
		Where("is_active =true").
		Order("name ASC").
		Limit(limit).
		Offset(offset)

	err = query.Select()
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// DeleteProjectsByUserId is doing a soft delete to this projects
func (repo *ProjectRepository) DeleteProjectsByUserId(ctx echo.Context, userId uuid.UUID) error {
	project := &domain.Project{}
	res, err := repo.db.Model(project).
		Set("is_active = ?", false).
		Set("deleted_at = ?", time.Now()).
		Where("user_id = ?", userId).
		//Returning("*"). // This ensures the updated project is returned
		Update()
	if res.RowsAffected() == 0 {
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}
