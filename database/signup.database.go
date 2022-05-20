package database

import (
	"context"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
)

func (repo *MongodbRepository) CreateSignup(ctx context.Context, signup *models.Signup) (string, error) {

	signup.Completed = false
	signup.CreatedAt = time.Now()

	_, err := repo.DB.Collection("inscriptions").InsertOne(ctx, signup)

	if err != nil {
		return "", err
	}

	return "New inscription created", nil

}
