package woocommerce

import (
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
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

type WoocommerceWebhookService struct {
	p port.WoocommerceRepository
}

func NewWoocommerceWebhookService(repo port.WoocommerceRepository) *WoocommerceWebhookService {
	return &WoocommerceWebhookService{
		repo,
	}
}
func (s *WoocommerceWebhookService) WoocommerceCreateAllWebHook(customerKey string, customerSecret string, domainUrl string, projectId uuid.UUID) error {
	client := initClient(customerKey, customerSecret, domainUrl)

	// Create all webhooks
	err := s.createAndSaveAllWebhooks(client, projectId)
	if err != nil {
		slog.Error("create all webhooks error", "error", err)
		return errors.Wrap(err, "create all webhooks error")
	}

	slog.Info("create all webhooks success")
	return nil
}

// initClient init woocommerce client
func initClient(customerKey string, customerSecret string, domainUrl string) *woocommerce.Client {
	app := woocommerce.App{
		CustomerKey:    customerKey,
		CustomerSecret: customerSecret,
		AppName:        "otoo",
		Scope:          "read_write",
	}

	client := woocommerce.NewClient(app, domainUrl,
		woocommerce.WithLog(&woocommerce.LeveledLogger{
			Level: woocommerce.LevelDebug, // open this for debug in dev environment
		}),
		woocommerce.WithRetry(3),
		woocommerce.WithVersion("v3"),
	)

	return client
}

// createAndSaveAllWebhooks creates WooCommerce webhooks and saves results to MongoDB concurrently
func (s *WoocommerceWebhookService) createAndSaveAllWebhooks(client *woocommerce.Client, projectId uuid.UUID) error {

	var wg sync.WaitGroup
	successCh := make(chan *w.WebhookRecord, len(allEvents)*3) // Buffered channel for successful results
	errorCh := make(chan *w.WebhookRecord, len(allEvents)*3)   // Buffered channel for error results

	for category, events := range allEvents {
		for _, event := range events {
			wg.Add(1)
			go func(cat, evt string) {
				defer wg.Done()
				webhook := initWebhook(evt, projectId)
				resp, err := client.Webhook.Create(webhook)
				if err != nil {

					errorCh <- &w.WebhookRecord{
						ProjectID: projectId.String(),
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
					ProjectID: projectId.String(),
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

// initWebhook initializes a webhook for a specific event
func initWebhook(event string, projectId uuid.UUID) woocommerce.Webhook {
	webhook := woocommerce.Webhook{
		Name:        "Otoo:" + event,
		Topic:       event,
		DeliveryUrl: os.Getenv("DELIVERY_URL") + "/woocommerce/" + strings.Replace(event, ".", "/", -1), // your callback URL
		Secret:      projectId.String(),
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
