package util

import (
	"encoding/json"
	"net/http"

	"go-learn/avito/internal/model"
)

// Успешный http-ответ
func Respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Http-ответ с ошибкой
func RespondWithError(w http.ResponseWriter, err model.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}
