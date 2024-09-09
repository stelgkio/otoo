package repository

import (
	"context"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// NotificationRepository struct
type NotificationRepository struct {
	mongo *mongo.Client
}

// NewNotificationRepository constructor
func NewNotificationRepository(mongo *mongo.Client) *NotificationRepository {
	return &NotificationRepository{
		mongo,
	}
}

// CreateNotification  add new notification for user and project
func (nr NotificationRepository) CreateNotification(ctx echo.Context, data *domain.Notification) error {
	coll := nr.mongo.Database("otoo").Collection("notitications")
	coll.InsertOne(context.TODO(), data)
	return nil
}

// UpdateNotification  update to read
func (nr NotificationRepository) UpdateNotification(ctx echo.Context, data *domain.Notification) error {
	coll := nr.mongo.Database("otoo").Collection("notifications")
	coll.InsertOne(context.TODO(), data)
	return nil
}

// FindNotification find all notification by projectID
func (nr NotificationRepository) FindNotification(projectID string, size, page int, sort, direction string, filters bool) ([]*domain.Notification, error) {
	coll := nr.mongo.Database("otoo").Collection("notifications")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	sortOrder := 1
	if direction == "desc" {
		sortOrder = -1
	} else if direction == "" {
		sortOrder = -1
	}

	filter := bson.M{}
	if filters {
		filter = bson.M{"project_id": projectID, "is_active": true, "is_read": false}

	}
	// Match active customers by projectID

	// Create an aggregation pipeline
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sort", Value: bson.D{{Key: sort, Value: sortOrder}}}},
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

	var notifications []*domain.Notification
	for cursor.Next(context.TODO()) {
		var notification domain.Notification
		if err := cursor.Decode(&notification); err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

// DeleteNotification delete Notification
func (nr NotificationRepository) DeleteNotification(ctx echo.Context, projectID, notificationID string) error {
	coll := nr.mongo.Database("otoo").Collection("notifications")
	ntfID, err := primitive.ObjectIDFromHex(notificationID)
	if err != nil {
		return errors.New("Invalid extension ID format")
	}
	filter := bson.M{"project_id": projectID, "is_active": true, "is_read": false, "_id": ntfID}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now().UTC(), "is_read": true}}
	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
