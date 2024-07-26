package repository

import (
	"context"
	"time"

	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AnalyticsRepository represents the repository for analytics data
type AnalyticsRepository struct {
	mongo *mongo.Client
}

// NewAnalyticsRepository creates a new instance of AnalyticsRepository
func NewAnalyticsRepository(mongo *mongo.Client) *AnalyticsRepository {
	return &AnalyticsRepository{
		mongo,
	}
}

// FindBestSellers returns the best sellers for a given project
func (r *AnalyticsRepository) FindBestSellers(projectID string, size, page int) ([]*domain.ProductBestSellerRecord, error) {
	coll := r.mongo.Database("otoo").Collection("woocommerce_product_best_seller")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"projectId": projectID, "is_active": true}
	findOptions := options.Find()
	findOptions.SetLimit(int64(size))
	findOptions.SetSkip(int64(size * (page - 1)))
	findOptions.SetSort(bson.D{{Key: "timestamp", Value: -1}})
	cursor, err := coll.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*w.ProductBestSellerRecord
	for cursor.Next(ctx) {
		var productBestSeller w.ProductBestSellerRecord
		if err := cursor.Decode(&productBestSeller); err != nil {
			return nil, err
		}
		products = append(products, &productBestSeller)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// CreateBestSellers creates the best sellers for a given project
func (r *AnalyticsRepository) CreateBestSellers(projectID string, data []*domain.ProductBestSellerRecord) error {
	coll := r.mongo.Database("otoo").Collection("woocommerce_product_best_seller")
	coll.InsertMany(context.TODO(), convertToInterfaceSlice(data))
	return nil
}

// convertToInterfaceSlice converts a slice of ProductBestSellerRecord to a slice of interface{}
func convertToInterfaceSlice(data []*domain.ProductBestSellerRecord) []interface{} {
	var result []interface{}
	for _, item := range data {
		result = append(result, item)
	}
	return result
}
