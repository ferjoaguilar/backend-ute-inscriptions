package handler

import (
	"encoding/json"
	"net/http"

	"github.com/snowball-devs/backend-utec-inscriptions/server"
)

type notFound struct {
	Url     string `json:"url"`
	Message string `json:"message"`
}

func NotFoundHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(notFound{
			Url:     r.URL.Path,
			Message: "Enpoint not found",
		})
	}
}
