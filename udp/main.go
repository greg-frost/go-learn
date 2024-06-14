package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// Сервер
func server() {
	conn, err := net.ListenPacket("udp", ":9999")
	if err != nil {
		fmt.Println("Ошибка инициализации сервера:", err)
		return
	}
	defer conn.Close()

	for {
		go connection(conn)
	}
}

// Соединение
func connection(c net.PacketConn) {
	b := make([]byte, 1024)
	n, _, err := c.ReadFrom(b)
	if err != nil {
		fmt.Println("Ошибка чтения сообщения:", err)
		return
	}

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
	c, err := net.Dial("udp", "localhost:9999")
	if err != nil {
		fmt.Println("Ошибка инициализации клиента:", err)
		return
	}

	msg := "Привет, сервер!"
	fmt.Println("Отправлено:", msg)

	err = json.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("Ошибка кодирования сообщения:", err)
	}

	c.Close()
}

func main() {
	fmt.Println(" \n[ UDP ]\n ")

	/* Локальный сервер */

	go server()
	go client()

	// Ожидание
	time.Sleep(1 * time.Second)
}
