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
	// Прослушивание TCP
	l, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Ошибка инициализации сервера:", err)
		return
	}
	defer l.Close()

	// Ожидание соединений
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Ошибка приема соединения:", err)
			continue
		}

		// Обработка запроса
		go connection(c)
	}
}

// Соединение
func connection(c net.Conn) {
	// Декодирование сообщения
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
	// Соединение по TCP
	c, err := net.Dial("tcp", "localhost:9999")
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
	fmt.Println(" \n[ TCP ]\n ")

	// Сторонний сервер
	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(status)

	// Локальный сервер
	go server()
	go client()

	// Ожидание
	time.Sleep(time.Second)
}
