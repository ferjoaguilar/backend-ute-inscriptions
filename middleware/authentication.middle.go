package middleware

import (
	"net/http"
	"strings"

	"github.com/ferjoaguilar/backend-utec-inscriptions/models"
	"github.com/ferjoaguilar/backend-utec-inscriptions/server"
	"github.com/ferjoaguilar/backend-utec-inscriptions/utils"
	"github.com/golang-jwt/jwt/v4"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
	}
)

func checkAuth(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func AuthenticationMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !checkAuth(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})

			if err != nil {
				utils.ResponseWriter(w, http.StatusUnauthorized, "Not authorized", err.Error())
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
