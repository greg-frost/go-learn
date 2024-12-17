package main

import (
	"fmt"
	"log"
	"net/http"

	"go-learn/rest2/app"
	"go-learn/rest2/controllers"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(" \n[ REST 2 (TPROGER) ]\n ")

	// Роутер
	router := mux.NewRouter()

	// Аутентификация
	router.Use(app.JwtAuthentication)

	// Обработчики
	router.HandleFunc("/api/user/new",
		controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login",
		controllers.Authenticate).Methods("POST")

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
