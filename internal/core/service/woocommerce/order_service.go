package woocommerce

import (
	"log"
	"log/slog"
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
	p            port.WoocommerceRepository
	s            port.ProjectRepository
	extensionSrv port.ExtensionService
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(woorepo port.WoocommerceRepository, projrepo port.ProjectRepository, extensionSrv port.ExtensionService) *OrderService {
	return &OrderService{
		p:            woorepo,
		s:            projrepo,
		extensionSrv: extensionSrv,
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

// GetOrderCount retrieves the count of orders for a given project ID
func (os *OrderService) GetOrderCount(projectID string, orderStatus w.OrderStatus, timeRange string) (int64, error) {
	orderCount, err := os.p.GetOrderCount(projectID, orderStatus, timeRange)
	if err != nil {
		return 0, err
	}
	return orderCount, nil
}

// Get10LatestOrders retrieves the latest 10 orders for a given project ID
func (os *OrderService) Get10LatestOrders(ctx echo.Context, projectID string, orderStatus w.OrderStatus, results chan<- []*domain.OrderRecord, errors chan<- error) {
	orders, err := os.p.OrderFindByProjectID(projectID, 10, 1, orderStatus, "", "")
	if err != nil {
		errors <- err
	} else {
		results <- orders
	}
}

// GetOrdersCountBetweenOrEquals retrieves the count of orders between or equal to a given date for a given project ID
func (os *OrderService) GetOrdersCountBetweenOrEquals(projectID string, timePeriod time.Time, orderStatus w.OrderStatus) (int64, error) {
	orderCount, err := os.p.GetOrdersCountBetweenOrEquals(projectID, timePeriod, w.OrderStatusCompleted)
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

// GetAllOrdersFromWoocommerce retrieves all orders from WooCommerce and saves them to MongoDB
func (os *OrderService) GetAllOrdersFromWoocommerce(customerKey string, customerSecret string, domainURL string, projectID string, totalProduct int64) error {
	client := InitClient(customerKey, customerSecret, domainURL)

	// Create all webhooks
	err := os.createAndSaveAllCustomers(client, projectID, totalProduct)
	if err != nil {
		slog.Error("get all order error", "error", err)
		return errors.Wrap(err, "create all order error")
	}

	slog.Info("get all orders success")
	return nil
}

// createAndSaveAllWebhooks creates WooCommerce products and saves results to MongoDB concurrently
func (os *OrderService) createAndSaveAllCustomers(client *commerce.Client, projectID string, totalOrder int64) error {

	var wg sync.WaitGroup
	orderCh := make(chan *w.OrderRecord, totalOrder) // Channel to distribute order to workers
	errorCh := make(chan *w.OrderRecord, 1)          // Buffered channel for error results

	// Worker pool to process products
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(projectID string) {
			defer wg.Done()
			for order := range orderCh {
				err := os.saveOrderResult(order, order.OrderID)
				if err != nil {
					log.Printf("Failed to save order result: %v", err)
					// Handle or log the error accordingly
				}
			}
		}(projectID)
	}

	// Goroutine to save error results
	go func(projectID string) {
		for result := range errorCh {
			err := os.saveOrderResult(result, result.OrderID)
			if err != nil {
				log.Printf("Failed to save order result: %v", err)
				// Handle or log the error accordingly
			}
		}
	}(projectID)

	// Main processing goroutine
	page := 1
	for {
		options := commerce.OrderListOption{
			ListOptions: commerce.ListOptions{
				Context: "view",
				Order:   "desc",
				Orderby: "id",
				Page:    page,
				PerPage: batchSize,
			},
		}
		resp, err := client.Order.List(options)
		if err != nil {
			errorCh <- &w.OrderRecord{
				ProjectID: projectID,
				Event:     "order.List",
				Error:     err.Error(),
				CreatedAt: time.Now().UTC(),
				OrderID:   0,
				Order:     commerce.Order{}, // Assuming empty product
				IsActive:  false,
			}
			break // Exit the loop on error
		}
		if len(resp) == 0 {
			break // Exit the loop if no more products are returned
		}
		os.extensionSrv.UpdateSynchronizerOrderReceivedExtension(nil, projectID, len(resp))
		for _, item := range resp {
			Status, _ := domain.StringToOrderStatus(item.Status)
			orderCh <- &w.OrderRecord{
				ProjectID: projectID,
				Event:     "order.created",
				Error:     "",
				CreatedAt: time.Now().UTC(),
				Timestamp: time.Now().UTC(),
				OrderID:   item.ID,
				Order:     item,
				IsActive:  true,
				Status:    Status,
			}

		}
		page++
	}

	close(orderCh) // Close the order channel after processing is complete
	close(errorCh) // Close the error channel after processing is complete

	wg.Wait() // Wait for all worker goroutines to finish

	return nil
}

// saveWebhookResult saves webhook creation result to MongoDB
func (os *OrderService) saveOrderResult(data *w.OrderRecord, orderID int64) error {

	err := os.p.OrderUpdate(data, orderID)
	if err != nil {
		return errors.Wrap(err, "failed to insert order result into MongoDB")
	}
	return nil
}
