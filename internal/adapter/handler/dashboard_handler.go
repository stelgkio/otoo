package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	c "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/customer/overview"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/default"
	chart "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/order/charts"
	o "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/order/overview"
	tab "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/order/table"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/product/overview"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
)

// DashboardHandler handles the dashboard routes
type DashboardHandler struct {
	projectSvc      port.ProjectService
	userSvc         port.UserService
	customerSvc     port.CustomerService
	productSvc      port.ProductService
	orderSvc        port.OrderService
	analyticsRepo   port.AnalyticsRepository
	extensionSvc    port.ExtensionService
	notificationSvc port.NotificationService
	voucherSvc      port.VoucherService
	paymentSvc      port.PaymentService
	hermesSvc       port.HermesService
}

// NewDashboardHandler returns a new DashboardHandler
func NewDashboardHandler(
	projectSvc port.ProjectService,
	userSvc port.UserService,
	customerSvc port.CustomerService,
	productSvc port.ProductService,
	orderSvc port.OrderService,
	analyticsRepo port.AnalyticsRepository,
	extension port.ExtensionService,
	notification port.NotificationService,
	voucherSvc port.VoucherService,
	paymentSvc port.PaymentService,
	hermesSvc port.HermesService) *DashboardHandler {
	return &DashboardHandler{
		projectSvc:      projectSvc,
		userSvc:         userSvc,
		customerSvc:     customerSvc,
		productSvc:      productSvc,
		orderSvc:        orderSvc,
		analyticsRepo:   analyticsRepo,
		extensionSvc:    extension,
		notificationSvc: notification,
		voucherSvc:      voucherSvc,
		paymentSvc:      paymentSvc,
		hermesSvc:       hermesSvc,
	}
}

// DefaultDashboard returns the default dashboard
func (dh *DashboardHandler) DefaultDashboard(ctx echo.Context) error {
	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(5)

	orderResults := make(chan int64, 1)
	orderErrors := make(chan error, 1)
	productResults := make(chan int64, 1)
	productErrors := make(chan error, 1)
	customerResults := make(chan int64, 1)
	customerErrors := make(chan error, 1)

	orderListResults := make(chan []*w.OrderRecord, 1)
	orderListErrors := make(chan error, 1)

	orderWeeklyBalance := make(chan *w.WeeklyAnalytics, 1)
	orderWeeklyBalanceErrors := make(chan error, 1)

	// Fetch order count
	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountWithDeleteAsync(ctx, projectID, w.OrderStatusCompleted, "", orderResults, orderErrors)
	}()
	// Fetch order count
	go func() {
		defer wg.Done()
		dh.orderSvc.GetLatestOrderWeeklyBalance(nil, projectID, orderWeeklyBalance, orderWeeklyBalanceErrors)
	}()
	// Fetch product count
	go func() {
		defer wg.Done()
		dh.productSvc.GetProductCount(ctx, projectID, w.Variation, productResults, productErrors)
	}()

	// Fetch customer count
	go func() {
		defer wg.Done()
		dh.customerSvc.GetCustomerCount(ctx, projectID, customerResults, customerErrors)
	}()

	// Fetch latest 10 order count
	go func() {
		defer wg.Done()
		dh.orderSvc.Get10LatestOrders(ctx, projectID, w.OrderStatusCompleted, "orderId", orderListResults, orderListErrors)
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
		close(orderWeeklyBalance)
		close(orderWeeklyBalanceErrors)
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
	for err := range orderWeeklyBalanceErrors {
		if err != nil {
			return fmt.Errorf("weekBalance find error: %v", err)
		}
	}

	var orderCount, productCount, customerCount int64

	var weeklyBalance *w.WeeklyAnalytics

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
	for item := range orderWeeklyBalance {
		weeklyBalance = item
	}

	response := map[string]string{
		"order_count":    fmt.Sprintf("%d", orderCount),
		"product_count":  fmt.Sprintf("%d", productCount),
		"customer_count": fmt.Sprintf("%d", customerCount),
	}
	bestSeller, err := dh.analyticsRepo.FindBestSellers(projectID, 5, 1)
	if err != nil {
		return fmt.Errorf("bestSeller error: %v", err)
	}

	return util.Render(ctx, t.DeafultTemplate(user, project.Name, projectID, response, orderList, bestSeller, weeklyBalance))
}

