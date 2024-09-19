package auth

import (
	"context"
	"database/sql"
	"math/rand"
	"net/http"
)

// Ключ пользователя в контексте
var userCtxKey = &contextKey{"user"}

// Структура "ключ в контексте"
type contextKey struct {
	name string
}

// Структура "пользователь"
type User struct {
	Name    string
	IsAdmin bool
}

// Промежуточный слой (аутентификация)
func Middleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

			// Пропуск авторизованных пользователей дальше
			// if err != nil || c == nil {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }

			// Получение ID пользователя из cookie
			userId, err := validateAndGetUserID(c)
			if err != nil {
				http.Error(w, "Неправильный cookie", http.StatusForbidden)
				return
			}

			// Получение пользователя из БД
			user := getUserByID(db, userId)

			// Добавление пользователя в контекст
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// Вызов следующего обработчика
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Извлечение пользователя из контекста
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}

// Получение ID пользователя из cookie
func validateAndGetUserID(c *http.Cookie) (int, error) {
	return rand.Intn(2) + 1, nil
}

// Получение пользователя по ID
func getUserByID(db *sql.DB, userId int) *User {
	if userId == 1 {
		return &User{
			Name:    "Greg Frost",
			IsAdmin: true,
		}
	}
	return &User{
		Name: "(гость)",
	}
}
