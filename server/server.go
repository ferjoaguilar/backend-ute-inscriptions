package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/ferjoaguilar/backend-utec-inscriptions/database"
	"github.com/ferjoaguilar/backend-utec-inscriptions/repository"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type broker struct {
	config *Config
	router mux.Router
}

func (b *broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*broker, error) {
	if config.Port == "" {
		return nil, errors.New("POST is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWTSecret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("DatabaseUrl is required")
	}

	broker := &broker{
		config: config,
		router: *mux.NewRouter(),
	}

	return broker, nil
}

func (b *broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = *mux.NewRouter()
	binder(b, &b.router)

	// Database
	repo, err := database.NewMongoRepository(b.config.DatabaseUrl)

	repository.SetUserRepository(repo)
	repository.SetSignupRepository(repo)

	// Cors settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://utec-inscriptions.feraguilar.tech", "http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
	})

	handler := c.Handler(&b.router)

	// INIT REST SERVER
	log.Println("Start server on port", b.config.Port)
	err = http.ListenAndServe(b.config.Port, handler)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
