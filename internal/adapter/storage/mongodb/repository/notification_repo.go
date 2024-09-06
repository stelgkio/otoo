package repository

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
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
	coll := nr.mongo.Database("otoo").Collection("notifications")
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
func (nr NotificationRepository) FindNotification(projectID string, size, page int, sort, direction string) ([]*domain.Notification, error) {
	coll := nr.mongo.Database("otoo").Collection("notifications")
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

	pipeline = append(pipeline, bson.D{{Key: "$sort",
		Value: bson.D{{Key: sort, Value: sortOrder}}}})

	// Pagination: skip and limit
	pipeline = append(pipeline, bson.D{{Key: "$skip", Value: int64(size * (page - 1))}})
	pipeline = append(pipeline, bson.D{{Key: "$limit", Value: int64(size)}})

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
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
