package repository

import (
	"github.com/go-pg/pg/v10"

	"github.com/labstack/echo/v4"
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
