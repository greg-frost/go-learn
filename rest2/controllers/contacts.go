package controllers

import (
	"encoding/json"
	"net/http"

	"go-learn/rest2/models"
	u "go-learn/rest2/utils"
)

// Создание контакта
var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	if err := json.NewDecoder(r.Body).Decode(contact); err != nil {
		u.Respond(w, u.Message(false, "Неправильный запрос"))
		return
	}

	contact.UserID = user
	resp := contact.Create()
	u.Respond(w, resp)
}

// Получение контактов
var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	data := models.GetContacts(user)
	resp := u.Message(true, "Контакты получены")
	resp["data"] = data
	u.Respond(w, resp)
}
