package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"github.com/snowball-devs/backend-utec-inscriptions/repository"
	"github.com/snowball-devs/backend-utec-inscriptions/repository/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo *mocks.UserRepository

// Init mock instance to unit testing
func TestMain(m *testing.M) {
	repo = &mocks.UserRepository{}
	repository.SetUserRepository(repo)
	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {

	testCases := []struct {
		Name          string
		Input         models.User
		ExpectedError error
	}{
		{
			Name:          "Success Create new user",
			Input:         models.User{},
			ExpectedError: nil,
		},
	}

	ctx := context.Background()
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.On("CreateUser", ctx, &tc.Input).Return("New user created", nil)
			_, err := repository.CreateUser(ctx, &tc.Input)
			if err != tc.ExpectedError {
				t.Errorf("Create user incorrect, go %v want %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestFindUserByEmail(t *testing.T) {
	testCases := []struct {
		Name            string
		Input           string
		ExpectedSuccess models.User
		ExpectedError   error
	}{
		{
			Name:  "Success Find user by email",
			Input: "estefany.lue99@gmail.com",
			ExpectedSuccess: models.User{
				ID:          primitive.NewObjectID(),
				Email:       "estefany.lue99@gmail.com",
				Username:    "estefany.lue99",
				Password:    "vanillagolang123",
				Permissions: "manager",
				Disable:     false,
				CreatedAt:   time.Now(),
			},
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.On("FindUserByEmail", ctx, tc.Input).Return(&tc.ExpectedSuccess, nil)
			_, err := repository.FindUserByEmail(ctx, tc.Input)
			if err != tc.ExpectedError {
				t.Errorf("Find user by email incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}
}

func TestDisableUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Input         primitive.ObjectID
		ExpectedError error
	}{
		{
			Name:          "Success disable user",
			Input:         primitive.NewObjectID(),
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.On("DisableUser", ctx, tc.Input.Hex()).Return("User updated successfully", nil)
			_, err := repository.DisableUser(ctx, tc.Input.Hex())
			if err != tc.ExpectedError {
				t.Errorf("Disable user incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}
}

func TestGetManagers(t *testing.T) {
	testCases := []struct {
		Name            string
		ExpectedSuccess []models.User
		ExpectedError   error
	}{
		{
			Name:            "Success get users managers",
			ExpectedSuccess: nil,
			ExpectedError:   nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.On("GetManagers", ctx).Return(tc.ExpectedSuccess, nil)
			_, err := repository.GetManagers(ctx)
			if err != tc.ExpectedError {
				t.Errorf("Get managers incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}

}

func TestGetUserById(t *testing.T) {
	testCases := []struct {
		Name            string
		Input           string
		ExpectedSuccess models.User
		ExpectedError   error
	}{
		{
			Name:            "Success Find user by Id",
			Input:           "627af6017c27bd2f8488c03f",
			ExpectedSuccess: models.User{},
			ExpectedError:   nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.On("GetUserById", ctx, tc.Input).Return(&tc.ExpectedSuccess, nil)
			_, err := repository.GetUserById(ctx, tc.Input)
			if err != tc.ExpectedError {
				t.Errorf("Get user by Id incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}

}
