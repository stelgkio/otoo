package woocommerce

import (
	"fmt"
	"log"
	"log/slog"
	"math"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	d "github.com/stelgkio/otoo/internal/core/domain"
	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	commerce "github.com/stelgkio/woocommerce"
)

// OrderService represents the service for managing orders
type OrderService struct {
	p             port.WoocommerceRepository
	s             port.ProjectRepository
	extensionSrv  port.ExtensionService
	voucherSvc    port.VoucherService
	analyticsRepo port.AnalyticsRepository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(woorepo port.WoocommerceRepository, projrepo port.ProjectRepository, extensionSrv port.ExtensionService, voucherSvc port.VoucherService, analyticsRepo port.AnalyticsRepository) *OrderService {
	return &OrderService{
		p:             woorepo,
		s:             projrepo,
		extensionSrv:  extensionSrv,
		voucherSvc:    voucherSvc,
		analyticsRepo: analyticsRepo,
	}
}

// GetOrderCountAsync retrieves the count of orders for a given project ID
func (os *OrderService) GetOrderCountAsync(ctx echo.Context, projectID string, orderStatus w.OrderStatus, timeRange string, results chan<- int64, errors chan<- error) {
	orderCount, err := os.p.GetOrderCount(projectID, orderStatus, timeRange)
	if err != nil {
		errors <- err
	} else {
		results <- orderCount
	}
}

// GetOrderCountWithDeleteAsync retrieves the count of orders for a given project ID
func (os *OrderService) GetOrderCountWithDeleteAsync(ctx echo.Context, projectID string, orderStatus w.OrderStatus, timeRange string, results chan<- int64, errors chan<- error) {
	orderCount, err := os.p.GetOrderCountWithDelete(projectID, orderStatus, timeRange)
	if err != nil {
		errors <- err
	} else {
		results <- orderCount
	}
}

// GetOrderCount retrieves the count of orders for a given project ID
func (os *OrderService) GetOrderCount(projectID string, orderStatus w.OrderStatus, timeRange string) (int64, error) {
	orderCount, err := os.p.GetOrderCount(projectID, orderStatus, timeRange)
	if err != nil {
		return 0, err
	}
	return orderCount, nil
}

// GetOrderCountWithDelete retrieves the count of orders for a given project ID
func (os *OrderService) GetOrderCountWithDelete(projectID string, orderStatus w.OrderStatus, timeRange string) (int64, error) {
	orderCount, err := os.p.GetOrderCountWithDelete(projectID, orderStatus, timeRange)
	if err != nil {
		return 0, err
	}
	return orderCount, nil
}

// Get10LatestOrders retrieves the latest 10 orders for a given project ID
func (os *OrderService) Get10LatestOrders(ctx echo.Context, projectID string, orderStatus w.OrderStatus, sort string, results chan<- []*domain.OrderRecord, errors chan<- error) {
	orders, err := os.p.OrderFindByProjectID(projectID, 10, 1, orderStatus, sort, "")
	if err != nil {
		errors <- err
	} else {
		results <- orders
	}
}

// GetOrdersCountBetweenOrEquals retrieves the count of orders between or equal to a given date for a given project ID
func (os *OrderService) GetOrdersCountBetweenOrEquals(projectID string, timePeriod time.Time, orderStatus w.OrderStatus) (int64, error) {
	orderCount, err := os.p.GetOrdersCountBetweenOrEquals(projectID, timePeriod, orderStatus)
	if err != nil {
		return 0, err
	}
	return orderCount, nil

}

// FindOrderByProjectIDAsync retrieves orders for a given project ID
func (os *OrderService) FindOrderByProjectIDAsync(projectID string, size, page int, orderStatus w.OrderStatus, sort, direction string, results chan<- []*domain.OrderRecord, errors chan<- error) {
	orderCount, err := os.p.OrderFindByProjectID(projectID, size, page, orderStatus, sort, direction)
	if err != nil {
		errors <- err
	} else {
		results <- orderCount
	}
}

// FindOrderByProjectIDWithTimePeriodAsync retrieves orders for a given project ID
func (os *OrderService) FindOrderByProjectIDWithTimePeriodAsync(projectID string, size, page int, orderStatus w.OrderStatus, sort, direction string, timePeriod time.Time, results chan<- []*domain.OrderRecord, errors chan<- error) {
	orderCount, err := os.p.OrderFindByProjectIDWithTimePedio(projectID, size, page, orderStatus, sort, direction, timePeriod)
	if err != nil {
		errors <- err
	} else {
		results <- orderCount
	}
}

// GetOrderByID retrieves an order by its ID
func (os *OrderService) GetOrderByID(projectID string, orderID int64) (*domain.OrderRecord, error) {
	order, err := os.p.GetOrderByID(projectID, orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// UpdateOrderStatusByID updates order by ID
func (os *OrderService) UpdateOrderStatusByID(projectID string, orderID int64, status string, project *d.Project) (*domain.OrderRecord, error) {
	panic("not implemented") // TODO: Implement
}

// BatchUpdateOrdersStatus updates batch orders
func (os *OrderService) BatchUpdateOrdersStatus(projectID string, orders []int64, status string, proj *d.Project) ([]*domain.OrderRecord, error) {
	client := InitClient(proj.WoocommerceProject.ConsumerKey, proj.WoocommerceProject.ConsumerSecret, proj.WoocommerceProject.Domain)

	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, len(orders))
	ordersToUpdate := make([]*domain.OrderRecord, len(orders))

	updateOrder := func(i int, orderID int64) {
		defer wg.Done()
		order, err := os.p.GetOrderByID(projectID, orderID)
		if err != nil {
			errChan <- err
			return
		}
		if order == nil {
			return
		}
		s, err := domain.StringToOrderStatus(status)
		if err != nil {
			errChan <- err
			return
		}

		order.Status = s
		order.Order.Status = status
		order.Timestamp = time.Now().UTC()
		order.UpdatedAt = time.Now().UTC()
		err = os.p.OrderUpdate(order, orderID)
		if err != nil {
			errChan <- err
			return
		}

		mu.Lock()
		ordersToUpdate[i] = order
		mu.Unlock()

		_, err = client.Order.Update(&order.Order)
		if err != nil {
			errChan <- err
			return
		}
	}

	wg.Add(len(orders))
	for i, orderID := range orders {
		go updateOrder(i, orderID)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return ordersToUpdate, nil
}

// UpdateOrder updates an order
func (os *OrderService) UpdateOrder(projectID string, orderID int64, orderTable *domain.OrderTableList, proj *d.Project) (*domain.OrderRecord, error) {
	client := InitClient(proj.WoocommerceProject.ConsumerKey, proj.WoocommerceProject.ConsumerSecret, proj.WoocommerceProject.Domain)

	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, 1)
	ordersToUpdate := make([]*domain.OrderRecord, 1)

	updateOrder := func(orderID int64) {
		defer wg.Done()
		order, err := os.p.GetOrderByID(projectID, orderID)
		if err != nil {
			errChan <- err
			return
		}
		if order == nil {
			return
		}

		order.Order.Shipping = &orderTable.Shipping
		order.Order.CustomerNote = orderTable.CustomerNote
		order.Order.Billing = &orderTable.Billing
		order.Timestamp = time.Now().UTC()
		order.UpdatedAt = time.Now().UTC()
		err = os.p.OrderUpdate(order, orderID)
		if err != nil {
			errChan <- err
			return
		}

		_, err = os.voucherSvc.UpdateVoucher(nil, order, projectID)
		if err != nil {
			errChan <- err
			return
		}

		mu.Lock()
		ordersToUpdate[0] = order
		mu.Unlock()

		_, err = client.Order.Update(&order.Order)
		if err != nil {
			errChan <- err
			return
		}
	}

	wg.Add(1)
	go updateOrder(orderID)

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return ordersToUpdate[0], nil
}

// GetAllOrdersFromWoocommerce retrieves all orders from WooCommerce and saves them to MongoDB
func (os *OrderService) GetAllOrdersFromWoocommerce(customerKey string, customerSecret string, domainURL string, projectID string, totalOrder int64) error {
	client := InitClient(customerKey, customerSecret, domainURL)

	// Create all webhooks
	err := os.createAndSaveAllOrders(client, projectID, totalOrder)
	if err != nil {
		slog.Error("get all order error", "error", err)
		return errors.Wrap(err, "create all order error")
	}

	slog.Info("get all orders success")
	return nil
}

// createAndSaveAllWebhooks creates WooCommerce products and saves results to MongoDB concurrently
func (os *OrderService) createAndSaveAllOrders(client *commerce.Client, projectID string, totalOrder int64) error {

	var wg sync.WaitGroup
	orderCh := make(chan *w.OrderRecord, totalOrder) // Channel to distribute orders to workers
	errorCh := make(chan *w.OrderRecord, 1)          // Buffered channel for error results

	// Number of worker goroutines for processing orders
	workers := int(math.Ceil(float64(totalOrder) / 100))
	if workers == 0 {
		workers = 1
	}

	// Worker pool to process orders
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(projectID string) {
			defer wg.Done()
			for order := range orderCh {
				err := os.saveOrderResult(order, order.OrderID)
				if err != nil {
					log.Printf("Failed to save order result: %v", err)
					errorCh <- order // Send the error order to the error channel
				}
			}
		}(projectID)
	}

	// Goroutine to save error results
	go func(projectID string) {
		for result := range errorCh {
			err := os.saveOrderResult(result, result.OrderID)
			if err != nil {
				log.Printf("Failed to save order error result: %v", err)
			}
		}
	}(projectID)

	// Create a semaphore to limit the number of concurrent fetches
	maxConcurrentFetches := 5
	sem := make(chan struct{}, maxConcurrentFetches) // Semaphore to limit concurrent requests
	var totalFetched int = 0
	var mu sync.Mutex
	// Concurrent fetching of orders
	go func() {
		var fetchWg sync.WaitGroup
		page := 1

		for {
			sem <- struct{}{} // Acquire a token before launching a new fetch goroutine
			fetchWg.Add(1)

			go func(page int) {
				defer fetchWg.Done()
				defer func() { <-sem }() // Release the semaphore token when done

				// Fetch orders from the API
				resp, err := os.fetchOrdesWithRetry(client, page)
				if err != nil {
					errorCh <- &w.OrderRecord{
						ProjectID: projectID,
						Event:     "order.List",
						Error:     "Failed to fetch orders",
						CreatedAt: time.Now().UTC(),
						OrderID:   0,
						Order:     commerce.Order{}, // Assuming empty order
						IsActive:  false,
					}
					return
				}

				if len(resp) == 0 {
					// No more orders, stop the fetching
					return
				}
				mu.Lock()
				totalFetched += len(resp)
				mu.Unlock()
				// Notify about the number of orders received
				os.extensionSrv.UpdateSynchronizerOrderReceivedExtension(nil, projectID, totalFetched)

				// Send fetched orders to the worker pool
				for _, item := range resp {
					Status, _ := domain.StringToOrderStatus(item.Status)
					_, orderCreated, _ := domain.ConvertDateString(item.DateCreatedGmt)

					orderCh <- &w.OrderRecord{
						ProjectID:    projectID,
						Event:        "order.created",
						Error:        "",
						CreatedAt:    time.Now().UTC(),
						Timestamp:    time.Now().UTC(),
						OrderID:      item.ID,
						Order:        item,
						IsActive:     true,
						Status:       Status,
						OrderCreated: orderCreated,
					}
				}
			}(page)

			page++

			// Define the exit condition to stop fetching orders
			if page > int(math.Ceil(float64(totalOrder)/100)) {
				break
			}
		}

		// Wait for all fetch goroutines to complete
		fetchWg.Wait()

		// Close the order channel after all fetching is done
		close(orderCh)
	}()

	// Wait for all worker goroutines to finish processing orders
	wg.Wait()

	// Close the error channel after all error processing is done
	close(errorCh)

	return nil
}

// saveWebhookResult saves webhook creation result to MongoDB
func (os *OrderService) saveOrderResult(data *w.OrderRecord, orderID int64) error {
	if data.Status == "processing" {
		go os.voucherSvc.CreateVoucher(nil, data, data.ProjectID)
	}
	err := os.p.OrderUpdate(data, orderID)
	if err != nil {
		return errors.Wrap(err, "failed to insert order result into MongoDB")
	}
	return nil
}

// GetLatestOrderWeeklyBalance retrieves the latest order weekly balance for a given project ID
func (os *OrderService) GetLatestOrderWeeklyBalance(ctx echo.Context, projectID string, results chan<- *domain.WeeklyAnalytics, errors chan<- error) {
	orderCount, err := os.analyticsRepo.FindLatestWeeklyBalance(projectID)
	if err != nil {
		errors <- err
	} else {
		results <- orderCount
	}

}

// CountOrdersByMonth retrieves the latest order monthly balance for a given project ID
func (os *OrderService) CountOrdersByMonth(projectID string) (map[string]int, error) {
	return os.p.CountOrdersByMonth(projectID)
}

// fetchOrdesWithRetry fetches a page of products with retry logic.
func (os *OrderService) fetchOrdesWithRetry(client *commerce.Client, page int) ([]commerce.Order, error) {
	options := commerce.OrderListOption{
		ListOptions: commerce.ListOptions{
			Context: "view",
			Order:   "desc",
			Orderby: "id",
			Page:    page,
			PerPage: batchSize,
		},
	}

	var orders []commerce.Order
	var err error

	for i := 0; i < maxRetries; i++ {
		orders, err = client.Order.List(options)
		if err == nil {
			return orders, nil // Success
		}

		log.Printf("Failed to fetch orders (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second) // Exponential backoff
	}

	return nil, fmt.Errorf("failed after %d retries: %v", maxRetries, err)
}
