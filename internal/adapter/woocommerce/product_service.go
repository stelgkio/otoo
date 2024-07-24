package woocommerce

import (
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	woo "github.com/stelgkio/woocommerce"
)

const (
	workerCount = 100 // Number of worker goroutines
	batchSize   = 100 // Number of products to process per batch
)
type ProductService struct {
	p port.WoocommerceRepository
	s port.ProjectRepository
}


func NewProductService(woorepo port.WoocommerceRepository ,projrepo port.ProjectRepository) *ProductService {
	return &ProductService{
		woorepo,
		projrepo,
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
func (s *ProductService) createAndSaveAllProducts(client *woo.Client, projectId uuid.UUID) error {

	var wg sync.WaitGroup
	productCh := make(chan *w.ProductRecord, batchSize) // Channel to distribute products to workers
	errorCh := make(chan *w.ProductRecord, 1)   // Buffered channel for error results

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
				ProjectID: projectId.String(),
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
				ProjectID: projectId.String(),
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




func (ps *ProductService) ExtractProductFromOrderAndUpsert(ctx echo.Context, req *w.OrderRecord) error {
	panic("unimplemented")
}
