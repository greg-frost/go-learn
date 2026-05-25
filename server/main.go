package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

// Список заголовков
func headersHandler(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, header)
		}
	}
}

func main() {
	fmt.Println(" \n[ HTTP-СЕРВЕР ]\n ")

	// Роутер
	router := http.NewServeMux()

	// Обработчик
	router.HandleFunc("/", headersHandler)

	// Свой сервер
	addr := "localhost:8080"
	s := &http.Server{
		Addr: addr,
		// Таймаут чтения заголовков запроса
		ReadHeaderTimeout: 500 * time.Millisecond,
		// Таймаут чтения всего запроса
		ReadTimeout: 500 * time.Millisecond,
		// Общий таймаут
		Handler: http.TimeoutHandler(
			router, time.Second, "timeout"),
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Printf("(на http://%s)\n", addr)
	log.Fatal(s.Serve(listener))
}
