package repository_test

import (
	"context"
	"testing"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"github.com/snowball-devs/backend-utec-inscriptions/repository"
)

func TestCreateSignup(t *testing.T) {
	testCases := []struct {
		Name          string
		Input         models.Signup
		ExpectedError error
	}{
		{
			Name:          "Success Create signup",
			Input:         models.Signup{},
			ExpectedError: nil,
		},
	}

	ctx := context.Background()
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			signupRepo.On("CreateSignup", ctx, &tc.Input).Return("New inscription created", nil)
			_, err := repository.CreateSignup(ctx, &tc.Input)
			if err != tc.ExpectedError {
				t.Errorf("Create signup incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}
}

func TestGetSignups(t *testing.T) {
	testCases := []struct {
		Name string

		ExpectedSuccess []models.Signup
		ExpectedError   error
	}{
		{
			Name:            "Success Get signups",
			ExpectedSuccess: []models.Signup{},
			ExpectedError:   nil,
		},
	}

	ctx := context.Background()
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			signupRepo.On("GetSignups", ctx).Return(tc.ExpectedSuccess, nil)
			_, err := repository.GetSignups(ctx)
			if err != tc.ExpectedError {
				t.Errorf("Get signups incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}
}

func TestCompleteSignup(t *testing.T) {
	testCases := []struct {
		Name            string
		Input           string
		ExpectedSuccess string
		ExpectedError   error
	}{
		{
			Name:            "Success Complete signup",
			Input:           "6287ca3afc46f44f97f656c3",
			ExpectedSuccess: "Signup completed successfully",
			ExpectedError:   nil,
		},
	}

	ctx := context.Background()
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			signupRepo.On("CompleteSignup", ctx, tc.Input).Return(tc.ExpectedSuccess, nil)
			_, err := repository.CompleteSignup(ctx, tc.Input)
			if err != tc.ExpectedError {
				t.Errorf("Completed signup incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}

}
