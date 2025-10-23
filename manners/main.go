package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

// Структура "обработчик"
type handler struct{}

// Получение нового обработчика
func newHandler() *handler {
	return &handler{}
}

// Обработка HTTP-запроса
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Мир"
	}
	fmt.Fprintf(w, "Привет, %s!", name)
}

// Ожидание завершения
func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	fmt.Print("Завершение работы...")
	manners.Close()
}

func main() {
	fmt.Println(" \n[ MANNERS ]\n ")

	// Обработчик
	handler := newHandler()

	// Планирование мягкого завершения
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	manners.ListenAndServe("localhost:8080", handler)
}
