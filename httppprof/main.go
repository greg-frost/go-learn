package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"time"
)

// Обработчик
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL:  %s\n", r.URL.Path)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
}

// Обработчик времени
func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Time: %s\n", t)
	fmt.Fprintf(w, "URL:  %s\n", r.URL.Path)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
}

func main() {
	fmt.Println(" \n[ HTTP-ПРОФИЛИРОВАНИЕ ]\n ")

	// Роутер
	r := http.NewServeMux()

	// Обработчики
	r.HandleFunc("/", handler)
	r.HandleFunc("/time/", timeHandler)

	// Профилирование
	r.HandleFunc("/debug/pprof", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
