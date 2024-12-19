package controllers

import (
	"encoding/json"
	"net/http"

	"go-learn/rest2/models"
	u "go-learn/rest2/utils"
)

// Создание аккаунта
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	if err := json.NewDecoder(r.Body).Decode(account); err != nil {
		u.Respond(w, u.Message(false, "Неправильный запрос"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}

// Аутентификация
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	if err := json.NewDecoder(r.Body).Decode(account); err != nil {
		u.Respond(w, u.Message(false, "Неправильный запрос"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
