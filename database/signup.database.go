package database

import (
	"context"
	"errors"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *MongodbRepository) GetSignups(ctx context.Context, page int) ([]models.Signup, error) {
	l := int64(5)
	skip := int64(page)
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	result, err := repo.DB.Collection("inscriptions").Find(ctx, bson.M{"completed": false}, &fOpt)
	if err != nil {
		return nil, err
	}

	var signups []models.Signup

	err = result.All(ctx, &signups)
	if err != nil {
		return nil, err
	}

	if len(signups) == 0 {
		return nil, errors.New("Signups not found")
	}

	return signups, nil
}
