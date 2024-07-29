package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	c "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/customer/overview"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/default"
	o "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/order/overview"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/product/overview"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
)

// DashboardHandler handles the dashboard routes
type DashboardHandler struct {
	projectSvc  port.ProjectService
	userSvc     port.UserService
	customerSvc port.CustomerService
	productSvc  port.ProductService
	orderSvc    port.OrderService
	bestSeller  port.BestSellers
}

// NewDashboardHandler returns a new DashboardHandler
func NewDashboardHandler(projectSvc port.ProjectService, userSvc port.UserService, customerSvc port.CustomerService, productSvc port.ProductService, orderSvc port.OrderService, bestSeller port.BestSellers) *DashboardHandler {
	return &DashboardHandler{
		projectSvc:  projectSvc,
		userSvc:     userSvc,
		customerSvc: customerSvc,
		productSvc:  productSvc,
		orderSvc:    orderSvc,
		bestSeller:  bestSeller,
	}
}

// DefaultDashboard returns the default dashboard
func (dh *DashboardHandler) DefaultDashboard(ctx echo.Context) error {
	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(4)

	orderResults := make(chan int64, 1)
	orderErrors := make(chan error, 1)
	productResults := make(chan int64, 1)
	productErrors := make(chan error, 1)
	customerResults := make(chan int64, 1)
	customerErrors := make(chan error, 1)

	orderListResults := make(chan []*w.OrderRecord, 1)
	orderListErrors := make(chan error, 1)

	// Fetch order count
	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, w.OrderStatusCompleted, orderResults, orderErrors)
	}()

	// Fetch product count
	go func() {
		defer wg.Done()
		dh.productSvc.GetProductCount(ctx, projectID, productResults, productErrors)
	}()

	// Fetch customer count
	go func() {
		defer wg.Done()
		dh.customerSvc.GetCustomerCount(ctx, projectID, customerResults, customerErrors)
	}()

	// Fetch latest 10 order count
	go func() {
		defer wg.Done()
		dh.orderSvc.Get10LatestOrders(ctx, projectID, w.OrderStatusCompleted, orderListResults, orderListErrors)
	}()
	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(orderResults)
		close(orderErrors)
		close(productResults)
		close(productErrors)
		close(customerResults)
		close(customerErrors)
		close(orderListResults)
		close(orderListErrors)
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

	var orderList []*w.OrderRecord

	for count := range orderResults {
		orderCount = count
	}
	for count := range productResults {
		productCount = count
	}
	for count := range customerResults {
		customerCount = count
	}

	for item := range orderListResults {
		orderList = item
	}

	response := map[string]string{
		"order_count":    fmt.Sprintf("%d", orderCount),
		"product_count":  fmt.Sprintf("%d", productCount),
		"customer_count": fmt.Sprintf("%d", customerCount),
	}
	bestSeller, err := dh.bestSeller.FindBestSellers(projectID, 5, 1)
	if err != nil {
		return fmt.Errorf("bestSeller error: %v", err)
	}

	return util.Render(ctx, t.DeafultTemplate(user, project.Name, projectID, response, orderList, bestSeller))
}

