package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"go-learn/rest3/internal/user"
	"go-learn/rest3/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println(" \n[ REST 3 (THE ART OF DEVELOPMENT) ]\n ")

	// Получение логгера
	log := logger.New()

	// Создание роутера
	log.Info("Создание роутера")
	router := httprouter.New()

	// Регистрация обработчиков
	log.Info("Регистрация обработчиков")
	handler := user.NewHandler()
	handler.Register(router)

	// Запуск сервера
	startServer(router)
}

// Запуск сервера
func startServer(router *httprouter.Router) {
	log := logger.New()

	log.Info("Запуск сервера")

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
	log.Info("Ожидаю обновлений...")
	log.Info("(на http://localhost:8080)")
	log.Fatal(server.Serve(listener))
}
