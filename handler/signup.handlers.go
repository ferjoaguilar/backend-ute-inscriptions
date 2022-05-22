package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ferjoaguilar/backend-utec-inscriptions/models"
	"github.com/ferjoaguilar/backend-utec-inscriptions/repository"
	"github.com/ferjoaguilar/backend-utec-inscriptions/server"
	"github.com/ferjoaguilar/backend-utec-inscriptions/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type signupNewRequest struct {
	Name      string             `json:"name" validate:"required,min=8,max=75"`
	Lastname  string             `json:"lastname" validate:"required,min=8,max=75"`
	Age       int                `json:"age" validate:"required,min=15,max=60"`
	Dni       string             `json:"dni" validate:"required,min=10,max=10"`
	Nit       string             `json:"nit" validate:"required,min=17,max=17"`
	Country   string             `json:"country" validate:"required,min=5,max=50"`
	City      string             `json:"city" validate:"required,min=5,max=50"`
	Address   string             `json:"address" validate:"required,min=8,max=100"`
	Cellphone string             `json:"cellphone" validate:"required,min=8,max=12"`
	Graduated string             `json:"graduated" validate:"required,min=8,max=100"`
	User      primitive.ObjectID `json:"user"`
}

type signupNewResponse struct {
	Message string `json:"message"`
}

type getSignups struct {
	Signups []models.SignupLookup `json:"signups"`
}

type changeStatusResponse struct {
	Message string `json:"message"`
}

type getSignupResponse struct {
	Signup *models.Signup `json:"signup"`
}

func CreateSignup(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = signupNewRequest{}

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

		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

		userId, err := utils.GetTokenInformation(s, tokenString)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Failed to getting token information", err.Error())
			return
		}

		var newSignup models.Signup = models.Signup{
			Name:      request.Name,
			Lastname:  request.Lastname,
			Age:       request.Age,
			Dni:       request.Dni,
			Nit:       request.Nit,
			Country:   request.Country,
			City:      request.City,
			Address:   request.Address,
			Cellphone: request.Cellphone,
			Graduated: request.Graduated,
			User:      userId,
		}
		response, err := repository.CreateSignup(r.Context(), &newSignup)

		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Failed to create inscription", err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(signupNewResponse{
			Message: response,
		})
	}
}

func GetSignupsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusStr := r.URL.Query().Get("status")
		if statusStr == "" {
			utils.ResponseWriter(w, http.StatusBadRequest, "This endpoint required query params call status", nil)
			return
		}

		response, err := repository.GetSignups(r.Context(), statusStr)

		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Failed get signups", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getSignups{Signups: response})
	}
}

func ChangeStatusHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userId := params["userId"]

		statusStr := r.URL.Query().Get("status")
		if statusStr == "" {
			utils.ResponseWriter(w, http.StatusBadRequest, "This endpoint required query params call status", nil)
			return
		}

		response, err := repository.ChangeStatus(r.Context(), userId, statusStr)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Failed to change status", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(changeStatusResponse{
			Message: response,
		})
	}
}

func GetSignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userId := params["userId"]
		response, err := repository.GetSignupById(r.Context(), userId)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Failed get signup", err.Error())
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getSignupResponse{
			Signup: response,
		})
	}
}
