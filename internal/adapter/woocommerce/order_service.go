package woocommerce

import (
	"time"

	"github.com/labstack/echo/v4"

	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
)

// OrderService represents the service for managing orders
type OrderService struct {
	p port.WoocommerceRepository
	s port.ProjectRepository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(woorepo port.WoocommerceRepository, projrepo port.ProjectRepository) *OrderService {
	return &OrderService{
		p: woorepo,
		s: projrepo,
	}
}

// GetOrderCountAsync retrieves the count of orders for a given project ID
func (os *OrderService) GetOrderCountAsync(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error) {
	orderCount, err := os.p.GetOrderCount(projectID)
	if err != nil {
		errors <- err
	}
	results <- orderCount
}

// GetOrderCount retrieves the count of orders for a given project ID
func (os *OrderService) GetOrderCount(projectID string) (int64, error) {
	orderCount, err := os.p.GetOrderCount(projectID)
	if err != nil {
		return 0, err
	}
	return orderCount, nil
}

// Get10LatestOrders retrieves the latest 10 orders for a given project ID
func (os *OrderService) Get10LatestOrders(ctx echo.Context, projectID string, results chan<- []*domain.OrderRecord, errors chan<- error) {
	orders, err := os.p.OrderFindByProjectID(projectID, 10, 1)
	if err != nil {
		errors <- err
	}
	results <- orders
}

// GetOrdersCountBetweenOrEquals retrieves the count of orders between or equal to a given date for a given project ID
func (os *OrderService) GetOrdersCountBetweenOrEquals(projectID string, timePeriod time.Time) (int64, error) {
	orderCount, err := os.p.GetOrdersCountBetweenOrEquals(projectID, timePeriod)
	if err != nil {
		return 0, err
	}
	return orderCount, nil

}
