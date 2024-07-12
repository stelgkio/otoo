package woocommerce

import (
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
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

type WoocommerceService struct {
}

func NewWoocommerceService() *WoocommerceService {
	return &WoocommerceService{}
}
func (s *WoocommerceService) WoocommerceCreateAllWebHook(customerKey string, customerSecret string, domainUrl string, projectId string) error {
	client := initClient(customerKey, customerSecret, domainUrl)

	// Create all webhooks
	err := createAllWebhooks(client, projectId)
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

// createAllWebhooks creates webhooks for all events concurrently
func createAllWebhooks(client *woocommerce.Client, projectId string) error {

	var wg sync.WaitGroup
	errorChannel := make(chan error, len(allEvents)*len(couponEvents)) // Maximum possible errors

	for category, events := range allEvents {
		for _, event := range events {
			wg.Add(1)
			go func(category, event string) {
				defer wg.Done()
				webhook := initWebhook(event, projectId)
				_, err := client.Webhook.Create(webhook)
				if err != nil {
					slog.Error("create webhook error", "category", category, "event", event, "error", err)
					errorChannel <- errors.Wrapf(err, "create webhook error for %s event %s", category, event)
				} else {
					slog.Info("create webhook success", "category", category, "event", event)
				}
			}(category, event)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errorChannel)

	// Collect errors
	var err error
	for e := range errorChannel {
		err = e // Return the last error encountered
	}

	return err
}

// initWebhook initializes a webhook for a specific event
func initWebhook(event string, projectId string) woocommerce.Webhook {
	webhook := woocommerce.Webhook{
		Name:        "Otoo:" + event,
		Topic:       event,
		DeliveryUrl: os.Getenv("DELIVERY_URL") + "/woocommerce/" + strings.Replace(event, ".", "/", -1), // your callback URL
		Secret:      projectId,
	}
	return webhook
}
