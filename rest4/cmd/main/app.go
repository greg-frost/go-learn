package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"go-learn/rest4/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println(" \n[ REST 4 (ART DEVELOPMENT) ]\n ")

	log.Print("create router")
	router := httprouter.New()

	log.Print("register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Print("start application")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("listen http://localhost:8080")
	log.Fatal(server.Serve(listener))
}
