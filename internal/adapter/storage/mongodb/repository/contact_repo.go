package repository

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactRepository struct {
	mongo *mongo.Client
}

func NewContactRepository(mongo *mongo.Client) *ContactRepository {
	return &ContactRepository{
		mongo,
	}
}

func (repo ContactRepository) InsertContact(ctx echo.Context, data *domain.Contact) error {
	coll := repo.mongo.Database("otoo").Collection("contacts")
	coll.InsertOne(context.TODO(), data)
	return nil
}
