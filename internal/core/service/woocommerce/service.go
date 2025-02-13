package woocommerce

import "github.com/stelgkio/woocommerce"

const (
	workerCount = 1   // Number of worker goroutines
	batchSize   = 100 // Number of products to process per batch
)

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
			Level: woocommerce.LevelError, // open this for debug in dev environment
		}),
		woocommerce.WithRetry(3),
		woocommerce.WithVersion("v3"),
	)

	return client
}
