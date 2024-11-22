package woocommerce

import (
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/woocommerce"
)

var (
	couponEvents = []string{
		"coupon.created",
		"coupon.updated",
		"coupon.deleted",
	}

	customerEvents = []string{
		"customer.created",
		"customer.updated",
		"customer.deleted",
	}

	orderEvents = []string{
		"order.created",
		"order.updated",
		"order.deleted",
	}

	productEvents = []string{
		"product.created",
		"product.updated",
		"product.deleted",
	}
)
var allEvents = map[string][]string{
	"coupon":   couponEvents,
	"customer": customerEvents,
	"order":    orderEvents,
	"product":  productEvents,
}

// WoocommerceWebhookService implements the WoocommerceWebhookService interface
type WoocommerceWebhookService struct {
	p port.WoocommerceRepository
}

// NewWoocommerceWebhookService creates a new instance of WoocommerceWebhookService
func NewWoocommerceWebhookService(repo port.WoocommerceRepository) *WoocommerceWebhookService {
	return &WoocommerceWebhookService{
		repo,
	}
}

// WoocommerceCreateAllWebHook creates WooCommerce webhooks and saves results to MongoDB concurrently
func (s *WoocommerceWebhookService) WoocommerceCreateAllWebHook(customerKey string, customerSecret string, domainUrl string, projectID string) error {
	client := InitClient(customerKey, customerSecret, domainUrl)

	// Create all webhooks
	err := s.createAndSaveAllWebhooks(client, projectID)
	if err != nil {
		slog.Error("create all webhooks error", "error", err)
		return errors.Wrap(err, "create all webhooks error")
	}

	slog.Info("create all webhooks success")
	return nil
}

// WoocommerceCreateAllWebHookAsync creates WooCommerce webhooks and saves results to MongoDB concurrently
func (s *WoocommerceWebhookService) WoocommerceCreateAllWebHookAsync(customerKey string, customerSecret string, domainUrl string, projectID string) error {
	client := InitClient(customerKey, customerSecret, domainUrl)

	// Create all webhooks
	err := s.createAndSaveAllWebhooksAsync(client, projectID)
	if err != nil {
		slog.Error("create all webhooks error", "error", err)
		return errors.Wrap(err, "create all webhooks error")
	}

	slog.Info("create all webhooks success")
	return nil
}

// createAndSaveAllWebhooks creates WooCommerce webhooks and saves results to MongoDB concurrently
func (s *WoocommerceWebhookService) createAndSaveAllWebhooksAsync(client *woocommerce.Client, projectID string) error {

	var wg sync.WaitGroup
	successCh := make(chan *w.WebhookRecord, len(allEvents)*3) // Buffered channel for successful results
	errorCh := make(chan *w.WebhookRecord, len(allEvents)*3)   // Buffered channel for error results

	for category, events := range allEvents {
		for _, event := range events {
			wg.Add(1)
			go func(cat, evt string) {
				defer wg.Done()
				webhook := initWebhook(evt, projectID)
				resp, err := client.Webhook.Create(webhook)
				if err != nil {

					errorCh <- &w.WebhookRecord{
						ProjectID: projectID,
						Event:     evt,
						Error:     err.Error(),
						CreatedAt: time.Now().UTC(),
						WebhookID: resp.ID,
						Webhook:   *resp,
						IsActive:  true,
					}
					return
				}

				successCh <- &w.WebhookRecord{
					ProjectID: projectID,
					Event:     evt,
					Error:     "",
					CreatedAt: time.Now().UTC(),
					WebhookID: resp.ID,
					Webhook:   *resp,
					IsActive:  true,
				}
			}(category, event)
		}
	}

	go func() {
		wg.Wait()
		close(successCh)
		close(errorCh)
	}()

	// Save successful results
	go func() {
		for result := range successCh {
			err := s.saveWebhookResult(result)
			if err != nil {
				slog.Error("Failed to save webhook result", "error", err)
				// Handle or log the error accordingly
			}
		}
	}()
	// Save error results
	go func() {
		for result := range errorCh {
			err := s.saveWebhookResult(result)
			if err != nil {
				slog.Error("Failed to save webhook", "error", err)
				// Handle or log the error accordingly
			}
		}
	}()

	return nil
}

