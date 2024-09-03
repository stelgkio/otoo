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

// ProjectRepository is the repository for the Project model
type ProjectRepository struct {
	db *pg.DB
}

// NewProjectRepository returns a new ProjectRepository
func NewProjectRepository(db *pg.DB) *ProjectRepository {
	return &ProjectRepository{
		db,
	}
}

// CreateProject creates a newProject in the database
func (repo *ProjectRepository) CreateProject(ctx echo.Context, project *domain.Project) (*domain.Project, error) {

	_, err := repo.db.Model(project).Insert()
	if err != nil {
		return nil, err
	}

	return project, nil
}

// FindProjects finds projects in the database
func (repo *ProjectRepository) FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error) {

	var projects []*domain.Project

	userID, err := auth.GetUserID(ctx)
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
		Where("user_id =?", userID).
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

// DeleteProjectsByUserID is doing a soft delete to this projects
func (repo *ProjectRepository) DeleteProjectsByUserID(ctx echo.Context, userID uuid.UUID) error {
	project := &domain.Project{}
	res, err := repo.db.Model(project).
		Set("is_active = ?", false).
		Set("deleted_at = ?", time.Now().UTC()).
		Where("user_id = ?", userID).
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

// GetProjectByID returns a project by id
func (repo *ProjectRepository) GetProjectByID(ctx echo.Context, id string) (*domain.Project, error) {
	project := domain.Project{}
	err := repo.db.Model(&project).
		Where("is_active = ?", true).
		Where("id =?", id).
		Select()

	if err != nil {
		return nil, err
	}

	return &project, nil
}

// GetProjectByDomain returns a project by domaiL
func (repo *ProjectRepository) GetProjectByDomain(ctx echo.Context, domainURL string) (*domain.Project, error) {
	project := domain.Project{}
	err := repo.db.Model(&project).
		Where("is_active = ?", true).
		Where("woocommerce_domain =?", domainURL).
		First()

	if err != nil {
		return nil, err
	}

	return &project, nil
}

// GetAllProjects returns all projects
func (repo *ProjectRepository) GetAllProjects() ([]*domain.Project, error) {
	var projects []*domain.Project
	err := repo.db.Model(&projects).
		Where("is_active = ?", true).
		Select()
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// UpdateProject creates a new project in the database
func (repo *ProjectRepository) UpdateProject(ctx echo.Context, project *domain.Project) (*domain.Project, error) {
	_, err := repo.db.Model(project).Where("id =?", project.Id).Where("is_active =true").Update()
	if err != nil {
		panic(err)
	}

	return project, nil
}
