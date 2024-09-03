package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	ps "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/project_secrets"
	wp "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/settings_general"
	"github.com/stelgkio/otoo/internal/core/auth"
	r "github.com/stelgkio/otoo/internal/core/util"
)

type updateProjectRequest struct {
	Name           string `form:"name"`
	Description    string `form:"description"`
	ConsumerKey    string `form:"consumer_key"`
	ConsumerSecret string `form:"consumer_secret"`
}

// ProjectSettings GET /project/settings/:projectId
func (ph *ProjectHandler) ProjectSettings(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	return r.Render(ctx, wp.SettingsGeneral(user, project))
}

// ProjectSettingsSercrets GET /project/settings/secret/:projectId
func (ph *ProjectHandler) ProjectSettingsSercrets(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	return r.Render(ctx, ps.ProjectSecrets(project))
}

// ProjectUpdate POST /project/settings/update/:projectId
func (ph *ProjectHandler) ProjectUpdate(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	req := new(updateProjectRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, wp.SettingsGeneral(user, project))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	project.WoocommerceProject.Name = req.Name
	project.WoocommerceProject.Description = req.Description
	project.Name = req.Name
	project.Description = req.Description

	proj, err := ph.svc.UpdateProject(ctx, project)
	return r.Render(ctx, wp.SettingsGeneral(user, proj))
}

// ProjectDelete POST /project/settings/:projectId
func (ph *ProjectHandler) ProjectDelete(ctx echo.Context) error {

	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}

	err = ph.svc.SoftDeleteProjects(ctx, userID)
	if err != nil {
		return err
	}
	ctx.Response().Header().Set("HX-Redirect", "/dashboard")
	return ctx.NoContent(http.StatusOK)

}

// ProjectSecretsUpdate POST /project/settings/secreate/update/:projectId
func (ph *ProjectHandler) ProjectSecretsUpdate(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}

	req := new(updateProjectRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, ps.ProjectSecrets(project))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	project.WoocommerceProject.ConsumerKey = req.ConsumerKey
	project.WoocommerceProject.ConsumerSecret = req.ConsumerSecret

	_, err = ph.reportSvc.GetCustomerTotalCountTestCredential(ctx, req.ConsumerKey, req.ConsumerSecret, project.WoocommerceProject.Domain)
	if err != nil {
		slog.Error("Project new secrets error", "error", err)
		return r.Render(ctx, ps.ProjectSecretsError(project))
	}
	proj, err := ph.svc.UpdateProject(ctx, project)

	return r.Render(ctx, ps.ProjectSecrets(proj))
}
