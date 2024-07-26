package port

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	woo "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WoocommerceRepository defines the methods for interacting with the Woocommerce repository
type WoocommerceRepository interface {
	// InsertWoocommerceOrder inserts a new order into the database
	OrderCreate(data *w.OrderRecord) error
	OrderUpdate(data *w.OrderRecord, orderID int64) error
	OrderDelete(data any) error
	OrderFindByProjectID(projectID string, size, page int) ([]*w.OrderRecord, error)
	GetOrderCount(projectID string, orderStatus w.OrderStatus) (int64, error)
	GetOrdersCountBetweenOrEquals(projectID string, timeperiod time.Time, orderStatus w.OrderStatus) (int64, error)

	CustomerCreate(data *w.CustomerRecord) error
	CustomerUpdate(data *w.CustomerRecord, email string) error
	CustomerDelete(data any) error
	CustomerFindByProjectID(projectID string) error
	CustomerFindByEmail(email string) (*w.CustomerRecord, error)
	GetCustomerCount(projectID string) (int64, error)

	ProductCreate(data *w.ProductRecord) error
	ProductDelete(productID int64) error
	ProductUpdate(data *w.ProductRecord, productID int64) error
	ProductFindByProjectID(projectID string) error
	GetProductCount(projectID string) (int64, error)
	GetProductByID(projectID string, orderID int64) (*w.ProductRecord, error)
	ProductBestSellerAggregate(projectID string) ([]bson.M, error)

	CouponCreate(data any) error
	CouponUpdate(data any) error
	CouponDelete(data any) error
	CouponFindByProjectID(projectID string) error

	WebhookCreate(data w.WebhookRecord) error
	WebhookUpdate(data w.WebhookRecord) (*w.WebhookRecord, error)
	WebhookDelete(id primitive.ObjectID) error
	WebhookFindByProjectID(projectID string) ([]w.WebhookRecord, error)
}

// WoocommerceWebhookService defines the methods for interacting with the Woocommerce service
type WoocommerceWebhookService interface {
	// WoocommerceCreateOrderWebHook create new order web hook for woocommerce
	WoocommerceCreateAllWebHook(customerKey string, customerSecret string, domainURL string, projectID uuid.UUID) error
}

// CustomerService defines the methods for interacting with the Woocommerce service
type CustomerService interface {
	ExtractCustomerFromOrderAndUpsert(ctx echo.Context, req *woo.OrderRecord) error
	GetCustomerCount(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error)
}

// ProductService defines the methods for interacting with the Product service
type ProductService interface {
	ExtractProductFromOrderAndUpsert(ctx echo.Context, req *woo.OrderRecord) error
	GetAllProductFromWoocommerce(ccustomerKey string, customerSecret string, domainURL string, projectID uuid.UUID) error
	GetProductCount(ctx echo.Context, projectID string, results chan<- int64, errors chan<- error)
	GetProductBestSeller(projectID string, results chan<- []*domain.ProductBestSellerRecord, errors chan<- error)
}

// OrderService defines the methods for interacting with the Order service
type OrderService interface {
	Get10LatestOrders(ctx echo.Context, projectID string, results chan<- []*domain.OrderRecord, errors chan<- error)
	GetOrderCountAsync(ctx echo.Context, projectID string, orderStatus w.OrderStatus, results chan<- int64, errors chan<- error)
	GetOrdersCountBetweenOrEquals(projectID string, timePeriod time.Time, orderStatus w.OrderStatus) (int64, error)
	GetOrderCount(projectID string, orderStatus w.OrderStatus) (int64, error)
	FindOrderByProjectIDAsync(projectID string, size, page int, results chan<- []*domain.OrderRecord, errors chan<- error)
}
