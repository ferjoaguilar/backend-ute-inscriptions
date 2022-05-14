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

func TestCreateUser(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {

		id1 := primitive.NewObjectID()

		first := mtest.CreateCursorResponse(1, "users.utec", mtest.FirstBatch, bson.D{
			primitive.E{Key: "id", Value: id1},
			primitive.E{Key: "email", Value: "stefanylue123@gmail.com"},
			primitive.E{Key: "username", Value: "stefanylue123"},
			primitive.E{Key: "password", Value: "vanilla12345"},
			primitive.E{Key: "permissions", Value: "manager"},
			primitive.E{Key: "disabled", Value: true},
			primitive.E{Key: "created_at", Value: time.Now()},
		})

		killCursors := mtest.CreateCursorResponse(0, "users.utec", mtest.NextBatch)
		mt.AddMockResponses(first, killCursors, mtest.CreateSuccessResponse())

		mockdb := mt.DB
		repo := database.MongodbRepository{DB: mockdb}

		insertedUser, err := repo.CreateUser(context.Background(), &models.User{
			Email:       "feraguilar6985@gmail.com",
			Username:    "feraguilar6985",
			Password:    "password2365889",
			Permissions: "manager",
		})

		if err != nil {
			t.Errorf("TestCreateUser(success) was incorrect, got %v, want %v", err, insertedUser.InsertedID)
		}
	})

	mt.Run("error duplicate user", func(mt *mtest.T) {

		id1 := primitive.NewObjectID()

		first := mtest.CreateCursorResponse(1, "users.utec", mtest.FirstBatch, bson.D{
			primitive.E{Key: "id", Value: id1},
			primitive.E{Key: "email", Value: "stefanylue123@gmail.com"},
			primitive.E{Key: "username", Value: "stefanylue123"},
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
		_, err := repo.CreateUser(context.Background(), &models.User{})

		if mongo.IsDuplicateKeyError(err) != true {
			t.Errorf("TestCreateUser(error duplicate user) was incorrect, got %v, want %v", mongo.IsDuplicateKeyError(err), true)
		}

	})
}

func TestLoginUser(t *testing.T) {
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {

		expectedUser := models.User{
			ID:        primitive.NewObjectID(),
			Email:     "stefanylue123@gmail.com",
			Username:  "stefanylue123",
			Password:  "password2365889",
			Disable:   false,
			CreatedAt: time.Now(),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "user.utec", mtest.FirstBatch, bson.D{
			primitive.E{Key: "_id", Value: expectedUser.ID},
			primitive.E{Key: "email", Value: expectedUser.Email},
			primitive.E{Key: "password", Value: expectedUser.Password},
			primitive.E{Key: "disabled", Value: false},
			primitive.E{Key: "createdat", Value: expectedUser.CreatedAt},
		}))

		mockdb := mt.DB
		repo := database.MongodbRepository{DB: mockdb}
		userResponse, err := repo.FindUserByEmail(context.Background(), expectedUser.Email)
		if err != nil {
			t.Errorf("TestLoginUser error, got %v", err)
		}

		if expectedUser.Email != userResponse.Email {
			t.Errorf("TestLoginUser(success) was incorrect, got %v, want %v", userResponse.Email, expectedUser.Email)
		}
	})
}

func TestDisabledUser(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		expectedUser := models.User{
			ID: primitive.NewObjectID(),
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		mockdb := mt.DB
		repo := database.MongodbRepository{DB: mockdb}
		stringObjectId := expectedUser.ID.Hex()
		_, err := repo.DisabledUser(context.Background(), stringObjectId)
		if err != nil {
			t.Errorf("TestDisabledUser(success) was incorrect, got %v, want %v", err, expectedUser.ID)
		}
	})
}

func TestGetManagers(t *testing.T) {

	// Start settings
	t.Parallel()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	// Mocks
	expectedUser := bson.D{
		primitive.E{Key: "id", Value: primitive.NewObjectID()},
		primitive.E{Key: "email", Value: "stefanylue123@gmail.com"},
		primitive.E{Key: "username", Value: "stefanylue123"},
		primitive.E{Key: "password", Value: "vanilla12345"},
		primitive.E{Key: "permissions", Value: "manager"},
		primitive.E{Key: "disabled", Value: true},
		primitive.E{Key: "created_at", Value: time.Now()},
	}

	// Table driven test
	testCases := []struct {
		Name     string
		UserMock bson.D
		Expected interface{}
	}{
		{
			Name:     "Success Get users managers",
			UserMock: expectedUser,
			Expected: 1,
		},
	}

	for _, tc := range testCases {
		mt.Run(tc.Name, func(mt *mtest.T) {
			first := mtest.CreateCursorResponse(1, "users.utec", mtest.FirstBatch, tc.UserMock)

			killCursors := mtest.CreateCursorResponse(0, "user.utec", mtest.NextBatch)
			mt.AddMockResponses(first, killCursors)

			mockdb := mt.DB
			repo := database.MongodbRepository{DB: mockdb}

			usersManager, _ := repo.GetManagers(context.Background())

			if len(usersManager) <= 0 {
				t.Errorf("Test manager was incorrect, got %d, want %d", len(usersManager), tc.Expected)
			}
		})
	}

}
