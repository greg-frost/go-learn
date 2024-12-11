package main

import (
	"fmt"
	"log"
	"net/http"

	"go-learn/rest2/app"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(" \n[ REST 2 (TPROGER) ]\n ")

	// Роутер
	router := mux.NewRouter()

	// Аутентификация
	router.Use(app.JwtAuthentication)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
