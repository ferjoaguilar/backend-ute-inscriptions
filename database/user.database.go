package database

import (
	"context"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongodbRepository) CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	newUser := models.User{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Disable:   false,
		CreatedAt: time.Now(),
	}

	result, err := repo.DB.Collection("users").InsertOne(ctx, newUser)

	if err != nil {
		return nil, err
	}
	return result, nil
}
