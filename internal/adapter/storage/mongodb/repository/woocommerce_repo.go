package repository

import (
	"context"
	"log"
	"log/slog"
	"time"

	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
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

	filter := bson.M{"orderId": orderID, "is_active": true, "projectId": order.ProjectID}
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
func (repo WoocommerceRepository) OrderDelete(orderID int64, projectID string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	filter := bson.M{"orderId": orderID, "projectId": projectID}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now().UTC()}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

// OrderFindByProjectID find all orders by projectID
func (repo WoocommerceRepository) OrderFindByProjectID(projectID string, size, page int, orderStatus w.OrderStatus, sort, direction string) ([]*w.OrderRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"projectId": projectID, "is_active": true, "status": orderStatus}
	if orderStatus == w.OrderStatusAll {
		filter = bson.M{"projectId": projectID, "is_active": true}
	}

	sortOrder := 1
	if direction == "desc" {
		sortOrder = -1
	} else if direction == "" {
		sortOrder = -1
	}

	// Set sort field
	sortField := sort
	if sort == "total_amount" {
		sortField = "order_total_amount_float"
	} else if sort == "" {
		sortField = "timestamp"
	}

	// Create an aggregation pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$addFields", Value: bson.M{"order_total_amount_float": bson.M{"$toDouble": "$order.total"}}}},
		{{Key: "$sort", Value: bson.D{{Key: sortField, Value: sortOrder}}}},
		{{Key: "$skip", Value: int64(size * (page - 1))}},
		{{Key: "$limit", Value: int64(size)}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx)

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

// OrderFindByProjectIDWithTimePedio find all orders by projectID
func (repo WoocommerceRepository) OrderFindByProjectIDWithTimePedio(projectID string, size, page int, orderStatus w.OrderStatus, sort, direction string, timeperiod time.Time) ([]*w.OrderRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"projectId": projectID, "is_active": true, "status": orderStatus, "timestamp": bson.M{"$gte": timeperiod}}
	if orderStatus == w.OrderStatusAll {
		filter = bson.M{"projectId": projectID, "is_active": true, "timestamp": bson.M{"$gte": timeperiod}}
	}

	sortOrder := 1
	if direction == "desc" {
		sortOrder = -1
	} else if direction == "" {
		sortOrder = -1
	}

	// Set sort field
	sortField := sort
	if sort == "total_amount" {
		sortField = "order_total_amount_float"
	} else if sort == "" {
		sortField = "timestamp"
	}

	// Create an aggregation pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$addFields", Value: bson.M{"order_total_amount_float": bson.M{"$toDouble": "$order.total"}}}},
		{{Key: "$sort", Value: bson.D{{Key: sortField, Value: sortOrder}}}},
		{{Key: "$skip", Value: int64(size * (page - 1))}},
		{{Key: "$limit", Value: int64(size)}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx)

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
func (repo WoocommerceRepository) GetOrderCount(projectID string, orderStatus w.OrderStatus, timeRange string) (int64, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Calculate the time range based on the current time
	var startTime time.Time
	now := time.Now().UTC()

	switch timeRange {
	case "24h":
		startTime = now.Add(-24 * time.Hour)
	case "7d":
		startTime = now.Add(-7 * 24 * time.Hour)
	case "1m":
		startTime = now.AddDate(0, -1, 0)
	default:
		startTime = time.Time{} // Default to the epoch time for no filtering
	}

	filter := bson.M{"projectId": projectID, "is_active": true, "status": orderStatus, "timestamp": bson.M{"$gte": startTime}}
	if orderStatus == w.OrderStatusAll {
		filter = bson.M{"projectId": projectID, "is_active": true, "timestamp": bson.M{"$gte": startTime}}
	}

	res, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
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

// GetOrderByID get order by id
func (repo WoocommerceRepository) GetOrderByID(projectID string, orderID int64) (*w.OrderRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_orders")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true, "orderId": orderID}
	var order w.OrderRecord
	err := coll.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}

// Customer

// CustomerCreate create customer
func (repo WoocommerceRepository) CustomerCreate(data *w.CustomerRecord, email string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	filter := bson.M{"email": email, "is_active": true, "projectId": data.ProjectID}
	update := bson.M{"$set": data}
	// Set upsert option to true
	opt := options.Update().SetUpsert(true)
	_, err := coll.UpdateOne(context.TODO(), filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

// CustomerUpdate update customer
func (repo WoocommerceRepository) CustomerUpdate(data *w.CustomerRecord, email string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	filter := bson.M{"email": email, "is_active": true, "projectId": data.ProjectID}
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
func (repo WoocommerceRepository) CustomerDelete(customerID int64, projectID string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	filter := bson.M{"customerId": customerID, "projectId": projectID}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now().UTC()}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

// CustomerFindByProjectID find all customers by projectID
func (repo WoocommerceRepository) CustomerFindByProjectID(projectID string, size, page int, sort, direction string) ([]*w.CustomerRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	sortOrder := 1
	if direction == "desc" {
		sortOrder = -1
	} else if direction == "" {
		sortOrder = -1
	}
	pipeline := mongo.Pipeline{
		// Match active customers by projectID
		{{Key: "$match", Value: bson.D{{Key: "projectId", Value: projectID}, {Key: "is_active", Value: true}}}},
	}

	// Conditionally add order count and sort by it
	if sort == "order_count" {
		pipeline = append(pipeline, bson.D{{Key: "$addFields", Value: bson.D{{Key: "order_count",
			Value: bson.D{{Key: "$size", Value: bson.D{{Key: "$ifNull", Value: bson.A{"$orders", bson.A{}}}}}}}}}})
		pipeline = append(pipeline, bson.D{{Key: "$sort", Value: bson.D{{Key: "order_count", Value: sortOrder}}}})
	} else {
		pipeline = append(pipeline, bson.D{{Key: "$sort",
			Value: bson.D{{Key: sort, Value: sortOrder}}}})
	}

	// Pagination: skip and limit
	pipeline = append(pipeline, bson.D{{Key: "$skip", Value: int64(size * (page - 1))}})
	pipeline = append(pipeline, bson.D{{Key: "$limit", Value: int64(size)}})

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var customers []*w.CustomerRecord
	for cursor.Next(context.TODO()) {
		var customer w.CustomerRecord
		if err := cursor.Decode(&customer); err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

// CustomerFindByEmail find customer by email
func (repo WoocommerceRepository) CustomerFindByEmail(projectID string, email string) (*w.CustomerRecord, error) {
	var result *w.CustomerRecord
	coll := repo.mongo.Database("otoo").Collection("woocommerce_customers")
	filter := bson.M{"email": email, "is_active": true, "projectId": projectID}
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
func (repo WoocommerceRepository) ProductUpdate(data *w.ProductRecord, productID int64, projectID string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	filter := bson.M{"productId": productID, "is_active": true, "projectId": projectID}
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
func (repo WoocommerceRepository) ProductDelete(productID int64, projectID string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	filter := bson.M{"productId": productID, "projectId": projectID}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now().UTC()}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

// ProductFindByProjectID find all products by projectID
func (repo WoocommerceRepository) ProductFindByProjectID(projectID string, size, page int, sort, direction string, productType w.ProductType) ([]*w.ProductRecord, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	// OrderFindByProjectID find all orders by projectID
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"projectId": projectID, "is_active": true, "product.type": bson.M{"$ne": productType.String()}}

	findOptions := options.Find()
	findOptions.SetLimit(int64(size))
	findOptions.SetSkip(int64(size * (page - 1)))
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	cursor, err := coll.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*w.ProductRecord
	for cursor.Next(context.TODO()) {
		var product w.ProductRecord
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
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
func (repo WoocommerceRepository) GetProductCount(projectID string, productType w.ProductType) (int64, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_products")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true, "product.type": bson.M{"$ne": productType.String()}}
	res, err := coll.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}
	return res, nil
}

// ProductBestSellerAggregate get best seller products
func (repo WoocommerceRepository) ProductBestSellerAggregate(projectID string) ([]bson.M, error) {
	collection := repo.mongo.Database("otoo").Collection("woocommerce_products")
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
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
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$_id"},
			{Key: "orderCount", Value: bson.D{{Key: "$first", Value: "$orderCount"}}},
			{Key: "product", Value: bson.D{{Key: "$first", Value: "$product"}}},
		}}},
		// Add fields instead of project
		{{Key: "$addFields", Value: bson.D{
			{Key: "productId", Value: "$_id"},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "productId", Value: 1},
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

// WebhookBatchDelete deactivates all webhooks for a given projectId
func (repo *WoocommerceRepository) WebhookBatchDelete(projectID string) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")

	// Create a filter to match all webhooks with the specified projectId
	filter := bson.M{"projectId": projectID}

	// Prepare the update to set is_active to false and add a deleted_at timestamp
	update := bson.M{
		"$set": bson.M{
			"is_active":  false,
			"deleted_at": time.Now().UTC(),
		},
	}

	// Update many documents that match the filter
	_, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}
	return err
}

// WebhookDelete delete webhook
func (repo *WoocommerceRepository) WebhookDelete(projectID string, webhookID int64) error {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")

	// Create a filter to match all webhooks with the specified projectId
	filter := bson.M{"projectId": projectID, "webhookId": webhookID}

	// Prepare the update to set is_active to false and add a deleted_at timestamp
	update := bson.M{
		"$set": bson.M{
			"is_active":  false,
			"deleted_at": time.Now().UTC(),
		},
	}

	// Update many documents that match the filter
	_, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

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

// WebhookCount get number of Webhook
func (repo WoocommerceRepository) WebhookCount(projectID string) (int64, error) {
	coll := repo.mongo.Database("otoo").Collection("woocommerce_webhooks")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true}
	res, err := coll.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}
	return res, nil
}
