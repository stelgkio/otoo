package port

import (
	"time"

	"github.com/labstack/echo/v4"
	d "github.com/stelgkio/otoo/internal/core/domain"
	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	woo "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
)

// WoocommerceRepository defines the methods for interacting with the Woocommerce repository
type WoocommerceRepository interface {
	// InsertWoocommerceOrder inserts a new order into the database
	OrderCreate(data *w.OrderRecord) error
	OrderUpdate(data *w.OrderRecord, orderID int64) error
	OrderDelete(orderID int64, projectID string) error
	GetOrderByID(projectID string, orderID int64) (*w.OrderRecord, error)
	OrderFindByProjectID(projectID string, size, page int, orderStatus w.OrderStatus, sort, direction string) ([]*w.OrderRecord, error)
	OrderFindByProjectIDWithTimePedio(projectID string, size, page int, orderStatus w.OrderStatus, sort, direction string, timeperiod time.Time) ([]*w.OrderRecord, error)
	GetOrderCount(projectID string, orderStatus w.OrderStatus, timeRange string) (int64, error)
	GetOrdersCountBetweenOrEquals(projectID string, timeperiod time.Time, orderStatus w.OrderStatus) (int64, error)

	CustomerCreate(data *w.CustomerRecord, email string) error
	CustomerUpdate(data *w.CustomerRecord, email string) error
	CustomerDelete(customerID int64, projectID string) error
	CustomerFindByProjectID(projectID string, size, page int, sort, direction string) ([]*w.CustomerRecord, error)
	CustomerFindByEmail(projectID string, email string) (*w.CustomerRecord, error)
	GetCustomerCount(projectID string) (int64, error)

	ProductCreate(data *w.ProductRecord) error
	ProductDelete(productID int64, projectID string) error
	ProductUpdate(data *w.ProductRecord, productID int64, projectID string) error
	ProductFindByProjectID(projectID string, size, page int, sort, direction string, productType w.ProductType) ([]*w.ProductRecord, error)
	GetProductCount(projectID string, productType w.ProductType) (int64, error)
	GetProductByID(projectID string, orderID int64) (*w.ProductRecord, error)
	ProductBestSellerAggregate(projectID string) ([]bson.M, error)

	CouponCreate(data any) error
	CouponUpdate(data any) error
	CouponDelete(data any) error
	CouponFindByProjectID(projectID string) error

	WebhookCreate(data w.WebhookRecord) error
	WebhookUpdate(data w.WebhookRecord) (*w.WebhookRecord, error)
	WebhookBatchDelete(projectID string) error
	WebhookDelete(projectID string, webhookID int64) error
	WebhookFindByProjectID(projectID string) ([]w.WebhookRecord, error)
	WebhookCount(projectID string) (int64, error)
}

// WoocommerceWebhookService defines the methods for interacting with the Woocommerce service
type WoocommerceWebhookService interface {
	// WoocommerceCreateOrderWebHook create new order web hook for woocommerce
	WoocommerceCreateAllWebHook(customerKey string, customerSecret string, domainURL string, projectID string) error
	WoocommerceCreateAllWebHookAsync(customerKey string, customerSecret string, domainURL string, projectID string) error
	FindWebhookByProjectIDAsync(ctx echo.Context, projectID string, results chan<- []domain.WebhookRecord, errors chan<- error)
	GetWebhookCountAsync(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error)
	DeleteAllWebhooksByProjectID(projectID string, customerKey string, customerSecret string, domainURL string) error
}

// CustomerService defines the methods for interacting with the Woocommerce service
type CustomerService interface {
	ExtractCustomerFromOrderAndUpsert(ctx echo.Context, req *woo.OrderRecord) error
	GetCustomerCount(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error)
	FindCustomerByProjectIDAsync(projectID string, size, page int, sort, direction string, results chan<- []*domain.CustomerRecord, errors chan<- error)
	GetAllCustomerFromWoocommerce(customerKey string, customerSecret string, domainURL string, projectID string, totalProduct int64) error
}

// ProductService defines the methods for interacting with the Product service
type ProductService interface {
	ExtractProductFromOrderAndUpsert(ctx echo.Context, req *woo.OrderRecord, project *d.Project) error
	GetAllProductFromWoocommerce(ccustomerKey string, customerSecret string, domainURL string, projectID string, totalProduct int64) error
	GetProductCount(ctx echo.Context, projectID string, productType w.ProductType, results chan<- int64, errors chan<- error)
	GetProductBestSeller(projectID string, totalCount int64, results chan<- []*domain.ProductBestSellerRecord, errors chan<- error)
	FindProductByProjectIDAsync(projectID string, size, page int, sort, direction string, productType w.ProductType, results chan<- []*domain.ProductRecord, errors chan<- error)
}

// OrderService defines the methods for interacting with the Order service
type OrderService interface {
	Get10LatestOrders(ctx echo.Context, projectID string, orderStatus w.OrderStatus, sort string, results chan<- []*domain.OrderRecord, errors chan<- error)
	GetOrderCountAsync(ctx echo.Context, projectID string, orderStatus w.OrderStatus, timeRange string, results chan<- int64, errors chan<- error)
	GetOrdersCountBetweenOrEquals(projectID string, timePeriod time.Time, orderStatus w.OrderStatus) (int64, error)
	GetOrderCount(projectID string, orderStatus w.OrderStatus, timeRange string) (int64, error)
	GetOrderByID(projectID string, orderID int64) (*domain.OrderRecord, error)
	FindOrderByProjectIDAsync(projectID string, size, page int, orderStatus w.OrderStatus, sort, direction string, results chan<- []*domain.OrderRecord, errors chan<- error)
	FindOrderByProjectIDWithTimePeriodAsync(projectID string, size, page int, orderStatus w.OrderStatus, sort, direction string, timePeriod time.Time, results chan<- []*domain.OrderRecord, errors chan<- error)
	GetAllOrdersFromWoocommerce(customerKey string, customerSecret string, domainURL string, projectID string, totalProduct int64) error
	UpdateOrderStatusByID(projectID string, orderID int64, status string, project *d.Project) (*domain.OrderRecord, error)
	BatchUpdateOrdersStatus(projectID string, orders []int64, status string, proj *d.Project) ([]*domain.OrderRecord, error)
	UpdateOrder(projectID string, orderID int64, orderTable *domain.OrderTableList, proj *d.Project) (*domain.OrderRecord, error)
	GetLatestOrderWeeklyBalance(ctx echo.Context, projectID string, results chan<- *domain.WeeklyAnalytics, errors chan<- error)
}

// ReportService defines the methods for interacting with the Report service
type ReportService interface {
	GetCustomerTotalCount(ctx echo.Context, projectID string) (int, error)
	GetOrderTotalCount(ctx echo.Context, projectID string) (int, error)
	GetProductTotalCount(ctx echo.Context, projectID string) (int, error)
	GetCustomerTotalCountTestCredential(ctx echo.Context, customerKey string, customerSecret string, domainURL string) (int, error)
}
