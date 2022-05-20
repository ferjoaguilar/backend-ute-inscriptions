package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/snowball-devs/backend-utec-inscriptions/handler"
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
	r.Use(middleware.AuthenticationMiddleware(s))

	//Not found endpoint
	r.NotFoundHandler = handler.NotFoundHandler(s)

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/users/signup", handler.SignupHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/users/login", handler.LoginHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/users/{userId}", handler.DisabledUserHandler(s)).Methods(http.MethodDelete)
	api.HandleFunc("/users/managers", handler.GetManagersHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/users/me", handler.MetHandler(s)).Methods(http.MethodGet)

	api.HandleFunc("/inscriptions", handler.CreateSignup(s)).Methods(http.MethodPost)
	api.HandleFunc("/inscriptions", handler.GetSignupsHandler(s)).Methods(http.MethodGet)
}
