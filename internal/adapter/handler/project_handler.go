package handler

import (
	"log/slog"
	"strings"

	"github.com/labstack/echo/v4"
	er "github.com/stelgkio/otoo/internal/adapter/web/view/component/error"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/create"
	l "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/list"
	wp "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/progress/webhooks"
	v "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/validation"
	d "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/view"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	r "github.com/stelgkio/otoo/internal/core/util"
)

// UserHandler represents the HTTP handler for user-related requests
type ProjectHandler struct {
	svc     port.ProjectService
	userSvc port.UserService
}

// NewProjectHandler creates a new ProjectHandler instance
func NewProjectHandler(svc port.ProjectService, userSvc port.UserService) *ProjectHandler {
	return &ProjectHandler{
		svc,
		userSvc,
	}
}

// POST /project/create
func (ph *ProjectHandler) CreateProject(ctx echo.Context) error {
	req := new(domain.ProjectRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Create project binding error", "error", err)
		return r.Render(ctx, p.ProjectCreateForm(true, nil, new(domain.ProjectRequest)))
	}
	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return r.Render(ctx, p.ProjectCreateForm(true, validationErrors, req))

	}

	dom, err := ph.svc.CreateProject(ctx, req)
	if err != nil {
		slog.Error("Create project error", "error", err)
		return r.Render(ctx, p.ProjectCreateForm(true, nil, new(domain.ProjectRequest)))
	}

	slog.Info("Create new project", "log_info", dom)
	return r.Render(ctx, wp.WebHooksProgress(dom.Id.String(), nil))
}

// GET /project/createform
func (ph *ProjectHandler) ProjectCreateForm(ctx echo.Context) error {
	return r.Render(ctx, p.ProjectCreateForm(false, nil, new(domain.ProjectRequest)))
}

// GET /project/list
func (ph *ProjectHandler) ProjectListPage(ctx echo.Context) error {

	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{}, 1, 10)
	if err != nil {
		return err
	}
	return r.Render(ctx, l.ProjectListPage(projects))
}

// GET /dashboard
func (ph *ProjectHandler) GetProjectDashboardPage(ctx echo.Context) error {

	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{}, 1, 10)
	if err != nil {
		return err
	}
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	return r.Render(ctx, d.ProjectDashboard(projects, user))
}

// POST /project/validation/name
func (ph *ProjectHandler) ProjectNameValidation(ctx echo.Context) error {
	req := new(domain.ProjectRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding request", "error", err)
	}
	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{Name: req.Name}, 1, 10)
	if err != nil {
		return err
	}
	var valid bool = true
	if len(projects) > 0 {
		valid = false
	}
	return r.Render(ctx, v.ProjectNameValidation(valid, req.Name))
}

// POST /project/validation/domain
func (ph *ProjectHandler) ProjectDomainValidation(ctx echo.Context) error {
	errors := make(map[string]string)
	errors["domain"] = ""
	req := new(domain.ProjectRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding request", "error", err)
		return r.Render(ctx, er.ErrorPage())
	}
	validationErrors := req.Validate()
	if validationErrors["domain"] != "" {
		return r.Render(ctx, v.DomainUrlValidation(false, req.Domain, validationErrors))

	}
	trimmedDomain := strings.TrimRight(req.Domain, "/")
	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{Domain: trimmedDomain}, 1, 10)
	if err != nil {
		return err
	}

	var valid bool = true
	if len(projects) > 0 {
		valid = false
		errors["domain"] = "Domain url already exist!"
	}

	return r.Render(ctx, v.DomainUrlValidation(valid, req.Domain, errors))
}

// GET /project/webhooks/:projectId
func (ph *ProjectHandler) CheckWebHooks(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	return r.Render(ctx, wp.CheckWebhookProgress(user, projectId))
}
