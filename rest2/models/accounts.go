package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// Структура "токен"
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// Структура "аккаунт"
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Структура "контакт"
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
}
