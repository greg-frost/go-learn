package main

import (
	"fmt"
	"io"
	"net/http"
)

// TXT-ответ
func textResponse(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Go, Baby, Go!")
}

// HTML-ответ
func htmlResponse(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`
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
		`,
	)
}

func main() {
	fmt.Println(" \n[ HTTP ]\n ")

	/* Настройка обработчиков */

	http.HandleFunc("/", textResponse)
	http.HandleFunc("/html/", htmlResponse)

	http.Handle(
		"/files/",
		http.StripPrefix(
			"/files/",
			http.FileServer(http.Dir("export")),
		),
	)

	/* Запуск сервера */

	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на localhost:8080)")

	http.ListenAndServe("localhost:8080", nil)
}
