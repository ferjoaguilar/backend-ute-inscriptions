package repository_test

import (
	"os"
	"testing"

	"github.com/snowball-devs/backend-utec-inscriptions/repository"
	"github.com/snowball-devs/backend-utec-inscriptions/repository/mocks"
)

var userRepo *mocks.UserRepository

// Init mock instance to unit testing
func TestMain(m *testing.M) {
	userRepo = &mocks.UserRepository{}
	repository.SetUserRepository(userRepo)
	code := m.Run()
	os.Exit(code)
}
