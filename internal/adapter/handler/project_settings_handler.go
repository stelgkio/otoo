package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	ac "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/acs-courier"
	cu "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/courier4u"
	pn "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/notification"
	ps "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/project_secrets"
	wp "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/settings_general"
	tm "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/team"
	st "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/template"
	pw "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/settings/webhooks"
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
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, wp.SettingsGeneral(user, project, projectExtensions))
	}
	return r.Render(ctx, st.SettginsTemplate(user, project.Name, projectID, project, projectExtensions))

}

// ProjectSettingsSercrets GET /project/settings/secret/:projectId
func (ph *ProjectHandler) ProjectSettingsSercrets(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	return r.Render(ctx, ps.ProjectSecrets(project, projectExtensions))
}

// ProjectSettingsNotification GET /project/settings/notification/:projectId
func (ph *ProjectHandler) ProjectSettingsNotification(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	notifications, err := ph.notificationSvc.FindNotification(projectID, 10, 1, "timestamp", "", true)
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	return r.Render(ctx, pn.SettingsNotifications(project, notifications, projectExtensions))
}

// ProjectSettingsWebHook GET /project/settings/webhook/:projectId
func (ph *ProjectHandler) ProjectSettingsWebHook(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	return r.Render(ctx, pw.SettingsWebhooks(project, projectExtensions))
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
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	req := new(updateProjectRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, wp.SettingsGeneral(user, project, projectExtensions))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	project.WoocommerceProject.Name = req.Name
	project.WoocommerceProject.Description = req.Description
	project.Name = req.Name
	project.Description = req.Description

	proj, err := ph.svc.UpdateProject(ctx, project)
	return r.Render(ctx, wp.SettingsGeneral(user, proj, projectExtensions))
}

// ProjectDelete POST /project/settings/:projectId
func (ph *ProjectHandler) ProjectDelete(ctx echo.Context) error {

	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	projectID := ctx.Param("projectId")

	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	ph.webhookSvc.DeleteAllWebhooksByProjectID(projectID, project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain)

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
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	req := new(updateProjectRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, ps.ProjectSecrets(project, projectExtensions))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	project.WoocommerceProject.ConsumerKey = req.ConsumerKey
	project.WoocommerceProject.ConsumerSecret = req.ConsumerSecret

	_, err = ph.reportSvc.GetCustomerTotalCountTestCredential(ctx, req.ConsumerKey, req.ConsumerSecret, project.WoocommerceProject.Domain)
	if err != nil {
		slog.Error("Project new secrets error", "error", err)
		return r.Render(ctx, ps.ProjectSecretsError(project, projectExtensions))
	}
	proj, err := ph.svc.UpdateProject(ctx, project)

	return r.Render(ctx, ps.ProjectSecrets(proj, projectExtensions))
}

// ProjectSettingsTeam GET /project/settings/team/:projectId
func (ph *ProjectHandler) ProjectSettingsTeam(ctx echo.Context) error {
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
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, tm.Team(project, projectExtensions))
	}
	return r.Render(ctx, st.TeamTemplate(user, project.Name, projectID, project, projectExtensions))
}

// ProjectSettingsAcsCourier GET /project/settings/team/:projectId
func (ph *ProjectHandler) ProjectSettingsAcsCourier(ctx echo.Context) error {
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
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, ac.SettingsCourier4u(projectID, projectExtensions))
	}
	return r.Render(ctx, st.TeamTemplate(user, project.Name, projectID, project, projectExtensions))
}

// ProjectSettingsCourier4u GET /project/settings/team/:projectId
func (ph *ProjectHandler) ProjectSettingsCourier4u(ctx echo.Context) error {
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
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, cu.SettingsCourier4u(projectID, projectExtensions))
	}
	return r.Render(ctx, st.TeamTemplate(user, project.Name, projectID, project, projectExtensions))
}
