package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	fmt.Println("[СЕРВЕР] Соединение:", r.Host)

	// Получение соединения WS
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ошибка upgrader.Upgrade:", err)
		return
	}
	defer ws.Close()

	// Ожидание соединений
	for {
		// Получение сообщения
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("[СЕРВЕР] Ошибка получения сообщения:", err)
			break
		}
		fmt.Println("[СЕРВЕР] Получено:", string(msg))

		// Отправка сообщения
		msg = []byte(strings.ToUpper(string(msg)))
		err = ws.WriteMessage(mt, msg)
		if err != nil {
			fmt.Println("[СЕРВЕР] Ошибка отправки сообщения:", err)
			break
		}
		fmt.Println("[СЕРВЕР] Отправлено:", string(msg))
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
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	fmt.Println()
	log.Fatal(s.ListenAndServe())
}

// Клиент
func client() {
	// Соединение по WebSocket
	URL := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "ws"}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		fmt.Println("Ошибка инициализации клиента:", err)
		return
	}
	defer c.Close()

	// Получение сообщений
	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				fmt.Println("[КЛИЕНТ] Ошибка получения сообщения:", err)
				return
			}
			fmt.Println("[КЛИЕНТ] Получено:", string(msg))
		}
	}()

	// Отправка сообщения
	msg := "Привет, сервер!"
	err = c.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Println("[КЛИЕНТ] Ошибка отправки сообщения:", err)
		return
	}
	fmt.Println("[КЛИЕНТ] Отправлено:", msg)

	// Закрытие соединения
	err = c.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		fmt.Println("[КЛИЕНТ] Ошибка закрытия соединения:", err)
		return
	}
	fmt.Println("[КЛИЕНТ] Закрытие соединения")

	time.Sleep(100 * time.Millisecond)
}

func main() {
	fmt.Println(" \n[ WEBSOCKET ]\n ")

	// Локальный сервер
	go server()
	go client()

	// Ожидание
	time.Sleep(time.Second)
}
