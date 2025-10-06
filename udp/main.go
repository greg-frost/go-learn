package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// Сервер
func server() {
	// Прослушивание UDP
	conn, err := net.ListenPacket("udp", ":9999")
	if err != nil {
		fmt.Println("Ошибка инициализации сервера:", err)
		return
	}
	defer conn.Close()

	// Ожидание соединений
	for {
		connection(conn)
		// go connection(conn)
	}
}

// Соединение
func connection(c net.PacketConn) {
	// Чтение сообщения
	b := make([]byte, 1024)
	n, _, err := c.ReadFrom(b)
	if err != nil {
		fmt.Println("Ошибка чтения сообщения:", err)
		return
	}

	// Декодирование сообщения
	var msg string
	err = json.Unmarshal(b[:n], &msg)
	if err != nil {
		fmt.Println("Ошибка декодирования сообщения:", err)
	} else {
		fmt.Println("Получено:", msg)
	}
}

// Клиент
func client() {
	// Соединение по UDP
	c, err := net.Dial("udp", "localhost:9999")
	if err != nil {
		fmt.Println("Ошибка инициализации клиента:", err)
		return
	}
	defer c.Close()

	msg := "Привет, сервер!"
	fmt.Println("Отправлено:", msg)

	// Кодирование и отправка сообщения
	err = json.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("Ошибка кодирования сообщения:", err)
	}
}

func main() {
	fmt.Println(" \n[ UDP ]\n ")

	// Локальный сервер
	go server()
	go client()

	// Ожидание
	time.Sleep(time.Second)
}