// createAndSaveAllWebhooks creates WooCommerce webhooks and saves results to MongoDB concurrently
func (s *WoocommerceWebhookService) createAndSaveAllWebhooks(client *woocommerce.Client, projectID string) error {

	var wg sync.WaitGroup
	successCh := make(chan *w.WebhookRecord, len(allEvents)*3) // Buffered channel for successful results
	errorCh := make(chan *w.WebhookRecord, len(allEvents)*3)   // Buffered channel for error results

	for category, events := range allEvents {
		for _, event := range events {
			wg.Add(1)
			go func(cat, evt string) {
				defer wg.Done()
				webhook := initWebhook(evt, projectID)
				resp, err := client.Webhook.Create(webhook)
				if err != nil {

					errorCh <- &w.WebhookRecord{
						ProjectID: projectID,
						Event:     evt,
						Error:     err.Error(),
						CreatedAt: time.Now().UTC(),
						WebhookID: resp.ID,
						Webhook:   *resp,
						IsActive:  true,
					}
					return
				}

				successCh <- &w.WebhookRecord{
					ProjectID: projectID,
					Event:     evt,
					Error:     "",
					CreatedAt: time.Now().UTC(),
					WebhookID: resp.ID,
					Webhook:   *resp,
					IsActive:  true,
				}
			}(category, event)
		}
	}

	go func() {
		wg.Wait()
		close(successCh)
		close(errorCh)
	}()

	// Save successful results

	for result := range successCh {
		err := s.saveWebhookResult(result)
		if err != nil {
			slog.Error("Failed to save webhook result", "error", err)
			// Handle or log the error accordingly
		}
	}

	// Save error results

	for result := range errorCh {
		err := s.saveWebhookResult(result)
		if err != nil {
			slog.Error("Failed to save webhook", "error", err)
			// Handle or log the error accordingly
		}
	}

	return nil
}

// initWebhook initializes a webhook for a specific event
func initWebhook(event string, projectID string) woocommerce.Webhook {
	webhook := woocommerce.Webhook{
		Name:        "Otoo:" + event,
		Topic:       event,
		DeliveryUrl: os.Getenv("DELIVERY_URL") + "/woocommerce/" + strings.Replace(event, ".", "/", -1), // your callback URL
		Secret:      projectID,
	}
	return webhook
}

// saveWebhookResult saves webhook creation result to MongoDB
func (s *WoocommerceWebhookService) saveWebhookResult(data *w.WebhookRecord) error {

	err := s.p.WebhookCreate(*data)
	if err != nil {
		return errors.Wrap(err, "failed to insert webhook result into MongoDB")
	}
	return nil
}

// GetWebhookCountAsync retrieves the number of webhooks for a given project ID
func (s *WoocommerceWebhookService) GetWebhookCountAsync(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error) {
	webhookCount, err := s.p.WebhookCount(projectID)
	if err != nil {
		errors <- err
	} else {
		results <- webhookCount
	}
}

// FindWebhookByProjectIDAsync retrieves webhooks for a given project ID
func (s *WoocommerceWebhookService) FindWebhookByProjectIDAsync(ctx echo.Context, projectID string, results chan<- []domain.WebhookRecord, errors chan<- error) {
	webhooks, err := s.p.WebhookFindByProjectID(projectID)
	if err != nil {
		errors <- err
	} else {
		results <- webhooks
	}
}

// DeleteAllWebhooksByProjectID deletes all webhooks for a given project ID
func (s *WoocommerceWebhookService) DeleteAllWebhooksByProjectID(projectID string, customerKey string, customerSecret string, domainURL string) error {
	webhooks, err := s.p.WebhookFindByProjectID(projectID)
	if err != nil {
		return err
	}
	client := InitClient(customerKey, customerSecret, domainURL)

	if len(webhooks) == 0 {
		return nil
	}
	for _, webhook := range webhooks {

		_, err := client.Webhook.Delete(webhook.WebhookID, woocommerce.DeleteOption{
			Force: true,
		})
		if err != nil {
			slog.Error("Error deleting webhook", "error", err)
			if err.Error() != "Invalid ID." {
				return err
			}

		}
		s.p.WebhookDelete(projectID, webhook.WebhookID)
	}

	return nil
}
