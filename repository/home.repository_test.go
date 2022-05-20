package repository_test

import (
	"os"
	"testing"

	"github.com/snowball-devs/backend-utec-inscriptions/repository"
	"github.com/snowball-devs/backend-utec-inscriptions/repository/mocks"
)

var userRepo *mocks.UserRepository
var signupRepo *mocks.SignupRepository

// Init mock instance to unit testing
func TestMain(m *testing.M) {
	userRepo = &mocks.UserRepository{}
	signupRepo = &mocks.SignupRepository{}

	repository.SetUserRepository(userRepo)
	repository.SetSignupRepository(signupRepo)
	code := m.Run()
	os.Exit(code)
}
