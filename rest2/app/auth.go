package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"go-learn/rest2/models"
	u "go-learn/rest2/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWT-аутентификация
var jwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		skipAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path
		for _, path := range skipAuth {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = u.Message(false, "Missing header token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 {
			response = u.Message(false, "Invalid or malformed header token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		tokenPart := tokenParts[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			response = u.Message(false, "Malformed jwt-token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Invalid jwt-token")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		log.Println("UserID:", tk.UserID)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
