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
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Пропуск некоторых эндпоинтов
		skipAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path
		for _, path := range skipAuth {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Запрос и заголовок токена
		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		// Пустой заголовок токена
		if tokenHeader == "" {
			response = u.Message(false, "Отсутствует заголовок-токен")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		// Получение нужной части заголовка токена
		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 {
			response = u.Message(false, "Неправильный или поврежденный заголовок-токен")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		// Структура токена
		tokenPart := tokenParts[1]
		tk := &models.Token{}

		// Парсинг JWT-токена
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil {
			response = u.Message(false, "Поврежденный jwt-токен")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		// Неправильный токен
		if !token.Valid {
			response = u.Message(false, "Неправильный jwt-токен")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		// Логирование, оборачивание контекста и пропуск далее
		log.Println("Аутентификация (UserID):", tk.UserID)
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
