package handler

import (
	"github.com/labstack/echo/v4"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/default"
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

func (ph *DashboardHandler) ProjectDashboard(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	_, err := ph.projectSvc.GetProjectByID(ctx, projectId)
	if err != nil {
		return err
	}
	return util.Render(ctx, t.DeafultTemplate())
}