// DefaultDashboardOverView returns the default dashboard
func (dh *DashboardHandler) DefaultDashboardOverView(ctx echo.Context) error {
	_, _, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(5)

	orderResults := make(chan int64, 1)
	orderErrors := make(chan error, 1)
	productResults := make(chan int64, 1)
	productErrors := make(chan error, 1)
	customerResults := make(chan int64, 1)
	customerErrors := make(chan error, 1)

	orderListResults := make(chan []*w.OrderRecord, 1)
	orderListErrors := make(chan error, 1)
	orderWeeklyBalance := make(chan *w.WeeklyAnalytics, 1)
	orderWeeklyBalanceErrors := make(chan error, 1)

	// Fetch order count
	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, w.OrderStatusCompleted, "", orderResults, orderErrors)
	}()
	// Fetch order count
	go func() {
		defer wg.Done()
		dh.orderSvc.GetLatestOrderWeeklyBalance(nil, projectID, orderWeeklyBalance, orderWeeklyBalanceErrors)
	}()
	// Fetch product count
	go func() {
		defer wg.Done()
		dh.productSvc.GetProductCount(ctx, projectID, w.Variation, productResults, productErrors)
	}()

	// Fetch customer count
	go func() {
		defer wg.Done()
		dh.customerSvc.GetCustomerCount(ctx, projectID, customerResults, customerErrors)
	}()

	// Fetch latest 10 order count
	go func() {
		defer wg.Done()
		dh.orderSvc.Get10LatestOrders(ctx, projectID, w.OrderStatusCompleted, "orderId", orderListResults, orderListErrors)
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
		close(orderWeeklyBalance)
		close(orderWeeklyBalanceErrors)
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
	for err := range orderWeeklyBalanceErrors {
		if err != nil {
			return fmt.Errorf("weekBalance find error: %v", err)
		}
	}

	var orderCount, productCount, customerCount int64

	var orderList []*w.OrderRecord
	var weeklyBalance *w.WeeklyAnalytics
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
	for item := range orderWeeklyBalance {
		if item != nil {
			weeklyBalance = item
		}
	}

	response := map[string]string{
		"order_count":    fmt.Sprintf("%d", orderCount),
		"product_count":  fmt.Sprintf("%d", productCount),
		"customer_count": fmt.Sprintf("%d", customerCount),
	}
	bestSeller, err := dh.analyticsRepo.FindBestSellers(projectID, 5, 1)
	if err != nil {
		return fmt.Errorf("bestSeller error: %v", err)
	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		return util.Render(ctx, t.DeafultDashboard(projectID, response, orderList, bestSeller, weeklyBalance))
	}
	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	return util.Render(ctx, t.DeafultTemplate(user, project.Name, projectID, response, orderList, bestSeller, weeklyBalance))
}

// CustomerDashboard returns the customer dashboard
func (dh *DashboardHandler) CustomerDashboard(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return util.Render(ctx, c.CustomerOverView(projectID))
	}
	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return util.Render(ctx, c.CustomerTemplate(user, project.Name, projectID))
}

