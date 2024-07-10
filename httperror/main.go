package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Структура "ошибка"
type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

// JSON-ошибка
func JSONError(w http.ResponseWriter, e Error) {
	data := struct {
		Err Error `json:"error"`
	}{e}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

// Обработчик ошибки
func handleError(w http.ResponseWriter, r *http.Request) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     40300,
		Message:  "Доступ запрещен навсегда",
	}
	JSONError(w, e)
}

func main() {
	fmt.Println(" \n[ HTTP-ОШИБКИ ]\n ")

	// Обработчик
	http.HandleFunc("/", handleError)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
