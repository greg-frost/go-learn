package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
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
	Token    string `json:"token" sql:"-"`
}
