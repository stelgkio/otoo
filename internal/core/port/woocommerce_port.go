package port

import (
	"github.com/google/uuid"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WoocommerceRepository interface {
	// InsertWoocommerceOrder inserts a new order into the database
	OrderCreate(data any) error
	OrderUpdate(data any) error
	OrderDelete(data any) error
	OrderFindByProjectId(projectId string) error

	CustomerCreate(data any) error
	CustomerUpdate(data any) error
	CustomerDelete(data any) error
	CustomerFindByProjectId(projectId string) error

	ProductCreate(data any) error
	ProductUpdate(data any) error
	ProductDelete(data any) error
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
