package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/database"
	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type userMock struct {
	Key   string
	Value interface{}
}

// Create user testing
func TestCreateUser(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		mockUser := mtest.CreateCursorResponse(1, "users.utec", mtest.FirstBatch, bson.D{
			primitive.E{Key: "id", Value: primitive.NewObjectID()},
			primitive.E{Key: "email", Value: "estefany.lue99@gmail.com"},
			primitive.E{Key: "username", Value: "estefany.lue99"},
			primitive.E{Key: "password", Value: "vanilla12345"},
			primitive.E{Key: "permissions", Value: "manager"},
			primitive.E{Key: "disabled", Value: true},
			primitive.E{Key: "created_at", Value: time.Now()},
		})

		killCursors := mtest.CreateCursorResponse(0, "users.utec", mtest.NextBatch)
		mt.AddMockResponses(mockUser, killCursors, mtest.CreateSuccessResponse())

		mockdb := mt.DB
		repo := database.MongodbRepository{DB: mockdb}

		insertedUser, err := repo.CreateUser(context.Background(), &models.User{
			Email:       "feraguilar6985@gmail.com",
			Username:    "feraguilar6985",
			Password:    "password2365889",
			Permissions: "manager",
		})

		if err != nil {
			t.Errorf(err.Error())
			return
		}
		stringObjectId := insertedUser.InsertedID.(primitive.ObjectID).Hex()
		mongoObjectId := primitive.IsValidObjectID(stringObjectId)

		if !mongoObjectId {
			t.Errorf("Create user was incorrect, got %t want %t", mongoObjectId, true)
		}

	})

	mt.Run("error duplicate user", func(mt *mtest.T) {

		id1 := primitive.NewObjectID()
		first := mtest.CreateCursorResponse(1, "users.utec", mtest.FirstBatch, bson.D{
			primitive.E{Key: "id", Value: id1},
			primitive.E{Key: "name", Value: "estefany.lue99@gmail.com"},
			primitive.E{Key: "username", Value: "estefany.lue99"},
			primitive.E{Key: "password", Value: "vanilla12345"},
			primitive.E{Key: "permissions", Value: "manager"},
			primitive.E{Key: "disabled", Value: true},
			primitive.E{Key: "created_at", Value: time.Now()},
		})

		killCursors := mtest.CreateCursorResponse(0, "user.utec", mtest.NextBatch)

		mt.AddMockResponses(first, killCursors, mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))
		mockdb := mt.DB
		repo := database.MongodbRepository{DB: mockdb}
		insertedUser, err := repo.CreateUser(context.Background(), &models.User{})

		if mongo.IsDuplicateKeyError(err) != true {
			t.Errorf(err.Error())
		}

		if insertedUser != nil {
			t.Errorf("Create user was incorrect, got %v want %v", insertedUser, "duplicate key error")
		}

	})
}
