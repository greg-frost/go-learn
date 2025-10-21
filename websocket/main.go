package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// Upgrader для WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HTTP-обработчик
func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Используйте /ws для WebSocket")
}

// WS-обработчик
func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Соединение:", r.Host)

	// Получение соединения WS
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка upgrader.Upgrade:", err)
		return
	}
	defer ws.Close()

	// Ожидание соединений
	for {
		// Получение сообщения
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Ошибка чтения сообщения с хоста %v: %v", r.Host, err)
			break
		}
		log.Println("Получено:", string(message))

		// Отправка сообщения
		message = []byte(strings.ToUpper(string(message)))
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("Ошибка записи сообщения:", err)
			break
		}
		log.Println("Отправлено:", string(message))
	}
}

// Сервер
func server() {
	// Роутер
	mux := http.NewServeMux()

	// Параметры сервера
	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	// Обработчики
	mux.HandleFunc("/", httpHandler)
	mux.HandleFunc("/ws", wsHandler)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(s.ListenAndServe())
}

func main() {
	fmt.Println(" \n[ WEBSOCKET ]\n ")

	server()
}
