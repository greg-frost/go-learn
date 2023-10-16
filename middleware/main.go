package main

import (
	"log"
	"net/http"
)

// Обработчик
func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

// Middleware
func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware")
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/", myMiddleware(myHandler))
	http.ListenAndServe(":8080", nil)
}
