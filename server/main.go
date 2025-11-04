package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	text = "Go, Baby, Go!"
	html = `
		<doctype html>
		<html>
			<head>
				<title>Go Server</title>
			</head>
			<body>
				<h1>Go Server</h1>
				<p>Go, Baby, Go!</p>
			</body>
		</html>
	`
)

// TXT-обработчик
func textHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, text)
}

// HTML-обработчик
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, html)
}

// Список заголовков
func headersHandler(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	fmt.Println(" \n[ HTTP-СЕРВЕР ]\n ")

	// Обработчики
	http.HandleFunc("/", textHandler)
	http.HandleFunc("/html/", htmlHandler)
	http.HandleFunc("/headers/", headersHandler)

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
