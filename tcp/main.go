package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// Сервер
func server() {
	conn, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Ошибка инициализации сервера:", err)
		return
	}
	defer conn.Close()

	for {
		c, err := conn.Accept()
		if err != nil {
			fmt.Println("Ошибка приема соединения:", err)
			continue
		}

		go connection(c)
	}
}

// Соединение
func connection(c net.Conn) {
	var msg string

	err := json.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println("Ошибка декодирования сообщения:", err)
	} else {
		fmt.Println("Получено:", msg)
	}

	c.Close()
}

// Клиент
func client() {
	c, err := net.Dial("tcp", "localhost:9999")
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
	fmt.Println(" \n[ TCP ]\n ")

	/* Сторонний сервер */

	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(status)

	/* Локальный сервер */

	go server()
	go client()

	// Ожидание
	time.Sleep(1 * time.Second)
}
