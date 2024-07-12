package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type WoocommerceRepository struct {
	mongo *mongo.Client
}

func NewWoocommerceRepository(mongo *mongo.Client) *WoocommerceRepository {
	return &WoocommerceRepository{
		mongo,
	}
}

// Order
func (repo WoocommerceRepository) OrderCreate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	coll.InsertOne(context.TODO(), data)
	return nil
}

func (repo WoocommerceRepository) OrderUpdate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	coll.UpdateByID(context.TODO(), 1, data)
	return nil
}
func (repo WoocommerceRepository) OrderDelete(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	coll.DeleteOne(context.TODO(), data)
	return nil
}
func (repo WoocommerceRepository) OrderFindByProjectId(projectId string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	coll.FindOne(context.TODO(), projectId)
	return nil
}

// Customer
func (repo WoocommerceRepository) CustomerCreate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	coll.InsertOne(context.TODO(), data)
	return nil
}

func (repo WoocommerceRepository) CustomerUpdate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	coll.UpdateByID(context.TODO(), 1, data)
	return nil
}
func (repo WoocommerceRepository) CustomerDelete(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	coll.DeleteOne(context.TODO(), data)
	return nil
}

func (repo WoocommerceRepository) CustomerFindByProjectId(projectId string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	coll.FindOne(context.TODO(), projectId)
	return nil
}

// Coupon
func (repo WoocommerceRepository) CouponCreate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_coupons")
	coll.InsertOne(context.TODO(), data)
	return nil
}

func (repo WoocommerceRepository) CouponUpdate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_coupons")
	coll.UpdateByID(context.TODO(), 1, data)
	return nil
}
func (repo WoocommerceRepository) CouponDelete(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_coupons")
	coll.DeleteOne(context.TODO(), data)
	return nil
}
func (repo WoocommerceRepository) CouponFindByProjectId(projectId string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_coupons")
	coll.FindOne(context.TODO(), projectId)
	return nil
}

// Product
func (repo WoocommerceRepository) ProductCreate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	coll.InsertOne(context.TODO(), data)
	return nil
}

func (repo WoocommerceRepository) ProductUpdate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	coll.UpdateByID(context.TODO(), 1, data)
	return nil
}
func (repo WoocommerceRepository) ProductDelete(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	coll.DeleteOne(context.TODO(), data)
	return nil
}
func (repo WoocommerceRepository) ProductFindByProjectId(projectId string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	coll.FindOne(context.TODO(), projectId)
	return nil
}
