package repository

import (
	"context"

	"github.com/snowball-devs/backend-utec-inscriptions/models"
)

//go:generate mockery --name=SignupRepository
type SignupRepository interface {
	CreateSignup(ctx context.Context, signup *models.Signup) (string, error)
	GetSignups(ctx context.Context, page int) ([]models.Signup, error)
	CompleteSignup(ctx context.Context, id string) (string, error)
}

var SignupImplementations SignupRepository

func SetSignupRepository(repository SignupRepository) {
	SignupImplementations = repository
}

func CreateSignup(ctx context.Context, signup *models.Signup) (string, error) {
	return SignupImplementations.CreateSignup(ctx, signup)
}

func GetSignups(ctx context.Context, page int) ([]models.Signup, error) {
	return SignupImplementations.GetSignups(ctx, page)
}

func CompleteSignup(ctx context.Context, id string) (string, error) {
	return SignupImplementations.CompleteSignup(ctx, id)
}
