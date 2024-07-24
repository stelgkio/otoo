package woocommerce

import "github.com/stelgkio/woocommerce"

// initClient init woocommerce client
func InitClient(customerKey string, customerSecret string, domainUrl string) *woocommerce.Client {
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