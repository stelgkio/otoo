package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	ac "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/acs-courier"
	cu "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/courier4u"
	pn "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/notification"
	pm "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/payment"
	ps "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/project_secrets"
	wp "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/settings_general"
	tm "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/team"
	st "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/template"
	pw "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/webhooks"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/util"
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
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
	}
	return r.Render(ctx, ps.ProjectSecrets(project, projectExtensions, user))
}

// ProjectSettingsNotification GET /project/settings/notification/:projectId
func (ph *ProjectHandler) ProjectSettingsNotification(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
	}
	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	notifications, err := ph.notificationSvc.FindNotification(projectID, 10, 1, "timestamp", "", true)
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	return r.Render(ctx, pn.SettingsNotifications(project, notifications, projectExtensions, user))
}

// ProjectSettingsWebHook GET /project/settings/webhook/:projectId
func (ph *ProjectHandler) ProjectSettingsWebHook(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
	}
	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	return r.Render(ctx, pw.SettingsWebhooks(project, projectExtensions, user))
}

// ProjectSettingsWebHook GET /project/settings/webhook/:projectId
func (ph *ProjectHandler) ProjectSettingsPayment(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
	}
	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	projectExtensions, err := ph.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	return r.Render(ctx, pm.SettingsPayments(project, projectExtensions, user))
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

	projectID := ctx.Param("projectId")

	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	ph.webhookSvc.DeleteAllWebhooksByProjectID(projectID, project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain)

	err = ph.svc.SoftDeleteProjects(ctx, project.Id)
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("HX-Redirect", "/dashboard")

	return ctx.NoContent(http.StatusOK)

}

// ProjectSecretsUpdate POST /project/settings/secreate/update/:projectId
func (ph *ProjectHandler) ProjectSecretsUpdate(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
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
		return r.Render(ctx, ps.ProjectSecrets(project, projectExtensions, user))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	project.WoocommerceProject.ConsumerKey = req.ConsumerKey
	project.WoocommerceProject.ConsumerSecret = req.ConsumerSecret

	_, err = ph.reportSvc.GetCustomerTotalCountTestCredential(ctx, req.ConsumerKey, req.ConsumerSecret, project.WoocommerceProject.Domain)
	if err != nil {
		slog.Error("Project new secrets error", "error", err)
		return r.Render(ctx, ps.ProjectSecretsError(project, projectExtensions, user))
	}
	proj, err := ph.svc.UpdateProject(ctx, project)

	return r.Render(ctx, ps.ProjectSecrets(proj, projectExtensions, user))
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
		return r.Render(ctx, tm.Team(project, projectExtensions, user))
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
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.AcsCode
	})
	acs := new(domain.AcsCourierExtension)
	if len(extension) > 0 {

		acs, err = ph.extensionSvc.GetACSProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
		if err != nil {
			return err
		}
		if acs == nil {
			acs = new(domain.AcsCourierExtension)
		}

	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, ac.SettingsACSCourier(projectID, projectExtensions, user, make(map[string]string), acs))
	}
	return r.Render(ctx, st.AcsCourierTemplate(user, project.Name, projectID, project, projectExtensions, make(map[string]string), acs))
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
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.Courier4u
	})
	courier4u := new(domain.Courier4uExtension)
	if len(extension) > 0 {

		courier4u, err = ph.extensionSvc.GetCourier4uProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
		if err != nil {
			return err
		}
		if courier4u == nil {
			courier4u = new(domain.Courier4uExtension)
		}

	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, cu.SettingsCourier4u(projectID, projectExtensions, user, make(map[string]string), courier4u))
	}
	return r.Render(ctx, st.Courier4uTemplate(user, project.Name, projectID, project, projectExtensions, make(map[string]string), courier4u))
}
