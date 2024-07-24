package handler

import (
	"github.com/labstack/echo/v4"
	c "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/customer/overview"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/default"
	o "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/order/overview"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/product/overview"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
)

type DashboardHandler struct {
	projectSvc port.ProjectService
	userSvc    port.UserService
}

func NewDashboardHandler(projectSvc port.ProjectService, userSvc port.UserService) *DashboardHandler {
	return &DashboardHandler{
		projectSvc: projectSvc,
		userSvc:    userSvc,
	}
}

func (ph *DashboardHandler) DefaultDashboard(ctx echo.Context) error {
	project, user,projectId,err:= GetProjectAndUser(ctx, ph)
	if err != nil {
		return err
	}
	



	return util.Render(ctx, t.DeafultTemplate(user,project.Name,projectId))
}
func (ph *DashboardHandler) DefaultDashboardOverView(ctx echo.Context) error {
	_, _,projectId,err:= GetProjectAndUser(ctx, ph)
	if err != nil {
		return err
	}
	return util.Render(ctx, t.DeafultDashboard(projectId))
}


func (ph *DashboardHandler) CustomerDashboard(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	return util.Render(ctx, c.CustomerOverView(projectId))
}


func (ph *DashboardHandler) ProductDashboard(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	return util.Render(ctx, p.ProductOverview(projectId))
}


func (ph *DashboardHandler) OrderDashboard(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	
	return util.Render(ctx, o.OrderOverView(projectId))
}


func GetProjectAndUser(ctx echo.Context, ph *DashboardHandler) (*domain.Project, *domain.User, string, error) {
	// Extract project ID from the context
	projectId := ctx.Param("projectId")
	
	// Get project details using the project service
	project, err := ph.projectSvc.GetProjectByID(ctx, projectId)
	if err != nil {
		return nil, nil,"", err
	}

	// Get user ID from the authentication context
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, nil,"", err
	}

	// Get user details using the user service
	user, err := ph.userSvc.GetUserById(ctx, userId)
	if err != nil {
		return nil, nil,"", err
	}

	// Return the project and user details
	return project, user, projectId, nil
}