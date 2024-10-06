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

// ProjectHandler represents the HTTP handler for user-related requests
type ProjectHandler struct {
	svc               port.ProjectService
	userSvc           port.UserService
	reportSvc         port.ReportService
	productSvc        port.ProductService
	customerSvc       port.CustomerService
	orderSvc          port.OrderService
	bestSellerSvc     port.ProductBestSellers
	notificationSvc   port.NotificationService
	extensionSvc      port.ExtensionService
	webhookSvc        port.WoocommerceWebhookService
	orderAnalyticsSvc port.OrderAnalyticsCron
}

// NewProjectHandler creates a new ProjectHandler instance
func NewProjectHandler(
	svc port.ProjectService,
	userSvc port.UserService,
	reportSvc port.ReportService,
	productSvc port.ProductService,
	customerSvc port.CustomerService,
	orderSvc port.OrderService,
	bestSellerSvc port.ProductBestSellers,
	notificationSvc port.NotificationService,
	extensionSvc port.ExtensionService,
	webhookSvc port.WoocommerceWebhookService,
	orderAnalyticsSvc port.OrderAnalyticsCron) *ProjectHandler {
	return &ProjectHandler{
		svc,
		userSvc,
		reportSvc,
		productSvc,
		customerSvc,
		orderSvc,
		bestSellerSvc,
		notificationSvc,
		extensionSvc,
		webhookSvc,
		orderAnalyticsSvc,
	}
}

// CreateProject POST /project/create
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
	validationErrors, err := ph.DomainValidation(ctx)
	validationErrors, err = ph.KeyValidation(ctx)
	validationErrors, err = ph.NameValidation(ctx)
	if len(validationErrors) > 0 {
		return r.Render(ctx, p.ProjectCreateForm(true, validationErrors, req))

	}
	if err != nil {
		slog.Error("Create project error", "error", err)
		return r.Render(ctx, p.ProjectCreateForm(true, nil, new(domain.ProjectRequest)))
	}

	_, err = ph.reportSvc.GetCustomerTotalCountTestCredential(ctx, req.ConsumerKey, req.ConsumerSecret, req.Domain)
	if err != nil {
		slog.Error("Create project error", "error", err)
		validationErrors["consumer_key"] = "Consumer Key is invalid"
		validationErrors["consumer_secret"] = "Consumer Secret is invalid "
		validationErrors["domain"] = "Domain is not available"
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

// ProjectCreateForm GET /project/createform
func (ph *ProjectHandler) ProjectCreateForm(ctx echo.Context) error {
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, p.ProjectCreateForm(false, nil, new(domain.ProjectRequest)))
	}
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	return r.Render(ctx, p.CreateProjectTemplate(user, false, nil, new(domain.ProjectRequest)))
}

// ProjectListPage  GET /project/list
func (ph *ProjectHandler) ProjectListPage(ctx echo.Context) error {

	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{}, 1, 10)
	if err != nil {
		return err
	}
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, l.ProjectListPage(projects))
	}
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	return r.Render(ctx, d.ProjectDashboard(projects, user))

}

// GetProjectDashboardPage GET /dashboard
func (ph *ProjectHandler) GetProjectDashboardPage(ctx echo.Context) error {

	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{}, 1, 10)
	if err != nil {
		return err
	}
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}

	return r.Render(ctx, d.ProjectDashboard(projects, user))
}

// ProjectNameValidation POST /project/validation/name
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

// ProjectDomainValidation POST /project/validation/domain
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

// ProjectKeyValidation POST /project/validation/key
func (ph *ProjectHandler) ProjectKeyValidation(ctx echo.Context) error {
	errors := make(map[string]string)

	req := new(domain.ProjectRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding request", "error", err)
		return r.Render(ctx, er.ErrorPage())
	}
	validationErrors := req.Validate()
	if validationErrors["consumer_key"] != "" || validationErrors["consumer_secret"] != "" {
		return r.Render(ctx, v.DomainKeyValidation(false, req.ConsumerKey, req.ConsumerSecret, validationErrors))

	}

	return r.Render(ctx, v.DomainKeyValidation(true, req.ConsumerKey, req.ConsumerSecret, errors))
}

// DomainValidation validate domain
func (ph *ProjectHandler) DomainValidation(ctx echo.Context) (map[string]string, error) {
	errors := make(map[string]string)

	req := new(domain.ProjectRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding request", "error", err)
		return errors, err
	}
	validationErrors := req.Validate()
	if validationErrors["domain"] != "" {
		return errors, nil

	}
	trimmedDomain := strings.TrimRight(req.Domain, "/")
	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{Domain: trimmedDomain}, 1, 10)
	if err != nil {
		return errors, nil
	}
	if len(projects) > 0 {
		errors["domain"] = "Domain url already exist!"
	}

	return errors, nil
}

// NameValidation name
func (ph *ProjectHandler) NameValidation(ctx echo.Context) (map[string]string, error) {
	errors := make(map[string]string)

	req := new(domain.ProjectRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding request", "error", err)
		return errors, err
	}
	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{Name: req.Name}, 1, 10)
	if err != nil {
		return errors, nil
	}

	if len(projects) > 0 {
		errors["name"] = "Project name already exists!"
	}
	return errors, nil
}

// KeyValidation validate consumer keys
func (ph *ProjectHandler) KeyValidation(ctx echo.Context) (map[string]string, error) {
	errors := make(map[string]string)

	req := new(domain.ProjectRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding request", "error", err)
		return errors, err
	}
	validationErrors := req.Validate()
	if validationErrors["consumer_key"] != "" || validationErrors["consumer_secret"] != "" {
		return errors, nil

	}

	return errors, nil
}

// CheckWebHooks GET /project/webhooks/:projectId
func (ph *ProjectHandler) CheckWebHooks(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	return r.Render(ctx, wp.CheckWebhookProgress(user, projectID))
}