// CustomerTable returns the order custmer list
func (dh *DashboardHandler) CustomerTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	page := ctx.Param("page")
	pageNum, err := strconv.Atoi(page)
	sort := ctx.QueryParam("sort")
	direction := ctx.QueryParam("direction")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid page number: %v", err))
	}

	if sort == "" {
		sort = "customerId"
	}
	if direction == "" {
		direction = "asc"
	}

	var wg sync.WaitGroup
	// Fetch	var wg sync.WaitGroup
	wg.Add(2)

	customerCountChan := make(chan int64, 1)
	customerListChan := make(chan []*w.CustomerRecord, 1)
	errChan := make(chan error, 1)
	errListChan := make(chan error, 1)

	go func() {
		defer wg.Done()
		dh.customerSvc.GetCustomerCount(ctx, projectID, customerCountChan, errChan)
	}()

	// Fetch  10 customers
	go func() {
		defer wg.Done()
		dh.customerSvc.FindCustomerByProjectIDAsync(projectID, 10, pageNum, sort, direction, customerListChan, errListChan)
	}()
	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(customerCountChan)
		close(customerListChan)
		close(errChan)
		close(errListChan)
	}()

	var totalItems int64
	var customerRecords []*w.CustomerRecord

	for err := range errChan {
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"failed to fetch customer count": err.Error()})
		}
	}
	for errList := range errListChan {
		if errList != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error to fetch customer": errList.Error()})
		}
	}

	for item := range customerCountChan {
		totalItems = item
	}
	for item := range customerListChan {
		customerRecords = item
	}

	// Convert customerRecords to customerTableList for the response
	var customers []w.CustomerTableList
	if customerRecords != nil {
		for _, record := range customerRecords {
			totalSpent := 0.0
			if record.Orders != nil || len(record.Orders) > 0 {
				for _, order := range record.Orders {
					order, err := dh.orderSvc.GetOrderByID(projectID, order)
					if err != nil {
						totalSpent += 0.0
						continue
					}
					floatValue, err := strconv.ParseFloat(order.Order.Total, 64)
					if err != nil {
						totalSpent += 0.0
						continue
					}
					totalSpent += floatValue
				}
			}
			fullName := record.Customer.FirstName + " " + record.Customer.LastName
			if strings.Trim(fullName, " ") == "" {
				fullName = record.Customer.Billing.FirstName + " " + record.Customer.Billing.LastName
			}

			customers = append(customers, w.CustomerTableList{
				ID:          record.ID,
				Name:        fullName,
				Email:       record.Email,
				TotalOrders: len(record.Orders),
				TotalSpent:  fmt.Sprintf("%.2f", totalSpent),
			})
		}
	}

	// Prepare metadata
	itemsPerPage := 10
	totalPages := int(totalItems) / itemsPerPage
	if int(totalItems)%itemsPerPage > 0 {
		totalPages++
	}

	// Create response object
	response := w.CustomerTableResponde{
		Data: customers,
		Meta: w.Meta{
			TotalItems:   int(totalItems),
			CurrentPage:  pageNum,
			ItemsPerPage: itemsPerPage,
			TotalPages:   totalPages,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}

// ProductDashboard returns the product dashboard
func (dh *DashboardHandler) ProductDashboard(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return util.Render(ctx, p.ProductOverview(projectID))
	}
	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return util.Render(ctx, p.ProductTemplate(user, project.Name, projectID))

}

