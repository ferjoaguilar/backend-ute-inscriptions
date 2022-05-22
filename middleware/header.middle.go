package middleware

import (
	"net/http"

	"github.com/ferjoaguilar/backend-utec-inscriptions/server"
)

func GlobalApplicationJson(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}
