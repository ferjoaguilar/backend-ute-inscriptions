package repository

import (
	"context"

	"github.com/ferjoaguilar/backend-utec-inscriptions/models"
)

//go:generate mockery --name=SignupRepository
type SignupRepository interface {
	CreateSignup(ctx context.Context, signup *models.Signup) (string, error)
	GetSignups(ctx context.Context, status string) ([]models.SignupLookup, error)
	ChangeStatus(ctx context.Context, id string, status string) (string, error)
	GetSignupById(ctx context.Context, id string) (*models.Signup, error)
}

var SignupImplementations SignupRepository

func SetSignupRepository(repository SignupRepository) {
	SignupImplementations = repository
}

func CreateSignup(ctx context.Context, signup *models.Signup) (string, error) {
	return SignupImplementations.CreateSignup(ctx, signup)
}

func GetSignups(ctx context.Context, status string) ([]models.SignupLookup, error) {
	return SignupImplementations.GetSignups(ctx, status)
}

func ChangeStatus(ctx context.Context, id string, status string) (string, error) {
	return SignupImplementations.ChangeStatus(ctx, id, status)
}

func GetSignupById(ctx context.Context, id string) (*models.Signup, error) {
	return SignupImplementations.GetSignupById(ctx, id)
}
