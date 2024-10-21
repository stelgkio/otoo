package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	e "github.com/stelgkio/otoo/internal/core/util"
)

type ProjectService struct {
	repo           port.ProjectRepository
	wp             port.WoocommerceWebhookService
	wc             port.ProductService
	extensionSrv   port.ExtensionService
	userSvc        port.UserService
	userprojectSvc port.UserProjectService
}

// NewProjectService creates a new user service instance
func NewProjectService(repo port.ProjectRepository, wp port.WoocommerceWebhookService, wc port.ProductService, extensionSrv port.ExtensionService, userSvc port.UserService, userprojectSvc port.UserProjectService) *ProjectService {
	return &ProjectService{
		repo,
		wp,
		wc,
		extensionSrv,
		userSvc,
		userprojectSvc,
	}
}

// CreateProject creates a new Project
func (ps *ProjectService) CreateProject(ctx echo.Context, req *domain.ProjectRequest) (*domain.Project, error) {

	userID, err := auth.GetUserID(ctx)
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

	project.IsActive = true

	pr, err := ps.repo.CreateProject(ctx, project)
	if err != nil {
		return nil, errors.New("project is not created")
	}
	//TODO: get user and add projectId

	ps.userprojectSvc.AddUserToProject(ctx, userID, pr.Id)

	extension, err := ps.extensionSrv.GetExtensionByCode(ctx, domain.DataSynchronizerCode)
	if err != nil {
		return nil, err
	}
	ps.extensionSrv.CreateProjectExtension(ctx, project.Id.String(), extension, 370, "")

	extension2, err := ps.extensionSrv.GetExtensionByCode(ctx, domain.TeamMember)
	if err != nil {
		return nil, err
	}
	ps.extensionSrv.CreateProjectExtension(ctx, project.Id.String(), extension2, 370, "")

	go ps.wp.WoocommerceCreateAllWebHookAsync(req.ConsumerKey, req.ConsumerSecret, req.Domain, pr.Id.String())
	return pr, nil
}

// FindProjects finds projects in the database
func (ps *ProjectService) FindProjects(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error) {
	return ps.repo.FindProjects(ctx, filters, skip, limit)
}

// SearchByDomain finds projects by domain
func (ps *ProjectService) SearchByDomain(ctx echo.Context, filters *domain.FindProjectRequest, skip, limit int) ([]*domain.Project, error) {
	return ps.repo.SearchByDomain(ctx, filters, skip, limit)
}

// SoftDeleteProjects is doing a soft delete to this projects
func (ps *ProjectService) SoftDeleteProjects(ctx echo.Context, projectID uuid.UUID) error {
	err := ps.repo.DeleteProjectsByID(ctx, projectID)
	if err != nil && err != e.ErrDataNotFound {
		return e.ErrInternal
	}

	return nil
}

// GetProjectByID gets a project by its ID
func (ps *ProjectService) GetProjectByID(ctx echo.Context, id string) (*domain.Project, error) {
	return ps.repo.GetProjectByID(ctx, id)
}

// GetAllProjects returns all projects
func (ps *ProjectService) GetAllProjects() ([]*domain.Project, error) {
	return ps.repo.GetAllProjects()
}

// UpdateProject updates a project
func (ps *ProjectService) UpdateProject(ctx echo.Context, project *domain.Project) (*domain.Project, error) {
	user, err := ps.repo.UpdateProject(ctx, project)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return user, nil

}
