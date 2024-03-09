package port

type WoocommerceRepository interface {
	// InsertWoocommerceOrder inserts a new order into the database
	InsertWoocommerceOrder(data any) error
}

type WoocommerceService interface {
	// WoocommerceCreateOrderWebHook create new order web hook for woocommerce
	WoocommerceCreateOrderWebHook(customerKey string, customerSecret string, domainUrl string, projectId string) error
}
