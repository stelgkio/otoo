package cronjob

import (
	"fmt"
	"math"
	"sync"
	"time"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
)

// ProductBestSellerCron  represents the cron job for order analytics
type ProductBestSellerCron struct {
	projectSvc  port.ProjectService
	userSvc     port.UserService
	customerSvc port.CustomerService
	productSvc  port.ProductService
	orderSvc    port.OrderService
}

// NewProductBestSellerCron creates a new OrderAnalyticsCron instance
func NewProductBestSellerCron(projectSvc port.ProjectService, userSvc port.UserService, customerSvc port.CustomerService, productSvc port.ProductService, orderSvc port.OrderService) *ProductBestSellerCron {
	return &ProductBestSellerCron{
		projectSvc:  projectSvc,
		userSvc:     userSvc,
		customerSvc: customerSvc,
		productSvc:  productSvc,
		orderSvc:    orderSvc,
	}
}

// RunAProductBestSellerDailyJob runs the analytics job
func (as *ProductBestSellerCron) RunAProductBestSellerDailyJob() error {
	// Get the current time
	now := time.Now()

	// Define time ranges
	last24Hours := now.Add(-24 * time.Hour)
	lastWeek := now.Add(-7 * 24 * time.Hour)
	lastMonth := now.Add(-30 * 24 * time.Hour)

	allProjects, err := as.projectSvc.GetAllProjects()
	if err != nil {
		return err
	}

	for _, project := range allProjects {
		projectID := project.Id.String()

		totalCount, err := as.orderSvc.GetOrderCount(projectID, w.OrderStatusCompleted)
		if err != nil {
			return err
		}

		last24HoursCount, err := as.orderSvc.GetOrdersCountBetweenOrEquals(projectID, last24Hours, w.OrderStatusCompleted)
		if err != nil {
			return err
		}
		lastWeekCount, err := as.orderSvc.GetOrdersCountBetweenOrEquals(projectID, lastWeek, w.OrderStatusCompleted)
		if err != nil {
			return err
		}
		lastMonthCount, err := as.orderSvc.GetOrdersCountBetweenOrEquals(projectID, lastMonth, w.OrderStatusCompleted)
		if err != nil {
			return err
		}

		// Calculate percentages
		last24HoursPercentage := (float64(last24HoursCount) / float64(totalCount)) * 100
		lastWeekPercentage := (float64(lastWeekCount) / float64(totalCount)) * 100
		lastMonthPercentage := (float64(lastMonthCount) / float64(totalCount)) * 100

		// Print the results
		fmt.Printf("Orders in the last 24 hours: %d (%.2f%%)\n", last24HoursCount, last24HoursPercentage)
		fmt.Printf("Orders in the last week: %d (%.2f%%)\n", lastWeekCount, lastWeekPercentage)
		fmt.Printf("Orders in the last month: %d (%.2f%%)\n", lastMonthCount, lastMonthPercentage)

		return nil
	}

	return nil
}

// RunAProductBestSellerInitializerJob runs the analytics job
func (as *ProductBestSellerCron) RunAProductBestSellerInitializerJob(projectID string) error {

	var wg sync.WaitGroup

	totalCount, err := as.orderSvc.GetOrderCount(projectID, w.OrderStatusCompleted)
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

			as.orderSvc.FindOrderByProjectIDAsync(projectID, 10, i+1, orderListResults, orderListErrors)
		}(i)
	}

	go func() {
		wg.Wait()
		close(orderListResults)
		close(orderListErrors)
	}()

	for items := range orderListResults {
		for _, order := range items {
			as.productSvc.ExtractProductFromOrderAndUpsert(nil, order)
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
		as.productSvc.GetProductBestSeller(projectID, productBestSellers, productBestSellersErrors)
	}()

	bestSellerWg.Wait()
	close(productBestSellers)
	close(productBestSellersErrors)

	for items := range productBestSellers {
		for _, product := range items {
			fmt.Println(product)
		}
	}
	for err := range productBestSellersErrors {
		fmt.Println("Error finding orders:", err)
		return err
	}

	return nil

}
