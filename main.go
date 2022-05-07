package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/snowball-devs/backend-utec-inscriptions/middleware"
	"github.com/snowball-devs/backend-utec-inscriptions/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error to loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWTSecret")
	DATABASE_URL := os.Getenv("DatabaseUrl")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(bindRoutes)

}

func bindRoutes(s server.Server, r *mux.Router) {

	r.Use(middleware.GlobalApplicationJson(s))

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/", homeHandler(s)).Methods(http.MethodGet)
}

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func homeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to API REST GOLANG",
			Status:  true,
		})
	}
}
