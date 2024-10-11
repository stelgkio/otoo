package repository

import (
	"context"
	"fmt"
	"log/slog"
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
	findOptions.SetSort(bson.D{{Key: "total_orders", Value: -1}})
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

// DeleteBestSellers creates the best sellers for a given project
func (r *AnalyticsRepository) DeleteBestSellers(projectID string) error {
	coll := r.mongo.Database("otoo").Collection("woocommerce_product_best_seller")
	filter := bson.M{"projectId": projectID, "is_active": true}
	coll.DeleteMany(context.TODO(), filter)
	return nil
}

// FindWeeklyBalance returns the weekly balance for a given project with pagination (size and page).
func (r *AnalyticsRepository) FindWeeklyBalance(projectID string, size, page int) ([]*domain.WeeklyAnalytics, error) {
	// Ensure size and page have valid values
	if size <= 0 {
		size = 10 // Default page size
	}
	if page <= 0 {
		page = 1 // Default to the first page
	}

	coll := r.mongo.Database("otoo").Collection("weekly_balance_analytics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter: Find documents where the projectID matches and is_active is true
	filter := bson.M{"projectId": projectID}

	// Define the pagination options: Sort by the Timestamp field in descending order
	opts := options.Find().
		SetSort(bson.D{{Key: "timestamp", Value: -1}}).
		SetLimit(int64(size)).            // Limit results to the page size
		SetSkip(int64((page - 1) * size)) // Skip records for pagination

	// Define a slice to hold the result
	var results []*domain.WeeklyAnalytics

	// Query the database and find the WeeklyAnalytics documents for the given project
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weekly analytics for projectID: %s, error: %v", projectID, err)
	}
	defer cursor.Close(ctx)

	// Decode each document and append to the results slice
	for cursor.Next(ctx) {
		var weeklyAnalytics domain.WeeklyAnalytics
		if err := cursor.Decode(&weeklyAnalytics); err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		results = append(results, &weeklyAnalytics)
	}

	// Check if there was any error during the cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor iteration error: %v", err)
	}

	// Return the results slice
	return results, nil
}

// FindLatestWeeklyBalance returns the weekly balance for a given project
func (r *AnalyticsRepository) FindLatestWeeklyBalance(projectID string) (*domain.WeeklyAnalytics, error) {
	coll := r.mongo.Database("otoo").Collection("weekly_balance_analytics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter: Find documents where the projectID matches
	filter := bson.M{"projectId": projectID}

	// Define options: Sort by the Timestamp field in descending order and limit the result to 1
	opts := options.FindOne().SetSort(bson.D{{Key: "timestamp", Value: -1}})

	// Define a variable to hold the result
	var result domain.WeeklyAnalytics

	// Query the database and find the latest WeeklyAnalytics for the given project
	err := coll.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle case when no document is found
			slog.Error("no weekly analytics found for projectID: %s", "error", projectID)
			return nil, nil
		}
		// Handle other potential errors
		return nil, fmt.Errorf("failed to fetch weekly analytics for projectID: %s, error: %v", projectID, err)
	}

	// Return the found document
	return &result, nil

}

// CreateWeeklyBalance creates a new weekly balance for a given project
func (r *AnalyticsRepository) CreateWeeklyBalance(projectID string, data *domain.WeeklyAnalytics) error {
	coll := r.mongo.Database("otoo").Collection("weekly_balance_analytics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set the project ID and timestamp for the new record
	data.ProjectID = projectID
	data.Timestamp = time.Now()

	// Insert the new WeeklyAnalytics record into the database
	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create weekly analytics for projectID: %s, error: %v", projectID, err)
	}

	return nil
}

