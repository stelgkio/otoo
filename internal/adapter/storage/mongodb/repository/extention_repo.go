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

// ExtentionRepository represents the repository for Extention-related operations
type ExtentionRepository struct {
	mongo *mongo.Client
}

// NewExtentionRepository creates a new Extention repository instance
func NewExtentionRepository(mongo *mongo.Client) *ExtentionRepository {
	return &ExtentionRepository{
		mongo,
	}
}

// CreateExtention creates a new Extention
func (ex *ExtentionRepository) CreateExtention(ctx echo.Context, c *domain.Extention) error {
	panic("unimplemented")
}

// GetAllExtentions gets all Extentions
func (ex *ExtentionRepository) GetAllExtentions(ctx echo.Context) ([]*domain.Extention, error) {
	collection := ex.mongo.Database("otoo").Collection("extentions")
	cursor, err := collection.Find(ctx.Request().Context(), bson.M{"is_active": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx.Request().Context())

	var extentions []*domain.Extention
	for cursor.Next(ctx.Request().Context()) {
		var extention domain.Extention
		if err := cursor.Decode(&extention); err != nil {
			return nil, err
		}
		extentions = append(extentions, &extention)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return extentions, nil
}

// GetExtentionsByID gets a Extention by ID
func (ex *ExtentionRepository) GetExtentionsByID(ctx echo.Context, extensionID string) (*domain.Extention, error) {
	collection := ex.mongo.Database("otoo").Collection("extentions")

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
	var extention domain.Extention

	// Execute the query with the filter
	err = collection.FindOne(ctx.Request().Context(), filter).Decode(&extention)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Extention not found")
		}
		return nil, err
	}

	// Return the found extension
	return &extention, nil
}

// DeleteExtention deletes a Extention by ID
func (ex *ExtentionRepository) DeleteExtention(ctx echo.Context, extensionID string) error {
	collection := ex.mongo.Database("otoo").Collection("extentions")
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

//////////////////PROJECT EXTENTIONS/////////////////////////

// CreateProjectExtention creates a new ProjectExtention
func (ex *ExtentionRepository) CreateProjectExtention(ctx echo.Context, projectID string, e *domain.Extention) error {
	collection := ex.mongo.Database("otoo").Collection("project_extentions")

	projectExtention := &domain.ProjectExtention{
		ID:          primitive.NewObjectID(),
		Title:       e.Title,
		Description: e.Description,
		Code:        e.Code,
		ProjectID:   projectID,
		ExtentionID: e.ID.Hex(),
		CreatedAt:   time.Now().UTC(),
		IsActive:    true,
	}

	_, err := collection.InsertOne(ctx.Request().Context(), projectExtention)
	if err != nil {
		return err
	}
	return nil
}

// GetAllProjectExtentions gets all ProjectExtentions
func (ex *ExtentionRepository) GetAllProjectExtentions(ctx echo.Context, projectID string) ([]*domain.ProjectExtention, error) {
	collection := ex.mongo.Database("otoo").Collection("project_extentions")

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

	var projectExtentions []*domain.ProjectExtention
	for cursor.Next(ctx.Request().Context()) {
		var projectExtention domain.ProjectExtention
		if err := cursor.Decode(&projectExtention); err != nil {
			return nil, err
		}
		projectExtentions = append(projectExtentions, &projectExtention)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return projectExtentions, nil
}

// GetProjectExtentionsByID gets a ProjectExtention by ID
func (ex *ExtentionRepository) GetProjectExtentionsByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtention, error) {
	collection := ex.mongo.Database("otoo").Collection("project_extentions")

	extID, err := primitive.ObjectIDFromHex(extensionID)
	if err != nil {
		return nil, errors.New("Invalid extension ID format")
	}

	filter := bson.M{
		"_id":        extID,
		"project_id": projectID,
		"is_active":  true,
	}

	var projectExtention domain.ProjectExtention
	err = collection.FindOne(ctx.Request().Context(), filter).Decode(&projectExtention)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &projectExtention, nil
}

// DeleteProjectExtention deletes a ProjectExtention by ID
func (ex *ExtentionRepository) DeleteProjectExtention(ctx echo.Context, extensionID, projectID string) error {
	collection := ex.mongo.Database("otoo").Collection("project_extentions")

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
		return errors.New("ProjectExtention not found")
	}

	return nil
}
