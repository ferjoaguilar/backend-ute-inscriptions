package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/database"
	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		mockdb := mt.DB
		repo := database.MongodbRepository{DB: mockdb}
		insertedUser, err := repo.CreateUser(context.Background(), &models.User{
			Email:     "stefanylue123@gmail.com",
			Username:  "stefanylue123",
			Password:  "password2365889",
			Disable:   false,
			CreatedAt: time.Now(),
		})

		if err != nil {
			t.Errorf("TestCreateUser(success) was incorrect, got %v, want %v", err, insertedUser.InsertedID)
		}
	})

	mt.Run("error duplicate user", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))
		mockdb := mt.DB
		repo := database.MongodbRepository{DB: mockdb}
		_, err := repo.CreateUser(context.Background(), &models.User{})

		if mongo.IsDuplicateKeyError(err) != true {
			t.Errorf("TestCreateUser(error duplicate user) was incorrect, got %v, want %v", mongo.IsDuplicateKeyError(err), true)
		}

	})

}
