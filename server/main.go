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

// TXT-ответ
func textResponse(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, text)
}

// HTML-ответ
func htmlResponse(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, html)
}

// Список заголовков
func headersResponse(res http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(res, "%v: %v\n", name, h)
		}
	}
}

func main() {
	fmt.Println(" \n[ HTTP-СЕРВЕР ]\n ")

	// Обработчики
	http.HandleFunc("/", textResponse)
	http.HandleFunc("/html/", htmlResponse)
	http.HandleFunc("/headers/", headersResponse)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
