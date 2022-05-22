package repository_test

import (
	"context"
	"testing"

	"github.com/ferjoaguilar/backend-utec-inscriptions/models"
	"github.com/ferjoaguilar/backend-utec-inscriptions/repository"
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
		Name            string
		Input           string
		ExpectedSuccess []models.SignupLookup
		ExpectedError   error
	}{
		{
			Name:            "Success Get signups",
			Input:           "approved",
			ExpectedSuccess: []models.SignupLookup{},
			ExpectedError:   nil,
		},
	}

	ctx := context.Background()
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			signupRepo.On("GetSignups", ctx, tc.Input).Return(tc.ExpectedSuccess, nil)
			_, err := repository.GetSignups(ctx, tc.Input)
			if err != tc.ExpectedError {
				t.Errorf("Get signups incorrect got %v want %v", err, tc.ExpectedError)
			}
		})
	}
}

func TestChangeStatus(t *testing.T) {
	testCases := []struct {
		Name            string
		InputOne        string
		InputTwo        string
		ExpectedSuccess string
		ExpectedError   error
	}{
		{
			Name:            "Success Change status",
			InputOne:        "627af6017c27bd2f8488c03f",
			InputTwo:        "pending",
			ExpectedSuccess: "Student status updated successfully",
			ExpectedError:   nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Parallel()
		signupRepo.On("ChangeStatus", ctx, tc.InputOne, tc.InputTwo).Return(tc.ExpectedSuccess, nil)
		_, err := repository.ChangeStatus(ctx, tc.InputOne, tc.InputTwo)
		if err != tc.ExpectedError {
			t.Errorf("Change status incorrect got %v want %v", err, tc.ExpectedError)
		}
	}

}

func TestGetSignupById(t *testing.T) {
	testCases := []struct {
		Name            string
		Input           string
		ExpectedSuccess models.Signup
		ExpectedError   error
	}{
		{
			Name:            "Success Get signup by Id",
			Input:           "62894506ddaf96a5df37244c",
			ExpectedSuccess: models.Signup{},
			ExpectedError:   nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Parallel()
		signupRepo.On("GetSignupById", ctx, tc.Input).Return(&tc.ExpectedSuccess, nil)
		_, err := repository.GetSignupById(ctx, tc.Input)
		if err != tc.ExpectedError {
			t.Errorf("Get signup by Id incorrect got %v want %v", err, tc.ExpectedError)
		}
	}
}
