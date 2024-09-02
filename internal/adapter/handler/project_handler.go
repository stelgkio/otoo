package handler

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	er "github.com/stelgkio/otoo/internal/adapter/web/view/component/error"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/create"
	l "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/list"
	syn "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/progress/synchronize"
	wp "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/progress/webhooks"
	v "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/validation"
	d "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/view"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	r "github.com/stelgkio/otoo/internal/core/util"
)

// ProjectHandler represents the HTTP handler for user-related requests
type ProjectHandler struct {
	svc           port.ProjectService
	userSvc       port.UserService
	reportSvc     port.ReportService
	productSvc    port.ProductService
	customerSvc   port.CustomerService
	orderSvc      port.OrderService
	bestSellerSvc port.ProductBestSellers
}

// NewProjectHandler creates a new ProjectHandler instance
func NewProjectHandler(svc port.ProjectService, userSvc port.UserService, reportSvc port.ReportService, productSvc port.ProductService, customerSvc port.CustomerService, orderSvc port.OrderService, bestSellerSvc port.ProductBestSellers) *ProjectHandler {
	return &ProjectHandler{
		svc,
		userSvc,
		reportSvc,
		productSvc,
		customerSvc,
		orderSvc,
		bestSellerSvc,
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
	return r.Render(ctx, p.ProjectCreateForm(false, nil, new(domain.ProjectRequest)))
}

// ProjectListPage  GET /project/list
func (ph *ProjectHandler) ProjectListPage(ctx echo.Context) error {

	projects, err := ph.svc.FindProjects(ctx, &domain.FindProjectRequest{}, 1, 10)
	if err != nil {
		return err
	}
	return r.Render(ctx, l.ProjectListPage(projects))
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
	return r.Render(ctx, wp.CheckWebhookProgress(user, projectID))
}

// ProjectSynchronize Synchronize otoo with eshop
// ProjectSettings GET /project/synchronize/:projectId
func (ph *ProjectHandler) ProjectSynchronize(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	_, err = ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	customerTotal, err := ph.reportSvc.GetCustomerTotalCount(ctx, projectID)
	productTotal, err := ph.reportSvc.GetProductTotalCount(ctx, projectID)
	orderTotal, err := ph.reportSvc.GetOrderTotalCount(ctx, projectID)

	return r.Render(ctx, syn.ProjectSynchronizer(user, projectID, customerTotal, productTotal, orderTotal))
}

// ProjectSynchronizeStart Synchronize otoo with eshop
// ProjectSettings POST /project/synchronize/:projectId
func (ph *ProjectHandler) ProjectSynchronizeStart(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	// Check which data options were selected
	customerSelected := ctx.FormValue("check_customer") == "on"
	productsSelected := ctx.FormValue("check_products") == "on"
	ordersSelected := ctx.FormValue("check_orders") == "on"
	customerTotal, err := strconv.ParseInt(ctx.Param("customerTotal"), 10, 64)
	productTotal, err := strconv.ParseInt(ctx.Param("productTotal"), 10, 64)
	orderTotal, err := strconv.ParseInt(ctx.Param("orderTotal"), 10, 64)

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
	// Perform actions based on the selected options
	if customerSelected {
		go ph.customerSvc.GetAllCustomerFromWoocommerce(project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain, projectID, customerTotal)
	} else {
		customerTotal = 0
	}

	if productsSelected {
		go ph.productSvc.GetAllProductFromWoocommerce(project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain, projectID, productTotal)
	} else {
		productTotal = 0
	}

	if ordersSelected {
		go ph.orderSvc.GetAllOrdersFromWoocommerce(project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain, projectID, productTotal)
	} else {
		orderTotal = 0
	}

	return r.Render(ctx, syn.ProjectSynchronizerStart(user, projectID, customerTotal, productTotal, orderTotal, 0.0, 0.0, 0.0))
}

// ProjectSynchronizeDone GET /project/synchronize/done/:projectId
func (ph *ProjectHandler) ProjectSynchronizeDone(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	customerTotal, err := strconv.ParseInt(ctx.Param("customerTotal"), 10, 64)
	productTotal, err := strconv.ParseInt(ctx.Param("productTotal"), 10, 64)
	orderTotal, err := strconv.ParseInt(ctx.Param("orderTotal"), 10, 64)

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
	_, err = ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if customerTotal == 0 && productTotal == 0 && orderTotal == 0 {
		ctx.Response().Header().Set("HX-Trigger", "done")
		return r.Render(ctx, syn.ProjectSynchronizerDone(user, projectID, customerTotal, productTotal, orderTotal, 100.0, 100.0, 100.0))
	}

	var wg sync.WaitGroup

	orderResults := make(chan int64, 1)
	orderErrors := make(chan error, 1)
	productResults := make(chan int64, 1)
	productErrors := make(chan error, 1)
	customerResults := make(chan int64, 1)
	customerErrors := make(chan error, 1)

	if customerTotal > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ph.customerSvc.GetCustomerCount(ctx, projectID, customerResults, customerErrors)
		}()
	}
	if productTotal > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ph.productSvc.GetProductCount(ctx, projectID, w.Variation, productResults, productErrors)
		}()
	}
	if orderTotal > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ph.orderSvc.GetOrderCountAsync(ctx, projectID, w.OrderStatusAll, "", orderResults, orderErrors)

		}()
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(orderResults)
		close(orderErrors)
		close(productResults)
		close(productErrors)
		close(customerResults)
		close(customerErrors)

	}()

	// Check if there were any errors
	// Check for errors
	for err := range orderErrors {
		if err != nil {
			return fmt.Errorf("order count error: %v", err)
		}
	}
	for err := range productErrors {
		if err != nil {
			return fmt.Errorf("product count error: %v", err)
		}
	}
	for err := range customerErrors {
		if err != nil {
			return fmt.Errorf("customer count error: %v", err)
		}
	}

	var orderCount, productCount, customerCount int64

	for count := range orderResults {
		orderCount = count
	}
	for count := range productResults {
		productCount = count
	}
	for count := range customerResults {
		customerCount = count
	}

	if productCount == productTotal && orderCount == orderTotal && (customerCount > customerTotal || customerCount == customerTotal) {
		ctx.Response().Header().Set("HX-Trigger", "done")
		ph.bestSellerSvc.RunAProductBestSellerInitializerJob(projectID)
		return r.Render(ctx, syn.ProjectSynchronizerDone(user, projectID, customerTotal, productTotal, orderTotal, 100.0, 100.0, 100.0))
	}

	return r.Render(ctx, syn.ProjectSynchronizerStart(user, projectID, customerTotal, productTotal, orderTotal, float64(customerCount), float64(productCount), float64(orderCount)))
}

// ProjectSynchronizeTest Synchronize otoo with eshop
// ProjectSettings GET /project/synchronize/:projectId
func (ph *ProjectHandler) ProjectSynchronizeTest(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	_, err = ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	customerTotal, err := ph.reportSvc.GetCustomerTotalCount(ctx, projectID)
	productTotal, err := ph.reportSvc.GetProductTotalCount(ctx, projectID)
	orderTotal, err := ph.reportSvc.GetOrderTotalCount(ctx, projectID)
	if err != nil {
		return err
	}

	return r.Render(ctx, syn.ProjectSynchronizerTest(user, projectID, customerTotal, productTotal, orderTotal))
}
