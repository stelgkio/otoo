package port

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	woo "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WoocommerceRepository interface {
	// InsertWoocommerceOrder inserts a new order into the database
	OrderCreate(data *w.OrderRecord) error
	OrderUpdate(data *w.OrderRecord, orderId int64) error
	OrderDelete(data any) error
	OrderFindByProjectId(projectId string) error

	CustomerCreate(data *w.CustomerRecord) error
	CustomerUpdate(data *w.CustomerRecord,email string) error
	CustomerDelete(data any) error
	CustomerFindByProjectId(projectId string) error
	CustomerFindByEmail(email string) (*w.CustomerRecord,error)

	ProductCreate(data *w.ProductRecord) error
	ProductUpdate(data *w.ProductRecord, productId int64) error
	ProductDelete(productId int64) error
	ProductFindByProjectId(projectId string) error

	CouponCreate(data any) error
	CouponUpdate(data any) error
	CouponDelete(data any) error
	CouponFindByProjectId(projectId string) error

	WebhookCreate(data w.WebhookRecord) error
	WebhookUpdate(data w.WebhookRecord) (*w.WebhookRecord, error)
	WebhookDelete(id primitive.ObjectID) error
	WebhookFindByProjectId(projectId string) ([]w.WebhookRecord, error)
}

type WoocommerceWebhookService interface {
	// WoocommerceCreateOrderWebHook create new order web hook for woocommerce
	WoocommerceCreateAllWebHook(customerKey string, customerSecret string, domainUrl string, projectId uuid.UUID) error
}


type CustomerService interface {
	
	ExtractCustomerFromOrderAndUpsert(ctx echo.Context, req *woo.OrderRecord) error 
}



type ProductService interface {
	
	ExtractProductFromOrderAndUpsert(ctx echo.Context, req *woo.OrderRecord) error 
	GetAllProductFromWoocommerce(ccustomerKey string, customerSecret string, domainUrl string, projectId uuid.UUID) error 
}