package handler

import (
	"fmt"
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

type DashboardHandler struct {
	projectSvc port.ProjectService
	userSvc    port.UserService
	customerSvc   port.CustomerService
	productSvc   port.ProductService
	orderSvc   port.OrderService
}

func NewDashboardHandler(projectSvc port.ProjectService, userSvc port.UserService,customerSvc port.CustomerService,productSvc port.ProductService,orderSvc port.OrderService) *DashboardHandler {
	return &DashboardHandler{
		projectSvc: projectSvc,
		userSvc: userSvc,
		customerSvc: customerSvc,
		productSvc: productSvc,
		orderSvc: orderSvc,
	}
}

func (dh *DashboardHandler) DefaultDashboard(ctx echo.Context) error {
	project, user,projectId,err:= GetProjectAndUser(ctx, dh)
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
		dh.orderSvc.GetOrderCount(ctx,projectId,orderResults, orderErrors)
	}()

	// Fetch product count
	go func() {
		defer wg.Done()
		dh.productSvc.GetProductCount(ctx,projectId,productResults, productErrors)
	}()

	// Fetch customer count
	go func() {
		defer wg.Done()
		dh.customerSvc.GetCustomerCount(ctx,projectId,customerResults, customerErrors)
	}()


	// Fetch latest 10 order count
	go func() {
		defer wg.Done()
		dh.orderSvc.Get10LatestOrders(ctx,projectId,orderListResults,orderListErrors)
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
		"order_count":   fmt.Sprintf("%d", orderCount),
		"product_count":   fmt.Sprintf("%d",productCount),
		"customer_count":  fmt.Sprintf("%d",customerCount),
	}


	return util.Render(ctx, t.DeafultTemplate(user,project.Name,projectId,response,orderList))
}
func (dh *DashboardHandler) DefaultDashboardOverView(ctx echo.Context) error {
	_, _,projectId,err:= GetProjectAndUser(ctx, dh)
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
		dh.orderSvc.GetOrderCount(ctx,projectId,orderResults, orderErrors)
	}()

	// Fetch product count
	go func() {
		defer wg.Done()
		dh.productSvc.GetProductCount(ctx,projectId,productResults, productErrors)
	}()

	// Fetch customer count
	go func() {
		defer wg.Done()
		dh.customerSvc.GetCustomerCount(ctx,projectId,customerResults, customerErrors)
	}()


	// Fetch latest 10 order count
	go func() {
		defer wg.Done()
		dh.orderSvc.Get10LatestOrders(ctx,projectId,orderListResults,orderListErrors)
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
		"order_count":   fmt.Sprintf("%d", orderCount),
		"product_count":   fmt.Sprintf("%d",productCount),
		"customer_count":  fmt.Sprintf("%d",customerCount),
	}

	return util.Render(ctx, t.DeafultDashboard(projectId,response,orderList))
}


func (dh *DashboardHandler) CustomerDashboard(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	return util.Render(ctx, c.CustomerOverView(projectId))
}


func (dh *DashboardHandler) ProductDashboard(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	return util.Render(ctx, p.ProductOverview(projectId))
}


func (dh *DashboardHandler) OrderDashboard(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	
	return util.Render(ctx, o.OrderOverView(projectId))
}


func GetProjectAndUser(ctx echo.Context, dh *DashboardHandler) (*domain.Project, *domain.User, string, error) {
	// Extract project ID from the context
	projectId := ctx.Param("projectId")
	
	// Get project details using the project service
	project, err := dh.projectSvc.GetProjectByID(ctx, projectId)
	if err != nil {
		return nil, nil,"", err
	}

	// Get user ID from the authentication context
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return nil, nil,"", err
	}

	// Get user details using the user service
	user, err := dh.userSvc.GetUserById(ctx, userId)
	if err != nil {
		return nil, nil,"", err
	}

	// Return the project and user details
	return project, user, projectId, nil
}