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
	fmt.Println(" \n[ ПОСРЕДНИК ]\n ")

	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на localhost:8080)")

	fmt.Println()

	http.HandleFunc("/", myMiddleware(myHandler))
	http.ListenAndServe(":8080", nil)
}
