package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Текст и HTML-разметка
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

// Клиент
func Client(addr string) (int, string) {
	time.Sleep(50 * time.Millisecond)

	get, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer get.Body.Close()
	fmt.Printf("GET: %s --> %s\n", addr, get.Status)

	body, err := io.ReadAll(get.Body)
	if err != nil {
		log.Fatal(err)
	}

	return get.StatusCode, string(body)
}

func main() {
	fmt.Println(" \n[ ТЕСТИРОВАНИЕ HTTP ]\n ")

	// Обработчики
	http.HandleFunc("/", textHandler)
	http.HandleFunc("/html/", htmlHandler)

	// Запуск клиентов
	go Client("http://localhost:8080")
	go Client("http://localhost:8080/html")

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	fmt.Println()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
