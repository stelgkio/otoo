package cronjob

import (
	"fmt"
	"time"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
)

// OrderAnalyticsCron represents the cron job for order analytics
type OrderAnalyticsCron struct {
	projectSvc  port.ProjectService
	userSvc     port.UserService
	customerSvc port.CustomerService
	productSvc  port.ProductService
	orderSvc    port.OrderService
}

// NewOrderAnalyticsCron creates a new OrderAnalyticsCron instance
func NewOrderAnalyticsCron(projectSvc port.ProjectService, userSvc port.UserService, customerSvc port.CustomerService, productSvc port.ProductService, orderSvc port.OrderService) *OrderAnalyticsCron {
	return &OrderAnalyticsCron{
		projectSvc:  projectSvc,
		userSvc:     userSvc,
		customerSvc: customerSvc,
		productSvc:  productSvc,
		orderSvc:    orderSvc,
	}
}

// RunAnalyticsJob runs the analytics job
func (as *OrderAnalyticsCron) RunAnalyticsJob() error {
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

		totalCount, err := as.orderSvc.GetOrderCount(projectID, w.OrderStatusCompleted, "")
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
