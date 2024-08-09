package repository

import (
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// GetExtensionsByID gets a Extension by ID
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

// GetExtensionsByCode gets a Extension by code
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

	_, err := collection.InsertOne(ctx.Request().Context(), projectExtension)
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

// GetProjectExtensionsByID gets a ProjectExtension by ID
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
