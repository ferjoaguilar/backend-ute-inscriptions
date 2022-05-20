package database

import (
	"context"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
)

func (repo *MongodbRepository) CreateSignup(ctx context.Context, inscription models.Signup) (string, error) {

	inscription.Completed = false
	inscription.CreatedAt = time.Now()

	_, err := repo.DB.Collection("inscriptions").InsertOne(ctx, inscription)

	if err != nil {
		return "", err
	}

	return "New inscription created", nil

}
