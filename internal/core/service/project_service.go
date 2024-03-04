package service

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
)

type ProjectService struct {
	repo port.ProjectRepository
}

// NewUserService creates a new user service instance
func NewProjectService(repo port.ProjectRepository) *ProjectService {
	return &ProjectService{
		repo,
	}
}

func (ps *ProjectService) CreateProject(ctx echo.Context, req *domain.ProjectRequest) (*domain.Project, error) {
	project, err := domain.NewProject(req)
	var woo *domain.WoocommerceProject
	var shop *domain.ShopifyProject

	if err != nil {
		return nil, errors.New("project is not created")
	}
	if req.ProjectType == domain.Woocommerce {
		woo, err = domain.NewWoocommerceProject(req)
		if err != nil {
			return nil, ctx.String(http.StatusInternalServerError, "woocommerce project is not created")
		}
	}
	if req.ProjectType == domain.Shopify {
		shop, err = domain.NewShopifyProject(req)
		if err != nil {
			return nil, errors.New("shopify project is not created")
		}
	}
	return ps.repo.CreateProject(ctx, project, woo, shop)
}
