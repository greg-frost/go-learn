package main

import (
	"fmt"
	"log"
	"net/http"
)

// Обработчик
func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("handler")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

// Middleware
func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("middleware")
		next(w, r)
	}
}

func main() {
	fmt.Println(" \n[ MIDDLEWARE ]\n ")

	// Обработчик
	http.HandleFunc("/", myMiddleware(myHandler))

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	fmt.Println()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
