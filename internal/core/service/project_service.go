package service

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
)

type ProjectService struct {
	repo port.ProjectRepository
	wp   port.WoocommerceService
}

// NewUserService creates a new user service instance
func NewProjectService(repo port.ProjectRepository, wp port.WoocommerceService) *ProjectService {
	return &ProjectService{
		repo,
		wp,
	}
}

func (ps *ProjectService) CreateProject(ctx echo.Context, req *domain.ProjectRequest) (*domain.Project, error) {

	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, errors.New("user is not found")
	}

	project, err := domain.NewProject(req)
	var woo domain.WoocommerceProject
	var shop domain.ShopifyProject

	if err != nil {
		return nil, errors.New("project is not created")
	}
	if req.ProjectType == domain.Woocommerce {
		woo, err = domain.NewWoocommerceProject(req)
		if err != nil {
			return nil, errors.New("woocommerce project is not created")
		}
		project.WoocommerceProject = woo
		project.ShopifyProject = domain.ShopifyProject{}
	}
	if req.ProjectType == domain.Shopify {
		shop, err = domain.NewShopifyProject(req)
		if err != nil {
			return nil, errors.New("shopify project is not created")
		}
		project.ShopifyProject = shop
		project.WoocommerceProject = domain.WoocommerceProject{}
	}
	project.UserId = userId
	pr, err := ps.repo.CreateProject(ctx, project)
	if err != nil {
		return nil, errors.New("project is not created")
	}

	go ps.wp.WoocommerceCreateOrderWebHook(req.ConsumerKey, req.ConsumerSecret, req.Domain, pr.Id.String())

	return pr, nil
}
