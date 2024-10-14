package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExtensionRepository represents the repository for Extension-related operations
type ExtensionRepository struct {
	mongo *mongo.Client
}

// NewExtensionRepository creates a new Extension repository instance
func NewExtensionRepository(mongo *mongo.Client) *ExtensionRepository {
	return &ExtensionRepository{
		mongo,
	}
}

// CreateExtension creates a new Extension
func (ex *ExtensionRepository) CreateExtension(ctx echo.Context, c *domain.Extension) error {
	panic("unimplemented")
}

// GetAllExtensions gets all Extensions
func (ex *ExtensionRepository) GetAllExtensions(ctx echo.Context) ([]*domain.Extension, error) {
	collection := ex.mongo.Database("otoo").Collection("extensions")
	cursor, err := collection.Find(ctx.Request().Context(), bson.M{"is_active": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx.Request().Context())

	var Extensions []*domain.Extension
	for cursor.Next(ctx.Request().Context()) {
		var Extension domain.Extension
		if err := cursor.Decode(&Extension); err != nil {
			return nil, err
		}
		Extensions = append(Extensions, &Extension)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return Extensions, nil
}

// GetExtensionByID gets a Extension by ID
func (ex *ExtensionRepository) GetExtensionByID(ctx echo.Context, extensionID string) (*domain.Extension, error) {
	collection := ex.mongo.Database("otoo").Collection("extensions")

	// Convert the extensionID from string to ObjectID
	id, err := primitive.ObjectIDFromHex(extensionID)
	if err != nil {
		return nil, errors.New("Invalid extension ID format")
	}

	// Create the filter to match both the extension ID and IsActive
	filter := bson.M{
		"_id":       id,
		"is_active": true,
	}

	// Define a variable to hold the result
	var Extension domain.Extension

	// Execute the query with the filter
	err = collection.FindOne(ctx.Request().Context(), filter).Decode(&Extension)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Extension not found")
		}
		return nil, err
	}

	// Return the found extension
	return &Extension, nil
}

// GetExtensionByCode gets a Extension by code
func (ex *ExtensionRepository) GetExtensionByCode(ctx echo.Context, code string) (*domain.Extension, error) {
	collection := ex.mongo.Database("otoo").Collection("extensions")

	// Create the filter to match both the extension ID and IsActive
	filter := bson.M{
		"code":      code,
		"is_active": true,
	}

	// Define a variable to hold the result
	var Extension domain.Extension

	// Execute the query with the filter
	err := collection.FindOne(ctx.Request().Context(), filter).Decode(&Extension)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Extension not found")
		}
		return nil, err
	}

	// Return the found extension
	return &Extension, nil

}

