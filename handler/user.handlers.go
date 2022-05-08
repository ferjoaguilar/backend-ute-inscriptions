package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"github.com/snowball-devs/backend-utec-inscriptions/repository"
	"github.com/snowball-devs/backend-utec-inscriptions/server"
	"github.com/snowball-devs/backend-utec-inscriptions/utils"
	"golang.org/x/crypto/bcrypt"
)

type signupRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Username string `json:"username,omitempty" validate:"required,min=3,max=25"`
	Password string `json:"password,omitempty" validate:"required,min=5,max=75"`
}

func SignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = signupRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to parse json information", err)
		}

		var validate *validator.Validate = validator.New()
		err = validate.Struct(&request)
		if err != nil {
			utils.ResponseWriter(w, http.StatusBadRequest, "Error to validation", err)
		}

		passHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to hash password", err)
		}

		var newUser models.User = models.User{
			Email:    request.Email,
			Username: request.Username,
			Password: string(passHash),
		}

		response, err := repository.CreateUser(r.Context(), &newUser)

		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to create new user", err)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
