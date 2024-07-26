package repository

import (
	"context"
	"log"
	"log/slog"
	"time"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// WoocommerceRepository represents the repository for woocommerce data
type WoocommerceRepository struct {
	mongo *mongo.Client
}

// NewWoocommerceRepository creates a new instance of WoocommerceRepository
func NewWoocommerceRepository(mongo *mongo.Client) *WoocommerceRepository {
	return &WoocommerceRepository{
		mongo,
	}
}

// Order

// OrderCreate create order
func (repo WoocommerceRepository) OrderCreate(data *w.OrderRecord) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	coll.InsertOne(context.TODO(), data)
	return nil
}

// OrderUpdate update order
func (repo WoocommerceRepository) OrderUpdate(order *w.OrderRecord, orderID int64) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")

	filter := bson.M{"orderId": orderID, "is_active": true}
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

// OrderDelete delete order
func (repo WoocommerceRepository) OrderDelete(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	coll.DeleteOne(context.TODO(), data)
	return nil
}

// OrderFindByProjectID find all orders by projectID
func (repo WoocommerceRepository) OrderFindByProjectID(projectID string, size, page int) ([]*w.OrderRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	filter := bson.M{"projectId": projectID, "is_active": true}
	findOptions := options.Find()
	findOptions.SetLimit(int64(size))
	findOptions.SetSkip(int64(size * (page - 1)))
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	cursor, err := coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var orders []*w.OrderRecord
	for cursor.Next(context.TODO()) {
		var order w.OrderRecord
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrderCount get number of orders
func (repo WoocommerceRepository) GetOrderCount(projectID string, orderStatus w.OrderStatus) (int64, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true, "status": orderStatus}
	res, err := coll.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}
	return res, nil
}

// GetOrdersCountBetweenOrEquals get number of orders between or equals to timeperiod
func (repo WoocommerceRepository) GetOrdersCountBetweenOrEquals(projectID string, timeperiod time.Time, orderStatus w.OrderStatus) (int64, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true, "status": orderStatus, "timestamp": bson.M{"$gte": timeperiod}}
	totalcount, err := coll.CountDocuments(ctx, filter)
	return totalcount, err
}

// Customer

// CustomerCreate create customer
func (repo WoocommerceRepository) CustomerCreate(data *w.CustomerRecord) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	coll.InsertOne(context.TODO(), data)
	return nil
}

// CustomerUpdate update customer
func (repo WoocommerceRepository) CustomerUpdate(data *w.CustomerRecord, email string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	filter := bson.M{"email": email, "is_active": true}
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

// CustomerDelete  error
func (repo WoocommerceRepository) CustomerDelete(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	coll.DeleteOne(context.TODO(), data)
	return nil
}

// CustomerFindByProjectID find all customers by projectID
func (repo WoocommerceRepository) CustomerFindByProjectID(projectID string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	coll.FindOne(context.TODO(), projectID)
	return nil
}

// CustomerFindByEmail find customer by email
func (repo WoocommerceRepository) CustomerFindByEmail(email string) (*w.CustomerRecord, error) {
	var result *w.CustomerRecord
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	filter := bson.M{"email": email, "is_active": true}
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// GetCustomerCount get number of customers
func (repo WoocommerceRepository) GetCustomerCount(projectID string) (int64, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true}
	res, err := coll.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}
	return res, nil
}

// Product

// ProductCreate create product
func (repo WoocommerceRepository) ProductCreate(data *w.ProductRecord) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	coll.InsertOne(context.TODO(), data)
	return nil
}

// ProductUpdate update product
func (repo WoocommerceRepository) ProductUpdate(data *w.ProductRecord, productID int64) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	filter := bson.M{"productId": productID, "is_active": true}
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

// ProductDelete delete product
func (repo WoocommerceRepository) ProductDelete(productID int64) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	filter := bson.M{"productId": productID}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now()}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

// ProductFindByProjectID find all products by projectID
func (repo WoocommerceRepository) ProductFindByProjectID(projectID string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	coll.FindOne(context.TODO(), projectID)
	return nil
}

// GetProductByID get product by projectID and orderID
func (repo WoocommerceRepository) GetProductByID(projectID string, productID int64) (*w.ProductRecord, error) {
	var result *w.ProductRecord
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true, "productId": productID}
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

// GetProductCount get number of products
func (repo WoocommerceRepository) GetProductCount(projectID string) (int64, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true}
	res, err := coll.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}
	return res, nil
}

// GetProductBestSeller get best seller products
func (repo WoocommerceRepository) ProductBestSellerAggregate(projectID string) ([]bson.M, error) {
	collection := repo.mongo.Database("otoo").Collection("woocommerce_products")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Define the aggregation pipeline
	// Define the aggregation pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "projectId", Value: projectID}}}},
		{{Key: "$unwind", Value: "$orders"}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$productId"},
			{Key: "orderCount", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
		{{Key: "$sort", Value: bson.D{{Key: "orderCount", Value: -1}}}},
		{{Key: "$limit", Value: 5}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "woocommerce_products"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "productId"},
			{Key: "as", Value: "product"},
		}}},
		{{Key: "$unwind", Value: "$product"}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "productId", Value: "$_id"},
			{Key: "orderCount", Value: 1},
			{Key: "product", Value: 1},
		}}},
	}

	// Execute the aggregation
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		slog.Error("collection.Aggregate", "error", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	//var products []*w.ProductRecord

	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		log.Fatal("Error decoding cursor result: ", err)
	}

	return results, nil
}

// Coupon

// CouponCreate create coupon
func (repo WoocommerceRepository) CouponCreate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_coupons")
	coll.InsertOne(context.TODO(), data)
	return nil
}

// CouponUpdate update coupon
func (repo WoocommerceRepository) CouponUpdate(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_coupons")
	coll.UpdateByID(context.TODO(), 1, data)
	return nil
}

// CouponDelete delete coupon
func (repo WoocommerceRepository) CouponDelete(data any) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_coupons")
	coll.DeleteOne(context.TODO(), data)
	return nil
}

// CouponFindByProjectID find all coupons by projectID
func (repo WoocommerceRepository) CouponFindByProjectID(projectID string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_coupons")
	coll.FindOne(context.TODO(), projectID)
	return nil
}

// woocommerce

// WebhookCreate create webhook
func (repo *WoocommerceRepository) WebhookCreate(data w.WebhookRecord) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")
	_, err := coll.InsertOne(context.TODO(), data)
	return err
}

// WebhookUpdate update webhook
func (repo *WoocommerceRepository) WebhookUpdate(data w.WebhookRecord) (*w.WebhookRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhookss")
	filter := bson.M{"projectID": data.ProjectID}
	update := bson.M{"$set": data}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedRecord w.WebhookRecord
	err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedRecord)
	if err != nil {
		return nil, err
	}
	return &updatedRecord, nil
}

// WebhookDelete delete webhook
func (repo *WoocommerceRepository) WebhookDelete(id primitive.ObjectID) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now()}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

// WebhookFindByProjectID find all webhooks by projectID
func (repo *WoocommerceRepository) WebhookFindByProjectID(projectID string) ([]w.WebhookRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")
	filter := bson.M{"projectId": projectID, "is_active": true}
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
