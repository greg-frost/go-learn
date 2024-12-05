package utils

import (
	"encoding/json"
	"net/http"
)

// Сообщение
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Ответ
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "applications/json")
	json.NewEncoder(w).Encode(data)
}
