package repository

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
	domain "github.com/stelgkio/otoo/internal/core/domain/courier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// VoucherRepository represents the repository for analytics data
type VoucherRepository struct {
	mongo *mongo.Client
}

// NewVoucherRepository creates a new instance of AnalyticsRepository
func NewVoucherRepository(mongo *mongo.Client) *VoucherRepository {
	return &VoucherRepository{
		mongo,
	}
}

// CreateVoucher inserts a new Voucher into the database.
func (r *VoucherRepository) CreateVoucher(ctx echo.Context, voucher *domain.Voucher, projectID string) (*domain.Voucher, error) {
	collection := r.mongo.Database("otoo").Collection("vouchers")

	// Define the filter to find an existing voucher
	filter := bson.M{"projectId": projectID, "is_active": true, "orderId": voucher.OrderID, "is_printed": false} // Ensure `OrderID` is a field in your voucher

	// Define the update document
	update := bson.M{
		"$set": voucher, // Set the voucher fields to the new values
	}

	// Upsert the voucher into the collection
	_, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
		options.Update().SetUpsert(true), // Enable upsert
	)

	if err != nil {
		slog.Error("Failed to create or update voucher", "error", err)
		return nil, err
	}

	return voucher, nil
}

// UpdateVoucher updates a Voucher by voucherID and returns the updated Voucher.
func (r *VoucherRepository) UpdateVoucher(ctx echo.Context, voucher *domain.Voucher, projectID string, voucherID string, orderID int64) (*domain.Voucher, error) {
	collection := r.mongo.Database("otoo").Collection("vouchers")

	// Prepare the filter for finding the voucher
	filter := bson.M{"voucher_id": voucherID, "projectId": projectID, "is_active": true, "orderId": orderID}
	if voucherID == "" {
		filter = bson.M{"projectId": projectID, "is_active": true, "orderId": orderID}
	}

	// Prepare the update data
	update := bson.M{"$set": voucher}

	// Set upsert option to false if you don't want to create a new document if it doesn't exist
	opt := options.Update().SetUpsert(true)

	// Perform the update operation
	result, err := collection.UpdateOne(context.Background(), filter, update, opt)
	if err != nil {
		return nil, err
	}

	// Check if any documents were modified
	if result.MatchedCount == 0 {
		return nil, errors.New("voucher not found")
	}

	// Retrieve the updated voucher
	var updatedVoucher domain.Voucher
	err = collection.FindOne(context.Background(), filter).Decode(&updatedVoucher)
	if err != nil {
		return nil, err
	}

	return &updatedVoucher, nil
}

// GetVoucherByVoucherID selects a Voucher by voucher ID.
func (r *VoucherRepository) GetVoucherByVoucherID(ctx echo.Context, voucherID string) (*domain.Voucher, error) {
	collection := r.mongo.Database("otoo").Collection("vouchers")
	var voucher domain.Voucher
	id, err := primitive.ObjectIDFromHex(voucherID)
	if err != nil {
		return nil, err
	}

	// Find the voucher in the collection by voucherID
	err = collection.FindOne(ctx.Request().Context(), bson.M{"_id": id}).Decode(&voucher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("voucher not found")
		}
		return nil, err
	}

	return &voucher, nil
}

