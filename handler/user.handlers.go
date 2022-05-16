package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"github.com/snowball-devs/backend-utec-inscriptions/repository"
	"github.com/snowball-devs/backend-utec-inscriptions/server"
	"github.com/snowball-devs/backend-utec-inscriptions/utils"
	"golang.org/x/crypto/bcrypt"
)

type signupRequest struct {
	Email       string `json:"email,omitempty" validate:"required,email"`
	Username    string `json:"username,omitempty" validate:"required,min=3,max=25"`
	Password    string `json:"password,omitempty" validate:"required,min=5,max=75"`
	Permissions string `json:"permissions,omitempty" validate:"required,oneof=manager student"`
}

type loginRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=5,max=75"`
}

type loginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type disabledResponse struct {
	Message string `json:"message"`
}

func SignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = signupRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to parse json information", err.Error())
			return
		}

		var validate *validator.Validate = validator.New()
		err = validate.Struct(&request)
		if err != nil {
			utils.ResponseWriter(w, http.StatusBadRequest, "Error to validation", err.Error())
			return
		}

		passHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to hash password", err.Error())
			return
		}

		var newUser models.User = models.User{
			Email:       request.Email,
			Username:    request.Username,
			Password:    string(passHash),
			Permissions: request.Permissions,
		}

		response, err := repository.CreateUser(r.Context(), &newUser)

		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to create new user", err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = loginRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to parse json information", err.Error())
			return
		}

		var validate *validator.Validate = validator.New()
		err = validate.Struct(&request)
		if err != nil {
			utils.ResponseWriter(w, http.StatusBadRequest, "Error to validation", err.Error())
			return
		}

		user, err := repository.FindUserByEmail(r.Context(), request.Email)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to login user", err.Error())
			return
		}

		if user == nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Email or password is incorrect", nil)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			utils.ResponseWriter(w, http.StatusUnauthorized, "Email or password is incorrect", nil)
			return
		}

		claims := models.AppClaims{
			UserId:   user.ID,
			Email:    user.Email,
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to generate session", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(loginResponse{
			Username: user.Username,
			Token:    tokenString,
		})
	}
}

func DisabledUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userId := params["userId"]

		response, err := repository.DisableUser(r.Context(), userId)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to disable user", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(disabledResponse{
			Message: response,
		})

	}
}

func GetManagersHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := repository.GetManagers(r.Context())
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to get managers", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
