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
