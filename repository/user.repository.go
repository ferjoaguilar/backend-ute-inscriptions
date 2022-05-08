package repository

import (
	"context"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
}

var UserImplementation UserRepository

func SetUserRepository(repository UserRepository) {
	UserImplementation = repository
}

func CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	return UserImplementation.CreateUser(ctx, user)
}
