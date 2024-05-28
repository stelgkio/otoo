package woocommerce

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/pkg/errors"
	"github.com/stelgkio/otoo/internal/core/auth"
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

type WoocommerceService struct {
}

func NewWoocommerceService() *WoocommerceService {
	return &WoocommerceService{}
}

func (s *WoocommerceService) WoocommerceCreateOrderWebHook(customerKey string, customerSecret string, domainUrl string, projectId string) error {
	client := initClient(customerKey, customerSecret, domainUrl)

	webhook := initOrderWebhook(projectId)
	_, err := client.Webhook.Create(webhook)
	if err != nil {
		slog.Error("create webhook error", "error", err)
		return errors.Wrap(err, "create webhook error")
	}
	slog.Info("create webhook success")
	return nil
}

// initClient init woocommerce client
func initClient(customerKey string, customerSecret string, domainUrl string) *woocommerce.Client {
	app := woocommerce.App{
		CustomerKey:    customerKey,
		CustomerSecret: customerSecret,
		AppName:        "otoo",
		//UserId:         "1",
		Scope: "read_write",
	}

	client := woocommerce.NewClient(app, domainUrl,
		woocommerce.WithLog(&woocommerce.LeveledLogger{
			Level: woocommerce.LevelDebug, // you should open this for debug in dev environment,  usefully.
		}),
		woocommerce.WithRetry(3),
		woocommerce.WithVersion("v3"),
	)

	//req,err :=client.NewRequest("GET", "/wp-json/wc/v3/products", nil,nil)
	return client
}

// initOrderWebhook init order webhook
func initOrderWebhook(projectId string) woocommerce.Webhook {
	timeNowStr := fmt.Sprintf("%d", time.Now().Unix())
	secretToken, _, err := auth.GenerateWebHookAccessToken(projectId)
	if err != nil {
		slog.Error("generate project id error: %v", err)
	}
	webhook := woocommerce.Webhook{
		Name:        "order create" + timeNowStr,
		Topic:       "order.created",
		DeliveryUrl: "http://localhost:8081/woocommerce/create", // your callback url for wooCommerce event cron job to notify
		Secret:      secretToken,
	}
	return webhook
}
