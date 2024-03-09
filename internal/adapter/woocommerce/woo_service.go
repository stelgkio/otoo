package woocommerce

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/chenyangguang/woocommerce"
	"github.com/stelgkio/otoo/internal/core/auth"
)

// const (
// 	customerKey    = "ck_8c344061bcfda558d2f114efb8d1b892b4330a73" // your customer_key
// 	customerSecret = "cs_562132bdb7e4e4c9e37b53a0fb703718e2dad5f7" // your customer_secret
// 	shopUrl        = "shop.gitvim.com"                             // your shop website domain
// )

// var client *woocommerce.Client

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
		slog.Error("create webhook error: %v", err)
		return err
	}
	slog.Info("create webhook success")
	return nil
}

// initClient init woocommerce client
func initClient(customerKey string, customerSecret string, domainUrl string) *woocommerce.Client {
	app := woocommerce.App{
		CustomerKey:    customerKey,
		CustomerSecret: customerSecret,
	}

	client := woocommerce.NewClient(app, domainUrl,
		woocommerce.WithLog(&woocommerce.LeveledLogger{
			Level: woocommerce.LevelDebug, // you should open this for debug in dev environment,  usefully.
		}),
		woocommerce.WithRetry(3),
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