// DeleteExtension deletes a Extension by ID
func (ex *ExtensionRepository) DeleteExtension(ctx echo.Context, extensionID string) error {
	collection := ex.mongo.Database("otoo").Collection("Extensions")
	id, err := primitive.ObjectIDFromHex(extensionID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(
		ctx.Request().Context(),
		bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"updated_at": time.Now().UTC(),
				"is_active":  false,
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

//////////////////PROJECT ExtensionS/////////////////////////

// CreateProjectExtension creates a new ProjectExtension
func (ex *ExtensionRepository) CreateProjectExtension(ctx echo.Context, projectID string, e *domain.Extension) error {
	collection := ex.mongo.Database("otoo").Collection("project_extensions")

	projectExtension := &domain.ProjectExtension{
		ID:          primitive.NewObjectID(),
		Title:       e.Title,
		Description: e.Description,
		Code:        e.Code,
		ProjectID:   projectID,
		ExtensionID: e.ID.Hex(),
		CreatedAt:   time.Now().UTC(),
		IsActive:    true,
	}
	filter := bson.M{"extension_id": e.ID.Hex(), "is_active": true, "project_id": projectID}
	update := bson.M{"$set": projectExtension}

	// Set upsert option to true
	opt := options.Update().SetUpsert(true)

	// Perform the upsert operation
	_, err := collection.UpdateOne(context.TODO(), filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

// GetAllProjectExtensions gets all ProjectExtensions
func (ex *ExtensionRepository) GetAllProjectExtensions(ctx echo.Context, projectID string) ([]*domain.ProjectExtension, error) {
	collection := ex.mongo.Database("otoo").Collection("project_extensions")

	filter := bson.M{
		"project_id": projectID,
		"is_active":  true,
	}

	cursor, err := collection.Find(ctx.Request().Context(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx.Request().Context())

	var projectExtensions []*domain.ProjectExtension
	for cursor.Next(ctx.Request().Context()) {
		var projectExtension domain.ProjectExtension
		if err := cursor.Decode(&projectExtension); err != nil {
			return nil, err
		}
		projectExtensions = append(projectExtensions, &projectExtension)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return projectExtensions, nil
}

// GetProjectExtensionByID gets a ProjectExtension by ID
func (ex *ExtensionRepository) GetProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtension, error) {
	collection := ex.mongo.Database("otoo").Collection("project_extensions")

	extID, err := primitive.ObjectIDFromHex(extensionID)
	if err != nil {
		return nil, errors.New("Invalid extension ID format")
	}

	filter := bson.M{
		"_id":        extID,
		"project_id": projectID,
		"is_active":  true,
	}

	var projectExtension domain.ProjectExtension
	err = collection.FindOne(ctx.Request().Context(), filter).Decode(&projectExtension)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &projectExtension, nil
}

// DeleteProjectExtension deletes a ProjectExtension by ID
func (ex *ExtensionRepository) DeleteProjectExtension(ctx echo.Context, extensionID, projectID string) error {
	collection := ex.mongo.Database("otoo").Collection("project_extensions")

	extID, err := primitive.ObjectIDFromHex(extensionID)
	if err != nil {
		return errors.New("Invalid extension ID format")
	}

	filter := bson.M{
		"_id":        extID,
		"project_id": projectID,
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"is_active":  false,
		},
	}

	result, err := collection.UpdateOne(ctx.Request().Context(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("ProjectExtension not found")
	}

	return nil
}

//////////////////  ACS Extension  /////////////////////////

// CreateACSProjectExtension creates a new ProjectExtension
func (ex *ExtensionRepository) CreateACSProjectExtension(ctx echo.Context, projectID string, e *domain.AcsCourierExtension) error {
	collection := ex.mongo.Database("otoo").Collection("acs_project_extensions")

	filter := bson.M{"extension_id": e.ExtensionID, "is_active": true, "project_id": projectID}
	update := bson.M{"$set": e}

	// Set upsert option to true
	opt := options.Update().SetUpsert(true)

	// Perform the upsert operation
	_, err := collection.UpdateOne(context.TODO(), filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

// GetAllACSProjectExtensions gets all ProjectExtensions
func (ex *ExtensionRepository) GetAllACSProjectExtensions(ctx echo.Context, projectID string) ([]*domain.AcsCourierExtension, error) {
	collection := ex.mongo.Database("otoo").Collection("acs_project_extensions")

	filter := bson.M{
		"project_id": projectID,
		"is_active":  true,
	}

	cursor, err := collection.Find(ctx.Request().Context(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx.Request().Context())

	var projectExtensions []*domain.AcsCourierExtension
	for cursor.Next(ctx.Request().Context()) {
		var projectExtension domain.AcsCourierExtension
		if err := cursor.Decode(&projectExtension); err != nil {
			return nil, err
		}
		projectExtensions = append(projectExtensions, &projectExtension)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return projectExtensions, nil
}

// GetACSProjectExtensionByID gets a ProjectExtension by ID
func (ex *ExtensionRepository) GetACSProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.AcsCourierExtension, error) {
	collection := ex.mongo.Database("otoo").Collection("acs_project_extensions")

	filter := bson.M{
		"project_id":   projectID,
		"extension_id": extensionID,
		"is_active":    true,
	}

	var projectExtension domain.AcsCourierExtension
	err := collection.FindOne(ctx.Request().Context(), filter).Decode(&projectExtension)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &projectExtension, nil
}

// DeleteACSProjectExtension deletes a ProjectExtension by ID
func (ex *ExtensionRepository) DeleteACSProjectExtension(ctx echo.Context, extensionID, projectID string) error {
	collection := ex.mongo.Database("otoo").Collection("acs_project_extensions")

	extID, err := primitive.ObjectIDFromHex(extensionID)
	if err != nil {
		return errors.New("Invalid extension ID format")
	}

	filter := bson.M{
		"_id":          extID,
		"project_id":   projectID,
		"extension_id": extensionID,
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"is_active":  false,
		},
	}

	result, err := collection.UpdateOne(ctx.Request().Context(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("ProjectExtension not found")
	}

	return nil
}

////////////////// Data Synchronizer Extension/////////////////////////

// CreateSynchronizerProjectExtension creates a new ProjectExtension
// CreateSynchronizerProjectExtension creates a new ProjectExtension
func (ex *ExtensionRepository) CreateSynchronizerProjectExtension(ctx echo.Context, projectID string, e *domain.DataSynchronizerExtension) error {
	collection := ex.mongo.Database("otoo").Collection("synchronizer_project_extensions")

	// Check for existing processing extensions
	processingFilter := bson.M{
		"is_active":  true,
		"project_id": projectID,
		"status":     "processing", // Look for status "processing"
	}

	// Count the number of processing extensions
	processingCount, err := collection.CountDocuments(context.TODO(), processingFilter)
	if err != nil {
		return err
	}

	// If any processing extensions exist, do not create a new one
	if processingCount > 0 {
		return util.ErrSynchronizerInProgress
	}

	// If there are no processing extensions, create a new one with status "processing"
	e.Status = "processing" // Set the status to "processing"

	// Insert the new extension
	_, err = collection.InsertOne(context.TODO(), e)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSynchronizerCustomerRecievedExtension updates the "customer_received" field and checks if the status should be set to "completed"
func (ex *ExtensionRepository) UpdateSynchronizerCustomerRecievedExtension(ctx echo.Context, projectID string, customerReceived int) error {
	collection := ex.mongo.Database("otoo").Collection("synchronizer_project_extensions")

	// Define the filter to find the active project with the given project ID
	filter := bson.M{
		"is_active":  true,
		"project_id": projectID,
		"status":     "processing", // Ensure the status is not "processing"
	}

	// First, retrieve the current document to get the totals and the current state of the "received" fields
	var currentDoc domain.DataSynchronizerExtension
	err := collection.FindOne(context.TODO(), filter).Decode(&currentDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no active synchronizer project found for projectID: %s", projectID)
		}
		return err
	}

	// Increment the "customer_received" field
	update := bson.M{
		"$inc": bson.M{
			"customer_received": customerReceived, // Increment customer_received by the provided value
		},
		"$set": bson.M{
			"updated_at": time.Now(), // Update the timestamp for updated_at
		},
	}

	// Perform the update operation
	_, err = collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	// Check if customer_received, order_received, and product_received match the totals
	if currentDoc.CustomerRecieved+customerReceived >= currentDoc.TotalCustomer &&
		currentDoc.OrderReceived >= currentDoc.TotalOrder &&
		currentDoc.ProductReceived >= currentDoc.TotalProduct {

		// If all received values match the totals, update the status to "completed"
		statusUpdate := bson.M{
			"$set": bson.M{
				"status":     "completed",
				"updated_at": time.Now(),
			},
		}

		_, err = collection.UpdateOne(context.TODO(), filter, statusUpdate)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateSynchronizerOrderReceivedExtension updates the "order_received" field in the "synchronizer_project_extensions" collection
func (ex *ExtensionRepository) UpdateSynchronizerOrderReceivedExtension(ctx echo.Context, projectID string, orderReceived int) error {
	collection := ex.mongo.Database("otoo").Collection("synchronizer_project_extensions")

	filter := bson.M{
		"is_active":  true,
		"project_id": projectID,
		"status":     "processing", // Ensure the status is not "processing"
	}

	var currentDoc domain.DataSynchronizerExtension
	err := collection.FindOne(context.TODO(), filter).Decode(&currentDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no active synchronizer project found for projectID: %s", projectID)
		}
		return err
	}

	update := bson.M{
		"$inc": bson.M{
			"order_received": orderReceived, // Increment order_received by the provided value
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	// Check if all received fields meet their total
	if currentDoc.CustomerRecieved >= currentDoc.TotalCustomer &&
		currentDoc.OrderReceived+orderReceived >= currentDoc.TotalOrder &&
		currentDoc.ProductReceived >= currentDoc.TotalProduct {

		statusUpdate := bson.M{
			"$set": bson.M{
				"status":     "completed",
				"updated_at": time.Now(),
			},
		}

		_, err = collection.UpdateOne(context.TODO(), filter, statusUpdate)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateSynchronizerProductReceivedExtension updates the "product_received" field in the "synchronizer_project_extensions" collection
func (ex *ExtensionRepository) UpdateSynchronizerProductReceivedExtension(ctx echo.Context, projectID string, productReceived int) error {
	// Get the collection from the MongoDB database
	collection := ex.mongo.Database("otoo").Collection("synchronizer_project_extensions")

	// Define the filter to find the active project with the given project ID
	filter := bson.M{
		"is_active":  true,
		"project_id": projectID,
		"status":     "processing", // Ensure the status is not "processing"
	}

	// First, retrieve the current document to get the totals and the current state of the "received" fields
	var currentDoc domain.DataSynchronizerExtension
	err := collection.FindOne(context.TODO(), filter).Decode(&currentDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("no active synchronizer project found for projectID: %s", projectID)
		}
		return err
	}

	// Increment the "product_received" field
	update := bson.M{
		"$inc": bson.M{
			"product_received": productReceived, // Increment product_received by the provided value
		},
		"$set": bson.M{
			"updated_at": time.Now(), // Update the timestamp for updated_at
		},
	}

	// Perform the update operation
	_, err = collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	// Check if customer_received, order_received, and product_received match the totals
	if currentDoc.CustomerRecieved >= currentDoc.TotalCustomer &&
		currentDoc.OrderReceived >= currentDoc.TotalOrder &&
		currentDoc.ProductReceived+productReceived >= currentDoc.TotalProduct {

		// If all received values match the totals, update the status to "completed"
		statusUpdate := bson.M{
			"$set": bson.M{
				"status":     "completed",
				"updated_at": time.Now(),
			},
		}

		_, err = collection.UpdateOne(context.TODO(), filter, statusUpdate)
		if err != nil {
			return err
		}
	}

	// Return nil if everything succeeded
	return nil
}

// GetAllSynchronizerProjectExtensions gets all ProjectExtensions
func (ex *ExtensionRepository) GetAllSynchronizerProjectExtensions(ctx echo.Context, projectID string) ([]*domain.DataSynchronizerExtension, error) {
	collection := ex.mongo.Database("otoo").Collection("synchronizer_project_extensions")

	filter := bson.M{
		"project_id": projectID,
		"is_active":  true,
	}

	cursor, err := collection.Find(ctx.Request().Context(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx.Request().Context())

	var projectExtensions []*domain.DataSynchronizerExtension
	for cursor.Next(ctx.Request().Context()) {
		var projectExtension domain.DataSynchronizerExtension
		if err := cursor.Decode(&projectExtension); err != nil {
			return nil, err
		}
		projectExtensions = append(projectExtensions, &projectExtension)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return projectExtensions, nil
}

// GetSynchronizerProjectExtensionByID gets a ProjectExtension by ID
func (ex *ExtensionRepository) GetSynchronizerProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.DataSynchronizerExtension, error) {
	collection := ex.mongo.Database("otoo").Collection("synchronizer_project_extensions")

	filter := bson.M{
		"project_id": projectID,
		"is_active":  true,
		"status":     "processing",
	}

	var projectExtension domain.DataSynchronizerExtension
	err := collection.FindOne(ctx.Request().Context(), filter).Decode(&projectExtension)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &projectExtension, nil
}

// DeleteSynchronizerProjectExtension deletes a ProjectExtension by ID
func (ex *ExtensionRepository) DeleteSynchronizerProjectExtension(ctx echo.Context, extensionID, projectID string) error {
	collection := ex.mongo.Database("otoo").Collection("synchronizer_project_extensions")

	extID, err := primitive.ObjectIDFromHex(extensionID)
	if err != nil {
		return errors.New("Invalid extension ID format")
	}

	filter := bson.M{
		"_id":        extID,
		"project_id": projectID,
		"is_active":  true,
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"is_active":  false,
		},
	}

	result, err := collection.UpdateOne(ctx.Request().Context(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("ProjectExtension not found")
	}

	return nil
}

//////////////////  Courier4u Extension  /////////////////////////

// CreateCourier4uProjectExtension creates a new ProjectExtension
func (ex *ExtensionRepository) CreateCourier4uProjectExtension(ctx echo.Context, projectID string, e *domain.Courier4uExtension) error {
	collection := ex.mongo.Database("otoo").Collection("courier4u_project_extensions")

	filter := bson.M{"extension_id": e.ExtensionID, "is_active": true, "project_id": projectID}
	update := bson.M{"$set": e}

	// Set upsert option to true
	opt := options.Update().SetUpsert(true)

	// Perform the upsert operation
	_, err := collection.UpdateOne(context.TODO(), filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCourier4uProjectExtensions gets all ProjectExtensions
func (ex *ExtensionRepository) GetAllCourier4uProjectExtensions(ctx echo.Context, projectID string) ([]*domain.Courier4uExtension, error) {
	collection := ex.mongo.Database("otoo").Collection("courier4u_project_extensions")

	filter := bson.M{
		"project_id": projectID,
		"is_active":  true,
	}

	cursor, err := collection.Find(ctx.Request().Context(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx.Request().Context())

	var projectExtensions []*domain.Courier4uExtension
	for cursor.Next(ctx.Request().Context()) {
		var projectExtension domain.Courier4uExtension
		if err := cursor.Decode(&projectExtension); err != nil {
			return nil, err
		}
		projectExtensions = append(projectExtensions, &projectExtension)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return projectExtensions, nil
}

// GetCourier4uProjectExtensionByID gets a ProjectExtension by ID
func (ex *ExtensionRepository) GetCourier4uProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.Courier4uExtension, error) {
	collection := ex.mongo.Database("otoo").Collection("courier4u_project_extensions")

	filter := bson.M{
		"project_id":   projectID,
		"extension_id": extensionID,
		"is_active":    true,
	}

	var projectExtension domain.Courier4uExtension
	err := collection.FindOne(ctx.Request().Context(), filter).Decode(&projectExtension)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &projectExtension, nil
}

// DeleteCourier4uProjectExtension deletes a ProjectExtension by ID
func (ex *ExtensionRepository) DeleteCourier4uProjectExtension(ctx echo.Context, extensionID, projectID string) error {
	collection := ex.mongo.Database("otoo").Collection("courier4u_project_extensions")

	extID, err := primitive.ObjectIDFromHex(extensionID)
	if err != nil {
		return errors.New("Invalid extension ID format")
	}

	filter := bson.M{
		"_id":          extID,
		"project_id":   projectID,
		"extension_id": extensionID,
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"is_active":  false,
		},
	}

	result, err := collection.UpdateOne(ctx.Request().Context(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("ProjectExtension not found")
	}

	return nil
}
