package repository_test

import (
	"context"
	"testing"

	"github.com/ferjoaguilar/backend-utec-inscriptions/models"
	"github.com/ferjoaguilar/backend-utec-inscriptions/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
			userRepo.On("CreateUser", ctx, &tc.Input).Return("New user created", nil)
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
			Name:            "Success Find user by email",
			Input:           "estefany.lue99@gmail.com",
			ExpectedSuccess: models.User{},
			ExpectedError:   nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			userRepo.On("FindUserByEmail", ctx, tc.Input).Return(&tc.ExpectedSuccess, nil)
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
			userRepo.On("DisableUser", ctx, tc.Input.Hex()).Return("User updated successfully", nil)
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
			userRepo.On("GetManagers", ctx).Return(tc.ExpectedSuccess, nil)
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
			userRepo.On("GetUserById", ctx, tc.Input).Return(&tc.ExpectedSuccess, nil)
			_, err := repository.GetUserById(ctx, tc.Input)
			if err != tc.ExpectedError {
				t.Errorf("Get user by Id incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}

}
