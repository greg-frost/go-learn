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

	log.Print("Создание роутера")
	router := httprouter.New()

	log.Print("Регистрация обработчиков")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Print("Запуск приложения")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Ожидаю обновлений...")
	log.Print("(на http://localhost:8080)")
	log.Fatal(server.Serve(listener))
}
