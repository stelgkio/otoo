package cronjob

import (
	"fmt"
	"math"
	"sync"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
)

// ProductBestSellerCron  represents the cron job for order analytics
type ProductBestSellerCron struct {
	projectSvc    port.ProjectService
	AnalyticsRepo port.AnalyticsRepository
	customerSvc   port.CustomerService
	productSvc    port.ProductService
	orderSvc      port.OrderService
}

// NewProductBestSellerCron creates a new OrderAnalyticsCron instance
func NewProductBestSellerCron(projectSvc port.ProjectService, AnalyticsRepo port.AnalyticsRepository, customerSvc port.CustomerService, productSvc port.ProductService, orderSvc port.OrderService) *ProductBestSellerCron {
	return &ProductBestSellerCron{
		projectSvc:    projectSvc,
		AnalyticsRepo: AnalyticsRepo,
		customerSvc:   customerSvc,
		productSvc:    productSvc,
		orderSvc:      orderSvc,
	}
}

// RunAProductBestSellerDailyJob runs the analytics job
func (as *ProductBestSellerCron) RunAProductBestSellerDailyJob() error {
	// Get all projects
	allProjects, err := as.projectSvc.GetAllProjects()
	if err != nil {
		return err
	}

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	// Channel to collect errors from the goroutines
	errChan := make(chan error, len(allProjects))

	// Iterate over projects and run the job concurrently
	for _, project := range allProjects {
		wg.Add(1) // Increment the WaitGroup counter

		// Capture projectID to avoid issues with goroutine closures
		projectID := project.Id.String()

		go func(projID string) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			// Run the initializer job for each project
			if err := as.RunAProductBestSellerInitializerJob(projID); err != nil {
				// Send error to the channel if any occurs
				errChan <- err
			}
		}(projectID)
	}

	// Close the error channel once all goroutines are done
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Check for errors from the error channel
	for e := range errChan {
		if e != nil {
			return e // Return the first error encountered
		}
	}

	return nil // Return nil if no errors
}

// RunAProductBestSellerInitializerJob runs the analytics job
func (as *ProductBestSellerCron) RunAProductBestSellerInitializerJob(projectID string) error {

	var wg sync.WaitGroup

	project, err := as.projectSvc.GetProjectByID(nil, projectID)
	totalCount, err := as.orderSvc.GetOrderCount(projectID, w.OrderStatusCompleted, "")
	if err != nil {
		return err
	}

	worker := int(math.Ceil(float64(totalCount) / 10))

	wg.Add(worker)
	orderListResults := make(chan []*w.OrderRecord, worker)
	orderListErrors := make(chan error, 1)
	for i := 0; i < worker; i++ {
		go func(i int) {
			defer wg.Done()

			as.orderSvc.FindOrderByProjectIDAsync(projectID, 10, i+1, w.OrderStatusCompleted, "orderId", "asc", orderListResults, orderListErrors)
		}(i)
	}

	go func() {
		wg.Wait()
		close(orderListResults)
		close(orderListErrors)
	}()

	for items := range orderListResults {
		for _, order := range items {
			as.productSvc.ExtractProductFromOrderAndUpsert(nil, order, project)
			as.customerSvc.ExtractCustomerFromOrderAndUpsert(nil, order)
		}
	}

	for err := range orderListErrors {
		fmt.Println("Error finding orders:", err)
		return err
	}

	productBestSellers := make(chan []*w.ProductBestSellerRecord, 1)
	productBestSellersErrors := make(chan error, 1)

	var bestSellerWg sync.WaitGroup
	bestSellerWg.Add(1)
	go func() {
		defer bestSellerWg.Done()
		as.productSvc.GetProductBestSeller(projectID, totalCount, productBestSellers, productBestSellersErrors)
	}()

	bestSellerWg.Wait()
	close(productBestSellers)
	close(productBestSellersErrors)

	as.AnalyticsRepo.DeleteBestSellers(projectID)
	//TODO: find best seller rates
	for items := range productBestSellers {
		as.AnalyticsRepo.CreateBestSellers(projectID, items)
	}
	for err := range productBestSellersErrors {
		fmt.Println("Error finding orders:", err)
		return err
	}

	return nil

}
