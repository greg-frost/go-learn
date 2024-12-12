package models

import (
	"strings"

	u "go-learn/rest2/utils"

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

// Валидация
func (a *Account) Validate() (map[string]interface{}, bool) {
	if strings.Contains(a.Email, "@") {
		return u.Message(false, "Email-адрес необходим"), false
	}

	if len(a.Password) < 6 {
		return u.Message(false, "Пароль необходим"), false
	}

	account := &Account{}
	err := DB().Table("accounts").Where("email=?", a.Email).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Ошибка обращения к БД"), false
	}
	if account.Email != "" {
		return u.Message(false, "Email-адрес уже занят"), false
	}

	return u.Message(true, "OK"), true
}

// Получение пользователя
func User(userID uint) *Account {
	account := &Account{}

	DB().Table("accounts").Where("id=?", userID).First(account)
	if account.Email == "" {
		return nil
	}
	account.Password = ""

	return account
}
