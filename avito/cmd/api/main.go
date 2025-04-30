package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"go-learn/avito/internal/auth"
	"go-learn/avito/internal/db/postgres"
	"go-learn/avito/internal/handler"

	"github.com/gorilla/mux"
)

func init() {
	os.Setenv("DB_NAME", "learn")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASS", "admin")
}

func main() {
	fmt.Println(" \n[ AVITO-СТАЖИРОВКА ]\n ")

	addr := flag.String("addr", "localhost", "адрес сервера")
	port := flag.Int("port", 8080, "порт сервера")
	flag.Parse()

	router := mux.NewRouter()

	pgParams := postgres.ConnectionParams{
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
	pgStorage, err := postgres.NewStorage(pgParams)
	if err != nil {
		log.Fatal(err)
	}

	muxHandler := handler.NewHandler(pgStorage)
	muxHandler.Register(router)

	router.Use(auth.JwtAuthentication)

	startServer(router, fmt.Sprintf("%s:%d", *addr, *port))
}

// Запуск сервера
func startServer(router *mux.Router, connAddr string) {
	listener, err := net.Listen("tcp", connAddr)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	fmt.Println("Ожидаю подключений...")
	fmt.Println("(на http://" + connAddr + ")")
	log.Fatal(server.Serve(listener))
}
