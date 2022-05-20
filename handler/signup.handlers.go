package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/snowball-devs/backend-utec-inscriptions/models"
	"github.com/snowball-devs/backend-utec-inscriptions/repository"
	"github.com/snowball-devs/backend-utec-inscriptions/server"
	"github.com/snowball-devs/backend-utec-inscriptions/utils"
)

type signupNewRequest struct {
	Name      string `json:"name" validate:"required,min=8,max=75"`
	Lastname  string `json:"lastname" validate:"required,min=8,max=75"`
	Age       int    `json:"age" validate:"required,min=15,max=60"`
	Dni       string `json:"dni" validate:"required,min=10,max=10"`
	Nit       string `json:"nit" validate:"required,min=17,max=17"`
	Country   string `json:"country" validate:"required,min=5,max=50"`
	City      string `json:"city" validate:"required,min=5,max=50"`
	Address   string `json:"address" validate:"required,min=8,max=100"`
	Cellphone string `json:"cellphone" validate:"required,min=8,max=12"`
	Graduated string `json:"graduated" validate:"required,min=8,max=100"`
	User      string `json:"user"`
}

type signupNewResponse struct {
	Message string `json:"message"`
}

type getSignups struct {
	Signups []models.Signup `json:"signups"`
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
			User:      userId.Hex(),
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
		pageStr := r.URL.Query().Get("page")

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Failed getting query", err.Error())
			return
		}

		response, err := repository.GetSignups(r.Context(), page)

		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Failed get signups", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getSignups{Signups: response})
	}
}

func CompleteSignupHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		signupId := params["signupId"]
		response, err := repository.CompleteSignup(r.Context(), signupId)
		if err != nil {
			utils.ResponseWriter(w, http.StatusInternalServerError, "Error to complete signup", err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(disabledResponse{
			Message: response,
		})
	}
}
