package repository

import (
	"context"
	"time"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (repo WoocommerceRepository) OrderCreate(data *w.OrderRecord) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	coll.InsertOne(context.TODO(), data)
	return nil
}
func (repo WoocommerceRepository) OrderUpdate(order *w.OrderRecord, orderId int64) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")

	filter := bson.M{ "orderId": orderId, "is_active": true }
	update := bson.M{"$set": order}

    // Set upsert option to true
    opt := options.Update().SetUpsert(true)

    // Perform the upsert operation
    _, err := coll.UpdateOne(context.TODO(), filter, update, opt)
    if err != nil {
        return err
    }
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
func (repo WoocommerceRepository) CustomerCreate(data *w.CustomerRecord) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	coll.InsertOne(context.TODO(), data)
	return nil
}
func (repo WoocommerceRepository) CustomerUpdate(data *w.CustomerRecord, email string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	filter := bson.M{ "email": email, "is_active": true }
	update := bson.M{"$set": data}

    // Set upsert option to true
    opt := options.Update().SetUpsert(true)

    // Perform the upsert operation
    _, err := coll.UpdateOne(context.TODO(), filter, update, opt)
    if err != nil {
        return err
    }
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
func (repo WoocommerceRepository) CustomerFindByEmail(email string) (*w.CustomerRecord,error) {
	var result *w.CustomerRecord
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	filter := bson.M{ "email": email, "is_active": true }
	err := coll.FindOne(context.TODO(),filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return result , nil
}
// Product
func (repo WoocommerceRepository) ProductCreate(data *w.ProductRecord) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	coll.InsertOne(context.TODO(), data)
	return nil
}
func (repo WoocommerceRepository) ProductUpdate(data *w.ProductRecord, productId int64) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	filter := bson.M{ "productId": productId, "is_active": true }
	update := bson.M{"$set": data}

    // Set upsert option to true
    opt := options.Update().SetUpsert(true)

    // Perform the upsert operation
    _, err := coll.UpdateOne(context.TODO(), filter, update, opt)
    if err != nil {
        return err
    }
	return nil
}
func (repo WoocommerceRepository) ProductDelete(productId int64) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	filter := bson.M{"productId": productId}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now()}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}
func (repo WoocommerceRepository) ProductFindByProjectId(projectId string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
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





// woocommerce
func (repo *WoocommerceRepository) WebhookCreate(data w.WebhookRecord) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")
	_, err := coll.InsertOne(context.TODO(), data)
	return err
}

func (repo *WoocommerceRepository) WebhookUpdate(data w.WebhookRecord) (*w.WebhookRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhookss")
	filter := bson.M{"projectId": data.ProjectID}
	update := bson.M{"$set": data}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedRecord w.WebhookRecord
	err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedRecord)
	if err != nil {
		return nil, err
	}
	return &updatedRecord, nil
}

func (repo *WoocommerceRepository) WebhookDelete(id primitive.ObjectID) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now()}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

func (repo *WoocommerceRepository) WebhookFindByProjectId(projectId string) ([]w.WebhookRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")
	filter := bson.M{"projectId": projectId, "is_active": true}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []w.WebhookRecord
	for cursor.Next(context.TODO()) {
		var record w.WebhookRecord
		if err := cursor.Decode(&record); err != nil {
			return nil, err
		}
		results = append(results, record)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
