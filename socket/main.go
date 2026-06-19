package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"go-learn/base"
)

// Путь и сокет
var path = base.Dir("socket")
var socket = filepath.Join(path, "unix.socket")

// Сервер
func Server() {
	// Проверка существования файла сокета
	if _, err := os.Stat(socket); os.IsExist(err) {
		fmt.Println("Удаление старого файла сокета")
		if err := os.Remove(socket); err != nil {
			fmt.Println("Ошибка удаления файла сокета:", err)
			return
		}
	}

	// Прослушивание сокета
	l, err := net.Listen("unix", socket)
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
func Client() {
	// Соединение по сокету
	c, err := net.Dial("unix", socket)
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
	fmt.Println(" \n[ СОКЕТ ]\n ")

	// Локальный сервер
	go Server()
	go Client()

	// Ожидание
	time.Sleep(time.Second)
}