// GetAllVouchers retrieves all vouchers for a specific project.
func (r *VoucherRepository) GetAllVouchers(ctx echo.Context, projectID string) ([]*domain.Voucher, error) {
	collection := r.mongo.Database("otoo").Collection("vouchers")

	var vouchers []*domain.Voucher

	// Find all vouchers for the given project ID
	cursor, err := collection.Find(context.Background(), bson.M{"projectId": projectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode the vouchers
	for cursor.Next(context.Background()) {
		var voucher domain.Voucher
		if err := cursor.Decode(&voucher); err != nil {
			return nil, err
		}
		vouchers = append(vouchers, &voucher)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return vouchers, nil
}

// FindVoucherByProjectID finds all vouchers by projectID with pagination and sorting.
func (r *VoucherRepository) FindVoucherByProjectID(projectID string, size, page int, sort, direction string, voucherStatus domain.VoucherStatus) ([]*domain.Voucher, error) {
	collection := r.mongo.Database("otoo").Collection("vouchers")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Prepare filter to find vouchers for the given projectID
	filter := bson.M{"projectId": projectID, "is_active": true, "status": voucherStatus}
	if voucherStatus == domain.VoucherStatusAll {
		filter = bson.M{"projectId": projectID, "is_active": true}
	}
	// Determine sort order
	sortOrder := 1
	if direction == "desc" {
		sortOrder = -1
	} else if direction == "" {
		sortOrder = -1
	}

	// Set sort field
	sortField := sort
	if sort == "" {
		sortField = "created_at" // Default sort field if none provided
	}

	// Create an aggregation pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sort", Value: bson.D{{Key: sortField, Value: sortOrder}}}},
		{{Key: "$skip", Value: int64(size * (page - 1))}}, // Skip for pagination
		{{Key: "$limit", Value: int64(size)}},             // Limit the number of documents returned
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No documents found
		}
		return nil, err // Return any other error
	}
	defer cursor.Close(ctx)

	var vouchers []*domain.Voucher
	for cursor.Next(ctx) {
		var voucher domain.Voucher
		if err := cursor.Decode(&voucher); err != nil {
			return nil, err // Return decoding error
		}
		vouchers = append(vouchers, &voucher) // Append to results
	}

	if err := cursor.Err(); err != nil {
		return nil, err // Return cursor error
	}

	return vouchers, nil // Return the list of vouchers
}

// DeleteVouchersByID marks a Voucher as inactive by setting is_active to false and adding a deleted_at timestamp.
func (r *VoucherRepository) DeleteVouchersByID(ctx echo.Context, voucherID string) error {
	collection := r.mongo.Database("otoo").Collection("vouchers")

	// Prepare the filter to find the voucher by voucherID
	filter := bson.M{"voucher_id": voucherID}

	// Prepare the update to mark the voucher as inactive and set the deletion timestamp
	update := bson.M{
		"$set": bson.M{
			"is_active":  false,
			"deleted_at": time.Now().UTC(),
		},
	}

	// Perform the update operation
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// GetVoucherCount gets the number of vouchers by projectID and status.
func (r *VoucherRepository) GetVoucherCount(projectID string, voucherStatus domain.VoucherStatus) (int64, error) {
	coll := r.mongo.Database("otoo").Collection("vouchers")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Prepare filter to count vouchers for the given projectID and status
	filter := bson.M{"projectId": projectID, "is_active": true}

	// If a specific status is provided, include it in the filter
	if voucherStatus != domain.VoucherStatusAll {
		filter["status"] = voucherStatus
	}

	// Count the documents that match the filter
	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil // No documents found
		}
		return 0, err // Return any other error
	}
	return count, nil // Return the count of vouchers
}

// GetVoucherByOrderIDAndProjectID retrieves a voucher by orderID and projectID.
func (r *VoucherRepository) GetVoucherByOrderIDAndProjectID(ctx echo.Context, orderID int64, projectID string) (*domain.Voucher, error) {
	collection := r.mongo.Database("otoo").Collection("vouchers")

	// Use the echo context to get the request context and set a timeout
	reqCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Prepare filter to find vouchers for the given projectID and active status
	filter := bson.M{"orderId": orderID, "projectId": projectID, "is_active": true}

	// Query the database
	cursor, err := collection.Find(reqCtx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No documents found
		}
		return nil, err // Return any other error
	}
	defer cursor.Close(reqCtx) // Ensure the cursor is closed when done

	// Initialize a variable to store the result
	var voucher domain.Voucher

	// Loop through the cursor and decode the first matching document
	for cursor.Next(reqCtx) {
		if err := cursor.Decode(&voucher); err != nil {
			return nil, err // Return decoding error
		}
		// Return the first matching voucher
		return &voucher, nil
	}

	// If no matching voucher was found, return nil
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return nil, nil // Return nil if no voucher was found
}
