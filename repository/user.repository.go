package repository

import (
	"context"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
)

//go:generate mockery --name=UserRepository
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	DisableUser(ctx context.Context, id string) (string, error)
	GetManagers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
}

var UserImplementation UserRepository

func SetUserRepository(repository UserRepository) {
	UserImplementation = repository
}

func CreateUser(ctx context.Context, user *models.User) (string, error) {
	return UserImplementation.CreateUser(ctx, user)
}

func FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return UserImplementation.FindUserByEmail(ctx, email)
}

func DisableUser(ctx context.Context, id string) (string, error) {
	return UserImplementation.DisableUser(ctx, id)
}

func GetManagers(ctx context.Context) ([]models.User, error) {
	return UserImplementation.GetManagers(ctx)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return UserImplementation.GetUserById(ctx, id)
}
