package database

import (
	"context"
	"errors"
	"time"

	"github.com/ferjoaguilar/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongodbRepository) CreateSignup(ctx context.Context, signup *models.Signup) (string, error) {

	signup.Status = "pending"
	signup.CreatedAt = time.Now()

	_, err := repo.DB.Collection("inscriptions").InsertOne(ctx, signup)

	if err != nil {
		return "", err
	}

	return "New inscription created", nil

}

func (repo *MongodbRepository) GetSignups(ctx context.Context, status string) ([]models.SignupLookup, error) {

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "users"}, {"localField", "user"}, {"foreignField", "_id"}, {"as", "user"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$user"}, {"preserveNullAndEmptyArrays", false}}}}
	matchStage := bson.D{{"$match", bson.D{{"status", status}}}}

	result, err := repo.DB.Collection("inscriptions").Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage, matchStage})
	if err != nil {
		return nil, err
	}

	var signups []models.SignupLookup

	err = result.All(ctx, &signups)
	if err != nil {
		return nil, err
	}

	if len(signups) == 0 {
		return nil, errors.New("Signups not found")
	}

	return signups, nil
}

func (repo *MongodbRepository) ChangeStatus(ctx context.Context, id string, status string) (string, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	result := repo.DB.Collection("inscriptions").FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.D{
		{"$set", bson.D{{"status", status}}},
	})

	if result.Err() != nil {
		return "", result.Err()
	}
	return "Student status updated successfully", nil
}

func (repo *MongodbRepository) GetSignupById(ctx context.Context, id string) (*models.Signup, error) {

	var signup *models.Signup
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = repo.DB.Collection("inscriptions").FindOne(ctx, bson.M{"user": objId}).Decode(&signup)

	if err != nil {
		return nil, err
	}

	return signup, nil
}
