package cronjob

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/stelgkio/otoo/internal/core/domain"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
)

// OrderAnalyticsCron represents the cron job for order analytics
type OrderAnalyticsCron struct {
	projectSvc    port.ProjectService
	userSvc       port.UserService
	customerSvc   port.CustomerService
	productSvc    port.ProductService
	orderSvc      port.OrderService
	analyticsRepo port.AnalyticsRepository
	smtpSvc       port.SmtpService
}

// NewOrderAnalyticsCron creates a new OrderAnalyticsCron instance
func NewOrderAnalyticsCron(projectSvc port.ProjectService,
	userSvc port.UserService,
	customerSvc port.CustomerService,
	productSvc port.ProductService,
	orderSvc port.OrderService,
	analyticsRepo port.AnalyticsRepository,
	smtpSvc port.SmtpService) *OrderAnalyticsCron {
	return &OrderAnalyticsCron{
		projectSvc:    projectSvc,
		userSvc:       userSvc,
		customerSvc:   customerSvc,
		productSvc:    productSvc,
		orderSvc:      orderSvc,
		analyticsRepo: analyticsRepo,
		smtpSvc:       smtpSvc,
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
		//	projectID := project.Id.String()
		//TODO: user should be admin
		usersSlice, err := as.userSvc.FindUsersByProjectId(nil, project.Id)

		if err != nil {
			return err
		}
		go func(project *domain.Project, users []*domain.User) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			// Run the initializer job for each project
			err := as.RunOrderWeeklyBalanceInitializeJob(project, users)
			if err != nil {
				// Send error to the channel if any occurs
				errChan <- err
			}
		}(project, usersSlice)
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
func (as *OrderAnalyticsCron) RunOrderWeeklyBalanceInitializeJob(project *domain.Project, users []*domain.User) error {
	var wg sync.WaitGroup
	// Log input parameters
	if project == nil {
		return fmt.Errorf("project is nil")
	}
	if len(users) == 0 {
		return fmt.Errorf("user is nil")
	}
	usersSlice := util.Filter(users, func(u *domain.User) bool {
		return u.ReseveNotification == true
	})
	now := time.Now().UTC()
	lastWeek := now.Add(-7 * 24 * time.Hour)

	//project, err := as.projectSvc.GetProjectByID(nil, projectID)
	lastWeekCount, err := as.orderSvc.GetOrdersCountBetweenOrEquals(project.Id.String(), lastWeek, w.OrderStatusCompleted)
	orderCount, err := as.orderSvc.GetOrderCount(project.Id.String(), w.OrderStatusCompleted, "")
	if err != nil {
		return err
	}

	worker := int(math.Ceil(float64(lastWeekCount) / 10))
	if worker == 0 {
		worker = 1
	}
	wg.Add(worker)
	orderListResults := make(chan []*w.OrderRecord, worker)
	orderListErrors := make(chan error, 1)
	for i := 0; i < worker; i++ {
		go func(i int) {
			defer wg.Done()

			as.orderSvc.FindOrderByProjectIDWithTimePeriodAsync(project.Id.String(), 10, i+1, w.OrderStatusCompleted, "orderId", "asc", lastWeek, orderListResults, orderListErrors)
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
		as.orderSvc.GetLatestOrderWeeklyBalance(nil, project.Id.String(), orderWeeklyBalance, orderWeeklyBalanceErrors)
	}()

	bestSellerWg.Wait()
	close(orderWeeklyBalance)
	close(orderWeeklyBalanceErrors)

	for items := range orderWeeklyBalance {

		weekBalance := w.NewWeeklyAnalytics(project.Id.String(), orderCount, lastWeekCount, ordertWeeklyBalance, lastWeek, now)
		weekBalance.CalculatePercentages()
		if items != nil {
			result := w.CompareAnalytics(weekBalance.AnalyticsBase, items.AnalyticsBase)
			weekBalance.AddComparisonResult(result)
			fmt.Println("The result of the 2 weeks is", result)
		}
		as.analyticsRepo.CreateWeeklyBalance(project.Id.String(), weekBalance)
		for _, user := range usersSlice {
			fullname := fmt.Sprintf("%s %s", user.Name, user.LastName)
			as.smtpSvc.SendWeeklyBalanceEmail(nil, weekBalance, user.Email, fullname)
		}
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
		usersList, err := as.userSvc.FindUsersByProjectId(nil, project.Id)
		users := util.Filter(usersList, func(u *domain.User) bool {
			return u.ReseveNotification == true
		})
		if err != nil {
			return err
		}
		go func(projID string) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			// Run the initializer job for each project
			if err := as.RunOrderMonthlyCountInitializeJob(project, users); err != nil {
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
func (as *OrderAnalyticsCron) RunOrderMonthlyCountInitializeJob(project *domain.Project, user []*domain.User) error {
	now := time.Now().UTC()
	nextMonth := now.Add(30 * 24 * time.Hour)
	orderMap, err := as.orderSvc.CountOrdersByMonth(project.Id.String())

	newData := w.NewMonthlyAnalytics(project.Id.String(), orderMap, now, nextMonth)

	err = as.analyticsRepo.CreateMonthlyCount(project.Id.String(), &newData)
	if err != nil {
		return err
	}
	//go as.smtpSvc.SendMonthlyOrdersEmail(nil, &newData, user.Email, fmt.Sprintf("%s %s", user.Name, user.LastName))
	return nil
}