// ProductTable returns the order product list
func (dh *DashboardHandler) ProductTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	page := ctx.Param("page")
	pageNum, err := strconv.Atoi(page)
	sort := ctx.QueryParam("sort")
	direction := ctx.QueryParam("direction")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid page number: %v", err))
	}

	if sort == "" {
		sort = "productId"
	}
	if direction == "" {
		direction = "asc"
	}

	var wg sync.WaitGroup
	// Fetch	var wg sync.WaitGroup
	wg.Add(2)

	productCountChan := make(chan int64, 1)
	productListChan := make(chan []*w.ProductRecord, 1)
	errChan := make(chan error, 1)
	errListChan := make(chan error, 1)

	go func() {
		defer wg.Done()
		dh.productSvc.GetProductCount(ctx, projectID, w.Variation, productCountChan, errChan)
	}()

	// Fetch  10 products
	go func() {
		defer wg.Done()
		dh.productSvc.FindProductByProjectIDAsync(projectID, 10, pageNum, sort, direction, w.Variation, productListChan, errListChan)
	}()
	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(productCountChan)
		close(productListChan)
		close(errChan)
		close(errListChan)
	}()

	var totalItems int64
	var productRecords []*w.ProductRecord

	for err := range errChan {
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"failed to fetch product count": err.Error()})
		}
	}
	for errList := range errListChan {
		if errList != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error to fetch product": errList.Error()})
		}
	}

	for item := range productCountChan {
		totalItems = item
	}
	for item := range productListChan {
		productRecords = item
	}

	// Convert productRecords to productTableList for the response
	var products []w.ProductTableList
	if productRecords != nil {
		for _, record := range productRecords {
			var imageURL string
			if len(record.Product.Images) > 0 {
				imageURL = record.Product.Images[0].Src
			} else {
				imageURL = "" // or a default image URL
			}
			var categoriesName string
			if len(record.Product.Categories) > 0 {
				categoriesName = record.Product.Categories[0].Name
			} else {
				categoriesName = "" // or a default image URL
			}
			products = append(products, w.ProductTableList{
				ID:               record.ID,
				ProjectID:        record.ProjectID,
				Timestamp:        record.Timestamp,
				ProductID:        record.ProductID,
				TotalAmountSpend: "",
				TotalOrders:      len(record.Orders),
				SKU:              record.Product.Sku,
				Price:            record.Product.Price,
				Category:         categoriesName,
				ProductImageURL:  imageURL,
				Name:             record.Product.Name,
				ProductType:      record.Product.Type,
			})
		}
	}

	// Prepare metadata
	itemsPerPage := 10
	totalPages := int(totalItems) / itemsPerPage
	if int(totalItems)%itemsPerPage > 0 {
		totalPages++
	}

	// Create response object
	response := w.ProductTableResponde{
		Data: products,
		Meta: w.Meta{
			TotalItems:   int(totalItems),
			CurrentPage:  pageNum,
			ItemsPerPage: itemsPerPage,
			TotalPages:   totalPages,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}

// OrderDashboard returns the order dashboard
func (dh *DashboardHandler) OrderDashboard(ctx echo.Context) error {

	projectID := ctx.Param("projectId")
	var wg sync.WaitGroup
	// Fetch	var wg sync.WaitGroup
	wg.Add(4)

	totalCountChan := make(chan int64, 1)
	count24hChan := make(chan int64, 1)
	count7dChan := make(chan int64, 1)
	count1mChan := make(chan int64, 1)
	errChan := make(chan error, 4)

	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountWithDeleteAsync(ctx, projectID, w.OrderStatusCompleted, "", totalCountChan, errChan)
	}()
	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, w.OrderStatusCompleted, "24h", count24hChan, errChan)
	}()
	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, w.OrderStatusCompleted, "7d", count7dChan, errChan)
	}()
	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, w.OrderStatusCompleted, "1m", count1mChan, errChan)
	}()

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(totalCountChan)
		close(count24hChan)
		close(count7dChan)
		close(count1mChan)
		close(errChan)
	}()
	var totalItems int64
	var total24hItems int64
	var total7hItems int64
	var total1mItems int64

	select {
	case countAll := <-totalCountChan:
		totalItems = countAll

	}
	select {
	case count24hChan := <-count24hChan:
		total24hItems = count24hChan

	}
	select {
	case count7 := <-count7dChan:
		total7hItems = count7

	}
	select {
	case count30 := <-count1mChan:
		total1mItems = count30

	}
	select {
	case err := <-errChan:
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"failed to fetch order count": err.Error()})
		}
	}
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return util.Render(ctx, o.OrderOverView(projectID, fmt.Sprintf("%d", totalItems), fmt.Sprintf("%d", total24hItems), fmt.Sprintf("%d", total7hItems), fmt.Sprintf("%d", total1mItems)))
	}
	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return util.Render(ctx, o.OrderTemplate(user, project.Name, projectID, fmt.Sprintf("%d", totalItems), fmt.Sprintf("%d", total24hItems), fmt.Sprintf("%d", total7hItems), fmt.Sprintf("%d", total1mItems)))
}

// OrderTable returns the order dashboard
func (dh *DashboardHandler) OrderTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	page := ctx.Param("page")
	status, err := w.StringToOrderStatus(ctx.Param("status"))
	pageNum, err := strconv.Atoi(page)
	sort := ctx.QueryParam("sort")
	direction := ctx.QueryParam("direction")
	itemsPerPage := 10
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid page number: %v", err))
	}

	if sort == "" {
		sort = "orderId"
	}
	if direction == "" {
		direction = "asc"
	}

	var wg sync.WaitGroup
	// Fetch	var wg sync.WaitGroup
	wg.Add(2)

	orderCountChan := make(chan int64, 1)
	orderListChan := make(chan []*w.OrderRecord, 1)
	errChan := make(chan error, 1)
	errListChan := make(chan error, 1)

	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, status, "", orderCountChan, errChan)
	}()

	// Fetch  10 orders
	go func() {
		defer wg.Done()
		dh.orderSvc.FindOrderByProjectIDAsync(projectID, itemsPerPage, pageNum, status, sort, direction, orderListChan, errListChan)
	}()
	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(orderCountChan)
		close(orderListChan)
		close(errChan)
		close(errListChan)
	}()

	var totalItems int64
	var orderRecords []*w.OrderRecord

	for err := range errChan {
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"failed to fetch order count": err.Error()})
		}
	}
	for errList := range errListChan {
		if errList != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error to fetch order": errList.Error()})
		}
	}

	for item := range orderCountChan {
		totalItems = item
	}
	for item := range orderListChan {
		orderRecords = item
	}

	// Convert orderRecords to OrderTableList for the response
	var orders []w.OrderTableList
	if orderRecords != nil {
		for _, record := range orderRecords {
			orders = append(orders, w.OrderTableList{
				ID:             record.ID,
				ProjectID:      record.ProjectID,
				OrderCreated:   record.OrderCreated,
				OrderID:        record.OrderID,
				TotalAmount:    record.Order.Total,
				Status:         record.Status,
				Billing:        *record.Order.Billing,
				Shipping:       *record.Order.Shipping,
				Products:       record.Order.LineItems,
				CurrencySymbol: record.Order.CurrencySymbol,
				PaymentMethod:  record.Order.PaymentMethodTitle,
				CustomerNote:   record.Order.CustomerNote,
			})
		}
	}

	// Prepare metadata

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

