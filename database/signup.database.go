package database

import (
	"context"
	"errors"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *MongodbRepository) GetSignups(ctx context.Context) ([]models.Signup, error) {

	result, err := repo.DB.Collection("inscriptions").Find(ctx, bson.M{"completed": false})
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

func (repo *MongodbRepository) CompleteSignup(ctx context.Context, id string) (string, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	result := repo.DB.Collection("inscriptions").FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.D{
		{"$set", bson.D{{"completed", true}}},
	})

	if result.Err() != nil {
		return "", result.Err()
	}

	return "Signup completed successfully", nil
}
