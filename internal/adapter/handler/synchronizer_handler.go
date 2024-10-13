package handler

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	synpage "github.com/stelgkio/otoo/internal/adapter/web/view/component/data_synchronizer/synchronize"
	syntmp "github.com/stelgkio/otoo/internal/adapter/web/view/component/data_synchronizer/synchronize/template"
	syn "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/progress/synchronize"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/util"
	r "github.com/stelgkio/otoo/internal/core/util"
)

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
	syncro, err := ph.extensionSvc.GetSynchronizerProjectExtensionByID(ctx, "", projectID)
	if syncro != nil {

		customerPercentage := (float64(syncro.CustomerRecieved) / float64(syncro.TotalCustomer)) * 100.0
		productPercentage := (float64(syncro.ProductReceived) / float64(syncro.TotalProduct)) * 100.0
		orderPercentage := (float64(syncro.OrderReceived) / float64(syncro.TotalOrder)) * 100.0

		return r.Render(ctx, syn.ProjectSynchronizerStart(user, projectID, int64(syncro.TotalCustomer), int64(syncro.TotalProduct), int64(syncro.TotalOrder), customerPercentage, productPercentage, orderPercentage))
	}
	var wg sync.WaitGroup
	wg.Add(3)
	// Create channels for customer, product, and order totals and errors
	customerChan := make(chan int, 1)
	productChan := make(chan int, 1)
	orderChan := make(chan int, 1)
	errorChan := make(chan error, 1)

	// Run service calls asynchronously using goroutines
	go func() {
		defer wg.Done()
		customerTotal, err := ph.reportSvc.GetCustomerTotalCount(ctx, projectID)
		if err != nil {
			errorChan <- err
			return
		}

		customerChan <- customerTotal
	}()

	go func() {
		defer wg.Done()
		productTotal, err := ph.reportSvc.GetProductTotalCount(ctx, projectID)
		if err != nil {
			errorChan <- err
			return
		}
		productChan <- productTotal
	}()

	go func() {
		defer wg.Done()
		orderTotal, err := ph.reportSvc.GetOrderTotalCount(ctx, projectID)
		if err != nil {
			errorChan <- err
			return
		}
		orderChan <- orderTotal
	}()

	go func() {
		wg.Wait()
		close(orderChan)
		close(productChan)
		close(customerChan)
		close(errorChan)

	}()

	// Initialize variables to store the totals
	var customerTotal, productTotal, orderTotal int

	for count := range orderChan {
		orderTotal = count
	}
	for count := range productChan {
		productTotal = count
	}
	for count := range customerChan {
		customerTotal = count
	}

	for err := range errorChan {
		if err != nil {
			return fmt.Errorf(" error: %v", err)
		}
	}

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

	nto := domain.CreateAnalyticsNotification(userID.String(), projectID)
	ph.notificationSvc.CreateNotification(ctx, nto)
	exten, err := ph.extensionSvc.GetExtensionByCode(ctx, domain.DataSynchronizerCode)
	if err != nil {
		return err
	}
	synchronizer := domain.NewDataSynchronizerExtension(projectID, exten.ID.Hex(), int(customerTotal), int(orderTotal), int(productTotal))
	err = ph.extensionSvc.CreateSynchronizerProjectExtension(ctx, projectID, synchronizer)
	if err != nil {
		if errors.Is(err, util.ErrSynchronizerInProgress) {
			syncro, _ := ph.extensionSvc.GetSynchronizerProjectExtensionByID(ctx, "", projectID)
			customerPercentage := (float64(syncro.CustomerRecieved) / float64(syncro.TotalCustomer)) * 100.0
			productPercentage := (float64(syncro.ProductReceived) / float64(syncro.TotalProduct)) * 100.0
			orderPercentage := (float64(syncro.OrderReceived) / float64(syncro.TotalOrder)) * 100.0

			return r.Render(ctx, syn.ProjectSynchronizerStart(user, projectID, int64(syncro.TotalCustomer), int64(syncro.TotalProduct), int64(syncro.TotalOrder), customerPercentage, productPercentage, orderPercentage))
		}
		return err
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
	project, err := ph.svc.GetProjectByID(ctx, projectID)
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

	var orderCount, productCount, customerCount int64 = 0, 0, 0

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
		go ph.bestSellerSvc.RunAProductBestSellerInitializerJob(projectID)
		go ph.orderAnalyticsSvc.RunOrderWeeklyBalanceInitializeJob(project, user)
		go ph.orderAnalyticsSvc.RunOrderMonthlyCountInitializeJob(project, user)
		return r.Render(ctx, syn.ProjectSynchronizerDone(user, projectID, customerTotal, productTotal, orderTotal, 100.0, 100.0, 100.0))
	}
	customerPercentage := (float64(customerCount) / float64(customerTotal)) * 100.0
	productPercentage := (float64(productCount) / float64(productTotal)) * 100.0
	orderPercentage := (float64(orderCount) / float64(orderTotal)) * 100.0

	return r.Render(ctx, syn.ProjectSynchronizerStart(user, projectID, customerTotal, productTotal, orderTotal, customerPercentage, productPercentage, orderPercentage))
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

