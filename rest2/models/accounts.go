package models

import (
	"os"
	"strings"

	u "go-learn/rest2/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

// Валидация аккаунта
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

	return u.Message(true, "Валидация аккаунта пройдена"), true
}

// Создание аккаунта
func (a *Account) Create() map[string]interface{} {
	if resp, ok := a.Validate(); !ok {
		return resp
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	a.Password = string(passwordHash)

	DB().Create(a)
	if a.ID <= 0 {
		return u.Message(false, "Не удалось создать аккаунт")
	}

	tk := &Token{UserID: a.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	a.Token = tokenString
	a.Password = ""

	resp := u.Message(true, "Аккаунт создан")
	resp["account"] = a
	return resp
}

// Вход в аккаунт
func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := DB().Table("accounts").Where("email=?", email).First(account).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return u.Message(false, "Email-адрес не найден")
		}
		return u.Message(false, "Ошибка обращения к БД")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Неправильные email или пароль")
	}
	account.Password = ""

	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := u.Message(true, "Вход выполнен")
	resp["account"] = account
	return resp
}

// Получение пользователя
// func GetUser(id uint) *Account {
// 	account := &Account{}
// 	DB().Table("accounts").Where("id=?", id).First(account)
// 	if account.Email == "" {
// 		return nil
// 	}
// 	account.Password = ""

// 	return account
// }
