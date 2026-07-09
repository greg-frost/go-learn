package model

import "testing"

// Тестирование пользователя
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    "user@example.com",
		Password: "password",
	}
}
