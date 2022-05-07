package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbRepository struct {
	DB *mongo.Database
}

func NewMongoRepository(url string) (*MongodbRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &MongodbRepository{DB: client.Database("utec-inscription")}, nil
}
