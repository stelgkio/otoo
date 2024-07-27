package woocommerce

import (
	"fmt"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/woocommerce"
	woo "github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	workerCount = 100 // Number of worker goroutines
	batchSize   = 100 // Number of products to process per batch
)

type ProductService struct {
	p port.WoocommerceRepository
	s port.ProjectRepository
}

func NewProductService(woorepo port.WoocommerceRepository, projrepo port.ProjectRepository) *ProductService {
	return &ProductService{
		p: woorepo,
		s: projrepo,
	}
}

func (s *ProductService) GetAllProductFromWoocommerce(customerKey string, customerSecret string, domainUrl string, projectId uuid.UUID) error {
	client := InitClient(customerKey, customerSecret, domainUrl)

	// Create all webhooks
	err := s.createAndSaveAllProducts(client, projectId)
	if err != nil {
		slog.Error("create all products error", "error", err)
		return errors.Wrap(err, "create all products error")
	}

	slog.Info("create all products success")
	return nil
}

// createAndSaveAllWebhooks creates WooCommerce products and saves results to MongoDB concurrently
func (s *ProductService) createAndSaveAllProducts(client *woo.Client, projectID uuid.UUID) error {

	var wg sync.WaitGroup
	productCh := make(chan *w.ProductRecord, batchSize) // Channel to distribute products to workers
	errorCh := make(chan *w.ProductRecord, 1)           // Buffered channel for error results

	// Worker pool to process products
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for product := range productCh {
				err := s.saveWebhookResult(product)
				if err != nil {
					log.Printf("Failed to save webhook result: %v", err)
					// Handle or log the error accordingly
				}
			}
		}()
	}

	// Goroutine to save error results
	go func() {
		for result := range errorCh {
			err := s.saveWebhookResult(result)
			if err != nil {
				log.Printf("Failed to save webhook result: %v", err)
				// Handle or log the error accordingly
			}
		}
	}()

	// Main processing goroutine
	page := 1
	for {
		options := woo.ProductListOptions{
			ListOptions: woo.ListOptions{
				Context: "view",
				Order:   "desc",
				Orderby: "date",
				Page:    page,
				PerPage: batchSize,
			},
		}
		resp, err := client.Product.List(options)
		if err != nil {
			errorCh <- &w.ProductRecord{
				ProjectID: projectID.String(),
				Event:     "Product.List",
				Error:     err.Error(),
				CreatedAt: time.Now(),
				ProductID: 0,
				Product:   woo.Product{}, // Assuming empty product
				IsActive:  false,
			}
			break // Exit the loop on error
		}
		if len(resp) == 0 {
			break // Exit the loop if no more products are returned
		}
		for _, item := range resp {
			productCh <- &w.ProductRecord{
				ProjectID: projectID.String(),
				Event:     "Product.List",
				Error:     "",
				CreatedAt: time.Now(),
				ProductID: item.ID,
				Product:   item,
				IsActive:  true,
			}
		}
		page++
	}

	close(productCh) // Close the product channel after processing is complete
	close(errorCh)   // Close the error channel after processing is complete

	wg.Wait() // Wait for all worker goroutines to finish

	return nil
}

// saveWebhookResult saves webhook creation result to MongoDB
func (s *ProductService) saveWebhookResult(data *w.ProductRecord) error {

	err := s.p.ProductCreate(data)
	if err != nil {
		return errors.Wrap(err, "failed to insert product result into MongoDB")
	}
	return nil
}

// ExtractProductFromOrderAndUpsert extracts product from order and upsert to MongoDB
func (s *ProductService) ExtractProductFromOrderAndUpsert(ctx echo.Context, req *w.OrderRecord) error {
	for _, orderItem := range req.Order.LineItems {
		if orderItem.ProductID != 0 {
			product, err := s.p.GetProductByID(req.ProjectID, orderItem.ProductID)
			if err != nil {
				slog.Error("error GetProductByID", "error", err)
				return err
			}

			if product != nil {
				// Check if the order ID already exists in the Orders slice
				orderExists := false
				for _, orderID := range product.Orders {
					if orderID == req.Order.ID {
						orderExists = true
						break
					}
				}
				if !orderExists {
					product.Orders = append(product.Orders, req.Order.ID)
					err = s.p.ProductUpdate(product, product.ProductID)
				}
			}
		}
	}

	return nil
}

// GetProductCount gets product count from MongoDB
func (s *ProductService) GetProductCount(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error) {
	productCount, err := s.p.GetProductCount(projectID)
	if err != nil {
		errors <- err
	}
	results <- productCount
}

// GetProductBestSeller gets product best seller from MongoDB
func (s *ProductService) GetProductBestSeller(projectID string, totalCount int64, results chan<- []*domain.ProductBestSellerRecord, errors chan<- error) {
	products, err := s.p.ProductBestSellerAggregate(projectID)
	if err != nil {
		errors <- err
	}
	var bestSellers []*domain.ProductBestSellerRecord
	for _, product := range products {
		productID, ok := product["productId"].(int64)
		if !ok {
			slog.Info("ProductID is not of type int64: %v", product["productId"])
			continue
		}

		orderCount, ok := product["orderCount"].(int32)
		if !ok {
			slog.Error("OrderCount is not of type int: %v", product["orderCount"])
			continue
		}

		productData, ok := product["product"].(bson.M)
		if !ok {
			slog.Error("Product is not of type bson.M: %v", product["product"])
			continue
		}

		var product woocommerce.Product
		productBytes, err := bson.MarshalExtJSON(productData["product"], false, false)
		if err != nil {
			slog.Error("Error marshalling product data", "error", err)
			continue
		}
		if err := bson.UnmarshalExtJSON(productBytes, false, &product); err != nil {
			slog.Error("Error unmarshalling product data", "error", err)
			continue
		}

		fmt.Printf("Product ID: %v, Order Count: %d, Name: %s\n", productID, orderCount, product.Name)

		newProductBestSellerRecord := &domain.ProductBestSellerRecord{
			ProjectID:   projectID,
			ProductID:   productID,
			TotalOrders: int64(orderCount),
			StartDate:   time.Now(),
			EndDate:     time.Now(),
			ProductName: product.Name,
			Timestamp:   time.Now(),
			IsActive:    true,
		}
		newProductBestSellerRecord.CalculatePercentages(totalCount)
		bestSellers = append(bestSellers, newProductBestSellerRecord)
	}
	if err != nil {
		errors <- err
	}
	results <- bestSellers
}
