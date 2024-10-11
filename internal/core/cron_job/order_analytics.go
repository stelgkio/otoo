package cronjob

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
)

// OrderAnalyticsCron represents the cron job for order analytics
type OrderAnalyticsCron struct {
	projectSvc    port.ProjectService
	userSvc       port.UserService
	customerSvc   port.CustomerService
	productSvc    port.ProductService
	orderSvc      port.OrderService
	analyticsRepo port.AnalyticsRepository
}

// NewOrderAnalyticsCron creates a new OrderAnalyticsCron instance
func NewOrderAnalyticsCron(projectSvc port.ProjectService, userSvc port.UserService, customerSvc port.CustomerService, productSvc port.ProductService, orderSvc port.OrderService, analyticsRepo port.AnalyticsRepository) *OrderAnalyticsCron {
	return &OrderAnalyticsCron{
		projectSvc:    projectSvc,
		userSvc:       userSvc,
		customerSvc:   customerSvc,
		productSvc:    productSvc,
		orderSvc:      orderSvc,
		analyticsRepo: analyticsRepo,
	}
}

// RunOrderWeeklyBalanceJob runs the analytics job
func (as *OrderAnalyticsCron) RunOrderWeeklyBalanceJob() error {
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
			if err := as.RunOrderWeeklyBalanceInitializeJob(projID); err != nil {
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

// RunOrderWeeklyBalanceInitializeJob runs the analytics job
func (as *OrderAnalyticsCron) RunOrderWeeklyBalanceInitializeJob(projectID string) error {
	var wg sync.WaitGroup

	now := time.Now().UTC()
	lastWeek := now.Add(-7 * 24 * time.Hour)

	//project, err := as.projectSvc.GetProjectByID(nil, projectID)
	lastWeekCount, err := as.orderSvc.GetOrdersCountBetweenOrEquals(projectID, lastWeek, w.OrderStatusCompleted)
	orderCount, err := as.orderSvc.GetOrderCount(projectID, w.OrderStatusCompleted, "")
	if err != nil {
		return err
	}

	worker := int(math.Ceil(float64(lastWeekCount) / 10))

	wg.Add(worker)
	orderListResults := make(chan []*w.OrderRecord, worker)
	orderListErrors := make(chan error, 1)
	for i := 0; i < worker; i++ {
		go func(i int) {
			defer wg.Done()

			as.orderSvc.FindOrderByProjectIDWithTimePeriodAsync(projectID, 10, i+1, w.OrderStatusCompleted, "orderId", "asc", lastWeek, orderListResults, orderListErrors)
		}(i)
	}

	go func() {
		wg.Wait()
		close(orderListResults)
		close(orderListErrors)
	}()

	var ordertWeeklyBalance float64 = 0.0
	for items := range orderListResults {
		for _, order := range items {
			total, err := strconv.ParseFloat(order.Order.Total, 64)
			if err != nil {
				fmt.Println("Error parsing order total:", err)
				continue
			}
			ordertWeeklyBalance += total
		}
	}

	for err := range orderListErrors {
		fmt.Println("Error finding orders:", err)
		return err
	}

	orderWeeklyBalance := make(chan *w.WeeklyAnalytics, 1)
	orderWeeklyBalanceErrors := make(chan error, 1)

	var bestSellerWg sync.WaitGroup
	bestSellerWg.Add(1)
	go func() {
		defer bestSellerWg.Done()
		as.orderSvc.GetLatestOrderWeeklyBalance(nil, projectID, orderWeeklyBalance, orderWeeklyBalanceErrors)
	}()

	bestSellerWg.Wait()
	close(orderWeeklyBalance)
	close(orderWeeklyBalanceErrors)

	//TODO: change this to save the weeklybalance
	//as.bestSellerSvc.DeleteBestSellers(projectID)

	for items := range orderWeeklyBalance {

		weekBalance := w.NewWeeklyAnalytics(projectID, orderCount, lastWeekCount, ordertWeeklyBalance, lastWeek, now)
		weekBalance.CalculatePercentages()
		if items != nil {
			result := w.CompareAnalytics(weekBalance.AnalyticsBase, items.AnalyticsBase)
			weekBalance.AddComparisonResult(result)
			fmt.Println("The result of the 2 weeks is", result)
		}
		as.analyticsRepo.CreateWeeklyBalance(projectID, weekBalance)
	}
	for err := range orderWeeklyBalanceErrors {
		fmt.Println("Error finding orders:", err)
		return err
	}

	return nil
}

// RunOrderMonthlyCountJob runs the analytics job
func (as *OrderAnalyticsCron) RunOrderMonthlyCountJob() error {
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
			if err := as.RunOrderMonthlyCountInitializeJob(projID); err != nil {
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

// RunOrderMonthlyCountInitializeJob runs the analytics job
func (as *OrderAnalyticsCron) RunOrderMonthlyCountInitializeJob(projectID string) error {
	now := time.Now().UTC()
	orderMap, err := as.orderSvc.CountOrdersByMonth(projectID)
	nownew := time.Now().UTC()
	newData := w.NewMonthlyAnalytics(projectID, orderMap, now, nownew)

	err = as.analyticsRepo.CreateMonthlyCount(projectID, &newData)
	if err != nil {
		return err
	}
	return nil
}
