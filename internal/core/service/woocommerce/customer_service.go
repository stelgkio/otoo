package woocommerce

import (
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	woo "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/woocommerce"
	commerce "github.com/stelgkio/woocommerce"
)

// CustomerService represents the service for managing customers
type CustomerService struct {
	p port.WoocommerceRepository
	s port.ProjectRepository
}

// NewCustomerService creates a new CustomerService instance
func NewCustomerService(woorepo port.WoocommerceRepository, projrepo port.ProjectRepository) *CustomerService {
	return &CustomerService{
		p: woorepo,
		s: projrepo,
	}
}

// ExtractCustomerFromOrderAndUpsert extracts customer from order and upserts customer to MongoDB
func (c *CustomerService) ExtractCustomerFromOrderAndUpsert(ctx echo.Context, req *woo.OrderRecord) error {
	ctm, err := c.p.CustomerFindByEmail(req.ProjectID, req.Order.Billing.Email)
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
			if err != nil {
				slog.Error("error customerUpdate", "error", err)
				return err
			}
		}
	} else {
		customer := &woo.CustomerRecord{
			ProjectID:  req.ProjectID,
			Event:      "customer.created",
			CustomerID: req.Order.CustomerId,
			Email:      req.Order.Billing.Email,
			Timestamp:  time.Now().UTC(),
			IsActive:   true,
			CreatedAt:  time.Now().UTC(),
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
		err = c.p.CustomerCreate(customer, req.Order.Billing.Email)
		if err != nil {
			slog.Error("error customerCreate", "error", err)
			return err
		}
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

// FindCustomerByProjectIDAsync finds customer by project ID asynchronously
func (c *CustomerService) FindCustomerByProjectIDAsync(projectID string, size, page int, sort, direction string, results chan<- []*domain.CustomerRecord, errors chan<- error) {
	products, err := c.p.CustomerFindByProjectID(projectID, size, page, sort, direction)
	if err != nil {
		errors <- err
	}
	results <- products
}

// GetAllCustomerFromWoocommerce retrieves all customers from WooCommerce and saves them to MongoDB
func (c *CustomerService) GetAllCustomerFromWoocommerce(customerKey string, customerSecret string, domainURL string, projectID string, totalProduct int64) error {
	client := InitClient(customerKey, customerSecret, domainURL)

	// Create all webhooks
	err := c.createAndSaveAllCustomers(client, projectID, totalProduct)
	if err != nil {
		slog.Error("get all customers error", "error", err)
		return errors.Wrap(err, "create all customer error")
	}

	slog.Info("get all customers success")
	return nil
}

// createAndSaveAllWebhooks creates WooCommerce products and saves results to MongoDB concurrently
func (c *CustomerService) createAndSaveAllCustomers(client *commerce.Client, projectID string, totalCustomer int64) error {

	var wg sync.WaitGroup
	customerCh := make(chan *w.CustomerRecord, totalCustomer) // Channel to distribute customer to workers
	errorCh := make(chan *w.CustomerRecord, 1)                // Buffered channel for error results

	// Worker pool to process products
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(projectID string) {
			defer wg.Done()
			for customer := range customerCh {
				err := c.saveCustomerResult(customer, customer.Email)
				if err != nil {
					log.Printf("Failed to save customer result: %v", err)
					// Handle or log the error accordingly
				}
			}
		}(projectID)
	}

	// Goroutine to save error results
	go func(projectID string) {
		for result := range errorCh {
			err := c.saveCustomerResult(result, projectID)
			if err != nil {
				log.Printf("Failed to save customer result: %v", err)
				// Handle or log the error accordingly
			}
		}
	}(projectID)

	// Main processing goroutine
	page := 1
	for {
		options := commerce.CustomerListOption{
			ListOptions: commerce.ListOptions{
				Context: "view",
				Order:   "desc",
				Orderby: "id",
				Page:    page,
				PerPage: batchSize,
			},
		}
		resp, err := client.Customer.List(options)
		if err != nil {
			errorCh <- &w.CustomerRecord{
				ProjectID:  projectID,
				Event:      "customer.List",
				Error:      err.Error(),
				CreatedAt:  time.Now().UTC(),
				CustomerID: 0,
				Customer:   commerce.Customer{}, // Assuming empty product
				IsActive:   false,
			}
			break // Exit the loop on error
		}
		if len(resp) == 0 {
			break // Exit the loop if no more products are returned
		}
		for _, item := range resp {
			customerCh <- &w.CustomerRecord{
				ProjectID:  projectID,
				Event:      "customer.created",
				Error:      "",
				Email:      item.Email,
				CreatedAt:  time.Now().UTC(),
				Timestamp:  time.Now().UTC(),
				CustomerID: item.ID,
				Customer:   item,
				IsActive:   true,
			}

		}
		page++
	}

	close(customerCh) // Close the customer channel after processing is complete
	close(errorCh)    // Close the error channel after processing is complete

	wg.Wait() // Wait for all worker goroutines to finish

	return nil
}

// saveWebhookResult saves webhook creation result to MongoDB
func (c *CustomerService) saveCustomerResult(data *w.CustomerRecord, email string) error {

	err := c.p.CustomerUpdate(data, email)
	if err != nil {
		return errors.Wrap(err, "failed to insert customer result into MongoDB")
	}
	return nil
}
