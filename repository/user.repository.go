package repository

import (
	"context"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	DisabledUser(ctx context.Context, id string) (string, error)
	GetManagers(ctx context.Context) ([]models.User, error)
}

var UserImplementation UserRepository

func SetUserRepository(repository UserRepository) {
	UserImplementation = repository
}

func CreateUser(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	return UserImplementation.CreateUser(ctx, user)
}

func FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return UserImplementation.FindUserByEmail(ctx, email)
}

func DisabledUser(ctx context.Context, id string) (string, error) {
	return UserImplementation.DisabledUser(ctx, id)
}

func GetManagers(ctx context.Context) ([]models.User, error) {
	return UserImplementation.GetManagers(ctx)
}