// DeleteWeeklyBalance deletes the weekly balance for a given project
func (r *AnalyticsRepository) DeleteWeeklyBalance(projectID string) error {
	coll := r.mongo.Database("otoo").Collection("weekly_balance_analytics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter: Find documents where the projectID matches
	filter := bson.M{"projectId": projectID}

	// Delete the document(s) matching the projectID
	res, err := coll.DeleteMany(ctx, filter) // Using DeleteMany, you can switch to DeleteOne if needed
	if err != nil {
		return fmt.Errorf("failed to delete weekly analytics for projectID: %s, error: %v", projectID, err)
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("no records deleted for projectID: %s", projectID)
	}

	return nil
}

// FindMonthlyCount returns the monthly order count for a given project with pagination (size and page).
func (r *AnalyticsRepository) FindMonthlyCount(projectID string, size, page int) ([]*domain.MonthlyOrderCountAnalytics, error) {
	// Ensure size and page have valid values
	if size <= 0 {
		size = 10 // Default page size
	}
	if page <= 0 {
		page = 1 // Default to the first page
	}

	coll := r.mongo.Database("otoo").Collection("monthly_order_count_analytics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter: Find documents where the projectID matches
	filter := bson.M{"projectId": projectID}

	// Define the pagination options: Sort by the Timestamp field in descending order
	opts := options.Find().
		SetSort(bson.D{{Key: "timestamp", Value: -1}}).
		SetLimit(int64(size)).            // Limit results to the page size
		SetSkip(int64((page - 1) * size)) // Skip records for pagination

	// Define a slice to hold the result
	var results []*domain.MonthlyOrderCountAnalytics

	// Query the database and find the MonthlyOrderCountAnalytics documents for the given project
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch monthly order counts for projectID: %s, error: %v", projectID, err)
	}
	defer cursor.Close(ctx)

	// Decode each document and append to the results slice
	for cursor.Next(ctx) {
		var monthlyCount domain.MonthlyOrderCountAnalytics
		if err := cursor.Decode(&monthlyCount); err != nil {
			return nil, fmt.Errorf("error decoding document: %v", err)
		}
		results = append(results, &monthlyCount)
	}

	// Check if there was any error during the cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor iteration error: %v", err)
	}

	// Return the results slice
	return results, nil
}

// FindLatestMonthlyCount returns the latest monthly order count for a given project.
func (r *AnalyticsRepository) FindLatestMonthlyCount(projectID string) (*domain.MonthlyOrderCountAnalytics, error) {
	coll := r.mongo.Database("otoo").Collection("monthly_order_count_analytics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter: Find documents where the projectID matches
	filter := bson.M{"projectId": projectID}

	// Define options: Sort by the Timestamp field in descending order and limit the result to 1
	opts := options.FindOne().SetSort(bson.D{{Key: "timestamp", Value: -1}})

	// Define a variable to hold the result
	var result domain.MonthlyOrderCountAnalytics

	// Query the database and find the latest MonthlyOrderCountAnalytics for the given project
	err := coll.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle case when no document is found
			return nil, nil
		}
		// Handle other potential errors
		return nil, fmt.Errorf("failed to fetch latest monthly order count for projectID: %s, error: %v", projectID, err)
	}

	// Return the found document
	return &result, nil
}

// CreateMonthlyCount creates a new monthly order count for a given project.
func (r *AnalyticsRepository) CreateMonthlyCount(projectID string, data *domain.MonthlyOrderCountAnalytics) error {
	coll := r.mongo.Database("otoo").Collection("monthly_order_count_analytics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set the project ID and timestamp for the new record
	data.ProjectID = projectID
	data.Timestamp = time.Now()

	// Insert the new MonthlyOrderCountAnalytics record into the database
	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create monthly order count for projectID: %s, error: %v", projectID, err)
	}

	return nil
}

// DeleteMonthlyCount deletes the monthly order counts for a given project.
func (r *AnalyticsRepository) DeleteMonthlyCount(projectID string) error {
	coll := r.mongo.Database("otoo").Collection("monthly_order_count_analytics")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the filter: Find documents where the projectID matches
	filter := bson.M{"projectId": projectID}

	// Delete the document(s) matching the projectID
	res, err := coll.DeleteMany(ctx, filter) // Using DeleteMany, you can switch to DeleteOne if needed
	if err != nil {
		return fmt.Errorf("failed to delete monthly order counts for projectID: %s, error: %v", projectID, err)
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("no records deleted for projectID: %s", projectID)
	}

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
