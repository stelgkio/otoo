package woocommerce

import (
	"fmt"
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
	"github.com/stelgkio/otoo/internal/core/util"
	"github.com/stelgkio/woocommerce"
	woo "github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
)

// ProductService represents the service for managing products
type ProductService struct {
	p            port.WoocommerceRepository
	s            port.ProjectRepository
	extensionSrv port.ExtensionService
}

// NewProductService creates a new ProductService instance
func NewProductService(woorepo port.WoocommerceRepository, projrepo port.ProjectRepository, extensionSrv port.ExtensionService) *ProductService {
	return &ProductService{
		p:            woorepo,
		s:            projrepo,
		extensionSrv: extensionSrv,
	}
}

// GetAllProductFromWoocommerce gets all products from WooCommerce and saves them to MongoDB
func (s *ProductService) GetAllProductFromWoocommerce(customerKey string, customerSecret string, domainURL string, projectID string, totalProduct int64) error {
	client := InitClient(customerKey, customerSecret, domainURL)

	// Create all webhooks
	err := s.createAndSaveAllProducts(client, projectID, totalProduct)
	if err != nil {
		slog.Error("get all products error", "error", err)
		return errors.Wrap(err, "create all products error")
	}

	slog.Info("get all products success")
	return nil
}

// createAndSaveAllWebhooks creates WooCommerce products and saves results to MongoDB concurrently
func (s *ProductService) createAndSaveAllProducts(client *woo.Client, projectID string, totalProduct int64) error {

	var wg sync.WaitGroup
	productCh := make(chan *w.ProductRecord, totalProduct) // Channel to distribute products to workers
	errorCh := make(chan *w.ProductRecord, 1)              // Buffered channel for error results

	// Worker pool to process products
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(projectID string) {
			defer wg.Done()
			for product := range productCh {
				err := s.saveProductResult(product, projectID)
				if err != nil {
					log.Printf("Failed to save webhook result: %v", err)
					// Handle or log the error accordingly
				}
			}
		}(projectID)
	}

	// Goroutine to save error results
	go func(projectID string) {
		for result := range errorCh {
			err := s.saveProductResult(result, projectID)
			if err != nil {
				log.Printf("Failed to save webhook result: %v", err)
				// Handle or log the error accordingly
			}
		}
	}(projectID)

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
			Status: "publish",
		}
		resp, err := client.Product.List(options)
		if err != nil {
			errorCh <- &w.ProductRecord{
				ProjectID: projectID,
				Event:     "product.List",
				Error:     err.Error(),
				CreatedAt: time.Now().UTC(),
				ProductID: 0,
				Product:   woo.Product{}, // Assuming empty product
				IsActive:  false,
			}
			break // Exit the loop on error
		}
		if len(resp) == 0 {
			break // Exit the loop if no more products are returned
		}
		s.extensionSrv.UpdateSynchronizerProductReceivedExtension(nil, projectID, len(resp))
		for _, item := range resp {
			productCh <- &w.ProductRecord{
				ProjectID: projectID,
				Event:     "product.created",
				Error:     "",
				CreatedAt: time.Now().UTC(),
				Timestamp: time.Now().UTC(),
				ProductID: item.ID,
				Product:   item,
				IsActive:  true,
			}
			if item.Type == domain.Variable.String() {
				for _, variationID := range item.Variations {
					variation, err := client.ProductVariation.Get(item.ID, variationID, nil)
					if err != nil {
						return err
					}

					variationRecord := &domain.ProductRecord{
						ProjectID: projectID,
						Error:     "",
						Event:     "product.created",
						ProductID: variation.ID,
						Product:   *variation,
						IsActive:  true,
						CreatedAt: time.Now().UTC(),
						Timestamp: time.Now().UTC(),
						ParentId:  item.ID,
					}

					err = s.p.ProductUpdate(variationRecord, variation.ID, projectID)
				}

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
func (s *ProductService) saveProductResult(data *w.ProductRecord, projectID string) error {

	err := s.p.ProductUpdate(data, data.ProductID, projectID)
	if err != nil {
		return errors.Wrap(err, "failed to insert product result into MongoDB")
	}
	return nil
}

// ExtractProductFromOrderAndUpsert upserts product data into the database
func (s *ProductService) ExtractProductFromOrderAndUpsert(ctx echo.Context, req *w.OrderRecord, proj *d.Project) error {
	client := InitClient(proj.WoocommerceProject.ConsumerKey, proj.WoocommerceProject.ConsumerSecret, proj.WoocommerceProject.Domain)
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
				if product.Orders != nil {
					for _, orderID := range product.Orders {
						if orderID == req.Order.ID {
							orderExists = true
							break
						}
					}
				}
				if !orderExists {
					if req.Order.Status != domain.OrderStatusCancelled.String() {
						if product.Orders == nil {
							product.Orders = []int64{req.Order.ID}
						} else {
							product.Orders = append(product.Orders, req.Order.ID)
						}
					} else {
						product.Orders = util.RemoveElement(product.Orders, util.FindIndex(product.Orders, req.Order.ID))
					}

					err = s.p.ProductUpdate(product, product.ProductID, req.ProjectID)
				}
			} else {
				wooproduct, err := client.Product.Get(orderItem.ProductID, nil)
				if err != nil {
					return err
				}
				productRecord := &domain.ProductRecord{
					ProjectID: proj.Id.String(),
					Error:     "",
					Event:     "product.created",
					ProductID: wooproduct.ID,
					Product:   *wooproduct,
					IsActive:  true,
					CreatedAt: time.Now().UTC(),
					Timestamp: time.Now().UTC(),
				}
				if req.Order.Status != domain.OrderStatusCancelled.String() {
					productRecord.Orders = []int64{req.Order.ID}
				}
				err = s.p.ProductUpdate(productRecord, wooproduct.ID, req.ProjectID)

				if wooproduct.Type == domain.Variable.String() {
					for _, variationID := range wooproduct.Variations {
						variation, err := client.ProductVariation.Get(wooproduct.ID, variationID, nil)
						if err != nil {
							return err
						}

						variationRecord := &domain.ProductRecord{
							ProjectID: proj.Id.String(),
							Error:     "",
							Event:     "product.created",
							ProductID: variation.ID,
							Product:   *variation,
							IsActive:  true,
							CreatedAt: time.Now().UTC(),
							Timestamp: time.Now().UTC(),
							ParentId:  wooproduct.ID,
						}

						err = s.p.ProductUpdate(variationRecord, variation.ID, req.ProjectID)
					}

				}
			}
		}

	}
	return nil
}

// GetProductCount gets product count from MongoDB
func (s *ProductService) GetProductCount(ctx echo.Context, projectID string, productType w.ProductType, results chan<- int64, errors chan<- error) {
	productCount, err := s.p.GetProductCount(projectID, productType)
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
			slog.Error("ProductID is not of type int64: %v", product["productId"])
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
			StartDate:   time.Now().UTC(),
			EndDate:     time.Now().UTC(),
			ProductName: product.Name,
			Timestamp:   time.Now().UTC(),
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

// FindProductByProjectIDAsync finds product by project ID asynchronously
func (s *ProductService) FindProductByProjectIDAsync(projectID string, size, page int, sort, direction string, productType w.ProductType, results chan<- []*domain.ProductRecord, errors chan<- error) {
	products, err := s.p.ProductFindByProjectID(projectID, size, page, sort, direction, productType)
	if err != nil {
		errors <- err
	}
	results <- products
}
