package woocommerce

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
	woo "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/woocommerce"
)

type CustomerService struct {
	p port.WoocommerceRepository
	s port.ProjectRepository
}

func NewCustomerService(woorepo port.WoocommerceRepository, projrepo port.ProjectRepository) *CustomerService {
	return &CustomerService{
		p: woorepo,
		s: projrepo,
	}
}

// ExtractCustomerFromOrderAndUpsert extracts customer from order and upserts customer to MongoDB
func (c *CustomerService) ExtractCustomerFromOrderAndUpsert(ctx echo.Context, req *woo.OrderRecord) error {
	ctm, err := c.p.CustomerFindByEmail(req.Order.Billing.Email)
	if err != nil {
		slog.Error("error customerFindByEmail", "error", err)
		return err
	}

	if ctm != nil {
		// Check if the order ID already exists in the Orders slice
		orderExists := false
		for _, orderID := range ctm.Orders {
			if orderID == req.Order.ID {
				orderExists = true
				break
			}
		}
		if !orderExists {
			ctm.Orders = append(ctm.Orders, req.Order.ID)
			err = c.p.CustomerUpdate(ctm, req.Order.Billing.Email)
		}
	} else {
		customer := &woo.CustomerRecord{
			ProjectID:  req.ProjectID,
			Event:      "customer_created",
			CustomerID: req.Order.CustomerId,
			Email:      req.Order.Billing.Email,
			Timestamp:  time.Now(),
			IsActive:   true,
			CreatedAt:  time.Now(),
			Customer: woocommerce.Customer{
				ID:        req.Order.CustomerId,
				Email:     req.Order.Billing.Email,
				FirstName: req.Order.Billing.FirstName,
				LastName:  req.Order.Billing.LastName,
				Billing: &woocommerce.Billing{
					FirstName: req.Order.Billing.FirstName,
					LastName:  req.Order.Billing.LastName,
					Company:   req.Order.Billing.Company,
					Address1:  req.Order.Billing.Address1,
					Address2:  req.Order.Billing.Address2,
					City:      req.Order.Billing.City,
					State:     req.Order.Billing.State,
					PostCode:  req.Order.Billing.PostCode,
					Country:   req.Order.Billing.Country,
					Email:     req.Order.Billing.Email,
					Phone:     req.Order.Billing.Phone,
				},
				Shipping: &woocommerce.Shipping{
					FirstName: req.Order.Shipping.FirstName,
					LastName:  req.Order.Shipping.LastName,
					Company:   req.Order.Shipping.Company,
					Address1:  req.Order.Shipping.Address1,
					Address2:  req.Order.Shipping.Address2,
					City:      req.Order.Shipping.City,
					State:     req.Order.Shipping.State,
				},
			},
		}
		customer.Orders = []int64{req.Order.ID}
		err = c.p.CustomerCreate(customer)
	}

	return err
}

// GetCustomerCount retrieves the count of customers for a given project ID
func (c *CustomerService) GetCustomerCount(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error) {
	customerCount, err := c.p.GetCustomerCount(projectID)
	if err != nil {
		errors <- err
	}
	results <- customerCount
}