// OrderTableHTML get html table
func (dh *DashboardHandler) OrderTableHTML(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	return util.Render(ctx, tab.OrderTable(projectID))
}

// OrderBulkAction bulk action for orders
func (dh *DashboardHandler) OrderBulkAction(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	var request w.BulkActionRequest

	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	intOrders, err := util.ConvertStringsToInt64(request.Orders)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	// Get project details using the project service
	project, err := dh.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	switch request.Status {
	case "asc_courier":
		dh.orderSvc.BatchUpdateOrdersStatus(projectID, intOrders, "completed", project)
		//TODO: check is use has this extention and then perform action
	case "completed":
		dh.orderSvc.BatchUpdateOrdersStatus(projectID, intOrders, request.Status, project)
	case "pending":
		dh.orderSvc.BatchUpdateOrdersStatus(projectID, intOrders, request.Status, project)
	case "processing":
		dh.orderSvc.BatchUpdateOrdersStatus(projectID, intOrders, request.Status, project)
	case "cancelled":
		dh.orderSvc.BatchUpdateOrdersStatus(projectID, intOrders, request.Status, project)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Orders status updated successfully"})

}

// OrderCharts return charts for orders
func (dh *DashboardHandler) OrderCharts(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	return util.Render(ctx, chart.OrderCharts(projectID))
}

// OrderMonthlyCharts return charts for orders
func (dh *DashboardHandler) OrderMonthlyCharts(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	now := time.Now()

	months := make([]string, 12)
	orderCounts := make([]int, 12)

	// Get order counts by month from the service
	orderMap, err := dh.analyticsRepo.FindLatestMonthlyCount(projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	fmt.Println(orderMap)

	// Iterate over the last 12 months
	for i := 0; i < 12; i++ {
		// Format as "Jan" for month name and "2006" for the year
		month := now.AddDate(0, -i, 0).Format("Jan") // Format as "Jan 2024", "Dec 2023", etc.
		months[11-i] = month                         // Store in reverse order

		// Generate key for looking up order count in the format "YYYY-MM"
		orderMonthKey := now.AddDate(0, -i, 0).Format("2006-01")

		// Get order count from the map if it exists, otherwise default to 0
		if count, exists := orderMap.MonthyData[orderMonthKey]; exists {
			orderCounts[11-i] = count
		} else {
			orderCounts[11-i] = 0
		}
	}

	// Struct for the JSON response
	type ChartData struct {
		Months []string `json:"months"`
		Orders []int    `json:"orders"`
	}

	// Create and return chart data
	chartData := ChartData{
		Months: months,
		Orders: orderCounts,
	}

	return ctx.JSON(http.StatusOK, chartData)
}

// OrderUpdate return charts for orders
func (dh *DashboardHandler) OrderUpdate(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	orderIDstr := ctx.Param("orderId")
	orderIO, err := strconv.ParseInt(orderIDstr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	req := new(w.OrderTableList)

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	// Get project details using the project service
	project, err := dh.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	_, err = dh.orderSvc.UpdateOrder(projectID, orderIO, req, project)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Order updated successfully"})
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
	userID, err := auth.GetUserID(ctx)
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
