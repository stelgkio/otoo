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
func (repo *ProjectRepository) CreateProject(ctx echo.Context, project *domain.Project, woo *domain.WoocommerceProject, shop *domain.ShopifyProject) (*domain.Project, error) {

	if woo != nil {
		_, err := repo.db.Model(woo).Insert()
		if err != nil {
			ctx.Echo().Logger.Error(err)
		}
		project.WoocommerceProjectId = woo.Id
		project.WoocommerceProject = woo
	}
	if shop != nil {
		_, err := repo.db.Model(shop).Insert()
		if err != nil {
			ctx.Echo().Logger.Error(err)
		}
		project.ShopifyProjectId = shop.Id
		project.ShopifyProject = shop
	}

	_, err := repo.db.Model(project).Insert()
	if err != nil {
		ctx.Echo().Logger.Error(err)
	}

	return project, nil
}