// ProjectSynchronizePage GET /project/synchronize/:projectId
func (ph *ProjectHandler) ProjectSynchronizePage(ctx echo.Context) error {
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
	syncro, err := ph.extensionSvc.GetSynchronizerProjectExtensionByID(ctx, "", projectID)

	if syncro != nil {

		customerPercentage := (float64(syncro.CustomerRecieved) / float64(syncro.TotalCustomer)) * 100.0
		productPercentage := (float64(syncro.ProductReceived) / float64(syncro.TotalProduct)) * 100.0
		orderPercentage := (float64(syncro.OrderReceived) / float64(syncro.TotalOrder)) * 100.0

		return r.Render(ctx, synpage.ProjectSynchronizerStartPage(user, projectID, int64(syncro.CustomerRecieved), int64(syncro.ProductReceived), int64(syncro.OrderReceived), customerPercentage, productPercentage, orderPercentage))
	}

	var wg sync.WaitGroup
	wg.Add(3)
	// Create channels for customer, product, and order totals and errors
	customerChan := make(chan int, 1)
	productChan := make(chan int, 1)
	orderChan := make(chan int, 1)
	errorChan := make(chan error, 1)

	// Run service calls asynchronously using goroutines
	go func() {
		defer wg.Done()
		customerTotal, err := ph.reportSvc.GetCustomerTotalCount(ctx, projectID)
		if err != nil {
			errorChan <- err
			return
		}

		customerChan <- customerTotal
	}()

	go func() {
		defer wg.Done()
		productTotal, err := ph.reportSvc.GetProductTotalCount(ctx, projectID)
		if err != nil {
			errorChan <- err
			return
		}
		productChan <- productTotal
	}()

	go func() {
		defer wg.Done()
		orderTotal, err := ph.reportSvc.GetOrderTotalCount(ctx, projectID)
		if err != nil {
			errorChan <- err
			return
		}
		orderChan <- orderTotal
	}()

	go func() {
		wg.Wait()
		close(orderChan)
		close(productChan)
		close(customerChan)
		close(errorChan)

	}()

	// Initialize variables to store the totals
	var customerTotal, productTotal, orderTotal int

	for count := range orderChan {
		orderTotal = count
	}
	for count := range productChan {
		productTotal = count
	}
	for count := range customerChan {
		customerTotal = count
	}

	for err := range errorChan {
		if err != nil {
			return fmt.Errorf(" error: %v", err)
		}
	}
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, synpage.ProjectSynchronizerPage(user, projectID, customerTotal, productTotal, orderTotal))
	}

	return r.Render(ctx, syntmp.ProjectSynchonizerTemplate(user, project.Name, projectID, customerTotal, productTotal, orderTotal))
}

// ProjectSynchronizeStartPage POST /project/synchronize/:projectId
func (ph *ProjectHandler) ProjectSynchronizeStartPage(ctx echo.Context) error {
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
	exten, err := ph.extensionSvc.GetExtensionByCode(ctx, domain.DataSynchronizerCode)
	if err != nil {
		return err
	}
	synchronizer := domain.NewDataSynchronizerExtension(projectID, exten.ID.Hex(), int(customerTotal), int(orderTotal), int(productTotal))
	err = ph.extensionSvc.CreateSynchronizerProjectExtension(ctx, projectID, synchronizer)
	if err != nil {
		if errors.Is(err, util.ErrSynchronizerInProgress) {
			syncro, _ := ph.extensionSvc.GetSynchronizerProjectExtensionByID(ctx, "", projectID)
			customerPercentage := (float64(syncro.CustomerRecieved) / float64(syncro.TotalCustomer)) * 100.0
			productPercentage := (float64(syncro.ProductReceived) / float64(syncro.TotalProduct)) * 100.0
			orderPercentage := (float64(syncro.OrderReceived) / float64(syncro.TotalOrder)) * 100.0

			return r.Render(ctx, synpage.ProjectSynchronizerStartPage(user, projectID, int64(syncro.CustomerRecieved), int64(syncro.ProductReceived), int64(syncro.OrderReceived), customerPercentage, productPercentage, orderPercentage))
		}
		return err
	}

	return r.Render(ctx, synpage.ProjectSynchronizerStartPage(user, projectID, customerTotal, productTotal, orderTotal, 0.0, 0.0, 0.0))
}

// ProjectSynchronizeDonePage GET /project/synchronize/done/:projectId
func (ph *ProjectHandler) ProjectSynchronizeDonePage(ctx echo.Context) error {
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
	project, err := ph.svc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if customerTotal == 0 && productTotal == 0 && orderTotal == 0 {
		ctx.Response().Header().Set("HX-Trigger", "done")
		return r.Render(ctx, synpage.ProjectSynchronizerDonePage(user, projectID, customerTotal, productTotal, orderTotal, 100.0, 100.0, 100.0))
	}

	syncro, err := ph.extensionSvc.GetSynchronizerProjectExtensionByID(ctx, "", projectID)

	if syncro != nil {

		customerPercentage := (float64(syncro.CustomerRecieved) / float64(syncro.TotalCustomer)) * 100.0
		productPercentage := (float64(syncro.ProductReceived) / float64(syncro.TotalProduct)) * 100.0
		orderPercentage := (float64(syncro.OrderReceived) / float64(syncro.TotalOrder)) * 100.0

		return r.Render(ctx, synpage.ProjectSynchronizerStartPage(user, projectID, int64(syncro.CustomerRecieved), int64(syncro.ProductReceived), int64(syncro.OrderReceived), customerPercentage, productPercentage, orderPercentage))
	}
	ctx.Response().Header().Set("HX-Trigger", "done")
	go ph.bestSellerSvc.RunAProductBestSellerInitializerJob(projectID)
	go ph.orderAnalyticsSvc.RunOrderWeeklyBalanceInitializeJob(project, user)
	go ph.orderAnalyticsSvc.RunOrderMonthlyCountInitializeJob(project, user)
	return r.Render(ctx, synpage.ProjectSynchronizerDonePage(user, projectID, customerTotal, productTotal, orderTotal, 100.0, 100.0, 100.0))

}
