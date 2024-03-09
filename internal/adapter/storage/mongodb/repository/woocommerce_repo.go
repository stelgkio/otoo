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

func (repo WoocommerceRepository) InsertWoocommerceOrder(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	coll.InsertOne(context.TODO(), data)
	return nil
}
