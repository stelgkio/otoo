package repository

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PaymentRepository struct
type PaymentRepository struct {
	mongo *mongo.Client
}

// NewPaymentRepository constructor
func NewPaymentRepository(mongo *mongo.Client) *PaymentRepository {
	return &PaymentRepository{
		mongo,
	}
}

// CreatePayment  add new Payment for user and project
func (nr PaymentRepository) CreatePayment(ctx echo.Context, data *domain.Payment) error {
	coll := nr.mongo.Database("otoo").Collection("payment")
	coll.InsertOne(context.TODO(), data)
	return nil
}

// UpdatePayment  update to read
func (nr PaymentRepository) UpdatePayment(ctx echo.Context, data *domain.Payment) error {
	coll := nr.mongo.Database("otoo").Collection("payment")
	coll.InsertOne(context.TODO(), data)
	return nil
}

// FindPayment find all Payment by projectID
func (nr PaymentRepository) FindPayment(projectID string, size, page int, sort, direction string, filters bool) ([]*domain.Payment, error) {
	coll := nr.mongo.Database("otoo").Collection("payment")
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
		filter = bson.M{"projectId": projectID, "is_active": true}

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

	var Payments []*domain.Payment
	for cursor.Next(ctx) {
		var Payment domain.Payment
		if err := cursor.Decode(&Payment); err != nil {
			return nil, err
		}
		Payments = append(Payments, &Payment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return Payments, nil
}

// DeletePayment delete Payment
func (nr PaymentRepository) DeletePayment(ctx echo.Context, projectID, PaymentID string) error {
	coll := nr.mongo.Database("otoo").Collection("payment")

	filter := bson.M{"project_id": projectID, "is_active": true}
	update := bson.M{"$set": bson.M{"is_active": false, "deleted_at": time.Now().UTC(), "is_read": true}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// PaymentCount get number of payments
func (nr PaymentRepository) PaymentCount(projectID string) (int64, error) {
	coll := nr.mongo.Database("otoo").Collection("payment")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.M{"projectId": projectID, "is_active": true}
	res, err := coll.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}
	return res, nil
}
