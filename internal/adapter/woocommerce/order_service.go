package woocommerce

import (
	"github.com/labstack/echo/v4"

	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
)

type OrderService struct {
	p port.WoocommerceRepository
	s port.ProjectRepository
}


func NewOrderService(woorepo port.WoocommerceRepository ,projrepo port.ProjectRepository) *OrderService {
	return &OrderService{
		p: woorepo,
		s: projrepo,
	}
}


func (os *OrderService) GetOrderCount(ctx echo.Context, projectId string ,results chan<- int64, errors chan<- error) {
	orderCount, err := os.p.GetOrderCount( projectId)
	if err != nil {
		errors <- err
	}
	results <- orderCount
}


func (os *OrderService) Get10LatestOrders(ctx echo.Context, projectId string ,results chan<- []*domain.OrderRecord, errors chan<- error) {
	orders, err := os.p.OrderFindByProjectId( projectId, 10, 1 )
	if err != nil {
		errors <- err
	}
	results <- orders
}