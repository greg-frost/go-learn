package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"go-learn/rest3/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println(" \n[ REST 3 (THE ART OF DEVELOPMENT) ]\n ")

	// Создание роутера
	log.Print("Создание роутера")
	router := httprouter.New()

	// Регистрация обработчиков
	log.Print("Регистрация обработчиков")
	handler := user.NewHandler()
	handler.Register(router)

	// Запуск сервера
	startServer(router)
}

// Запуск сервера
func startServer(router *httprouter.Router) {
	log.Print("Запуск сервера")

	// Адрес и порт
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// Настройка
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Запуск
	log.Print("Ожидаю обновлений...")
	log.Print("(на http://localhost:8080)")
	log.Fatal(server.Serve(listener))
}