// DefaultDashboardOverView returns the default dashboard
func (dh *DashboardHandler) DefaultDashboardOverView(ctx echo.Context) error {
	_, _, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(4)

	orderResults := make(chan int64, 1)
	orderErrors := make(chan error, 1)
	productResults := make(chan int64, 1)
	productErrors := make(chan error, 1)
	customerResults := make(chan int64, 1)
	customerErrors := make(chan error, 1)

	orderListResults := make(chan []*w.OrderRecord, 1)
	orderListErrors := make(chan error, 1)

	// Fetch order count
	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, w.OrderStatusCompleted, orderResults, orderErrors)
	}()

	// Fetch product count
	go func() {
		defer wg.Done()
		dh.productSvc.GetProductCount(ctx, projectID, productResults, productErrors)
	}()

	// Fetch customer count
	go func() {
		defer wg.Done()
		dh.customerSvc.GetCustomerCount(ctx, projectID, customerResults, customerErrors)
	}()

	// Fetch latest 10 order count
	go func() {
		defer wg.Done()
		dh.orderSvc.Get10LatestOrders(ctx, projectID, w.OrderStatusCompleted, orderListResults, orderListErrors)
	}()
	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(orderResults)
		close(orderErrors)
		close(productResults)
		close(productErrors)
		close(customerResults)
		close(customerErrors)
		close(orderListResults)
		close(orderListErrors)
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

	var orderList []*w.OrderRecord

	for count := range orderResults {
		orderCount = count
	}
	for count := range productResults {
		productCount = count
	}
	for count := range customerResults {
		customerCount = count
	}

	for item := range orderListResults {
		orderList = item
	}

	response := map[string]string{
		"order_count":    fmt.Sprintf("%d", orderCount),
		"product_count":  fmt.Sprintf("%d", productCount),
		"customer_count": fmt.Sprintf("%d", customerCount),
	}
	bestSeller, err := dh.bestSeller.FindBestSellers(projectID, 5, 1)
	if err != nil {
		return fmt.Errorf("bestSeller error: %v", err)
	}

	return util.Render(ctx, t.DeafultDashboard(projectID, response, orderList, bestSeller))
}

// CustomerDashboard returns the customer dashboard
func (dh *DashboardHandler) CustomerDashboard(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	return util.Render(ctx, c.CustomerOverView(projectID))
}

// ProductDashboard returns the product dashboard
func (dh *DashboardHandler) ProductDashboard(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	return util.Render(ctx, p.ProductOverview(projectID))
}

// OrderDashboard returns the order dashboard
func (dh *DashboardHandler) OrderDashboard(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	return util.Render(ctx, o.OrderOverView(projectID))
}

// OrderTable returns the order dashboard
func (dh *DashboardHandler) OrderTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	page := ctx.Param("page")
	status, err := w.StringToOrderStatus(ctx.Param("status"))
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid page number: %v", err))
	}

	var wg sync.WaitGroup
	// Fetch	var wg sync.WaitGroup
	wg.Add(2)

	orderCountChan := make(chan int64, 1)
	orderListChan := make(chan []*w.OrderRecord, 1)
	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, status, orderCountChan, errChan)
	}()

	// Fetch  10 orders
	go func() {
		defer wg.Done()
		dh.orderSvc.FindOrderByProjectIDAsync(projectID, 10, pageNum, status, orderListChan, errChan)
	}()
	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(orderCountChan)
		close(orderListChan)
		close(errChan)
	}()

	var totalItems int64
	var orderRecords []*w.OrderRecord

	select {
	case count := <-orderCountChan:
		totalItems = count
	case err := <-errChan:
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"failed to fetch order count": err.Error()})
	}

	select {
	case list := <-orderListChan:
		orderRecords = list
	case err := <-errChan:
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Convert orderRecords to OrderTableList for the response
	var orders []w.OrderTableList
	for _, record := range orderRecords {
		orders = append(orders, w.OrderTableList{
			ID:          record.ID,
			ProjectID:   record.ProjectID,
			Timestamp:   record.Timestamp,
			OrderID:     record.OrderID,
			TotalAmount: record.Order.Total,
			Status:      record.Status,
		})
	}

	// Prepare metadata
	itemsPerPage := 10
	totalPages := int(totalItems) / itemsPerPage
	if int(totalItems)%itemsPerPage > 0 {
		totalPages++
	}

	// Create response object
	response := w.OrderTableResponde{
		Data: orders,
		Meta: w.Meta{
			TotalItems:   int(totalItems),
			CurrentPage:  pageNum,
			ItemsPerPage: itemsPerPage,
			TotalPages:   totalPages,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetProjectAndUser retrieves project and user details from the context
func GetProjectAndUser(ctx echo.Context, dh *DashboardHandler) (*domain.Project, *domain.User, string, error) {
	// Extract project ID from the context
	projectID := ctx.Param("projectId")

	// Get project details using the project service
	project, err := dh.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, nil, "", err
	}

	// Get user ID from the authentication context
	userID, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, nil, "", err
	}

	// Get user details using the user service
	user, err := dh.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return nil, nil, "", err
	}

	// Return the project and user details
	return project, user, projectID, nil
}
