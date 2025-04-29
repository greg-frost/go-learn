package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"go-learn/avito/internal/model"
	u "go-learn/avito/internal/utils"

	"github.com/dgrijalva/jwt-go"
)

// JWT-аутентификация
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		skipAuth := []string{"/dummyLogin", "/register", "/login"}
		requestPath := r.URL.Path
		for _, path := range skipAuth {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		token, err := parseToken(r.Header.Get("Authorization"))
		if err != nil {
			u.RespondWithError(w, model.Error{
				Code:    http.StatusForbidden,
				Message: err.Error(),
			})
			return
		}

		ctx := u.PutRoleIntoContext(r.Context(), token.Role)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// Парсинг токена
func parseToken(tokenHeader string) (*model.Token, error) {
	if tokenHeader == "" {
		return nil, errors.New("пустой токен авторизации")
	}

	tokenParts := strings.Split(tokenHeader, " ")
	if len(tokenParts) != 2 {
		return nil, errors.New("неверный или поврежденный токен авторизации")
	}

	tk := new(model.Token)
	tokenPart := tokenParts[1]
	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, errors.New("неверный или поврежденный jwt-токен")
	}
	if !token.Valid {
		return nil, errors.New("неверный jwt-токен")
	}

	return tk, nil
}
