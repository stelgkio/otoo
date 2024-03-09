package mongo

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Dbconnect -> connects mongo
func MongoDbConnect(mongoUrl string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoUrl)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		slog.Error("⛒ Connection Failed to Database")
		slog.Error("Fail to connecto to mongo", slog.Any("error", err))
		return nil, err
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error("⛒ Connection Failed to Database")
		slog.Error("Fail to connecto to mongo", slog.Any("error", err))
		return nil, err
	}
	slog.Info("⛁ Connected to Mongo Database")
	coll := client.Database("otoo").Collection("woocommerce_orders")
	coll.Database().Client()

	return client, nil
}
