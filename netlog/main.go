package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Сервер
func server(protocol, address string) {
	ln, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go connection(c)
	}
}

// Соединение
func connection(c net.Conn) {
	defer c.Close()
	for {
		msg := make([]byte, 1024)
		n, err := c.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			return
		}
		fmt.Print(string(msg[:n]))
	}
}

func main() {
	fmt.Println(" \n[ СЕТЕВОЕ ЛОГИРОВАНИЕ ]\n ")

	// Запуск сервера
	go server("tcp", "localhost:8080")
	time.Sleep(100 * time.Millisecond)

	timeout := 30 * time.Second
	flags := log.LstdFlags | log.Lshortfile

	// Логирование по TCP
	conn, err := net.DialTimeout("tcp", "localhost:8080", timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Протокол TCP:")
	logger := log.New(conn, "[tcp] ", flags)
	logger.Print("Hello")
	logger.Printf("Cruel\n")
	logger.Println("World")

	time.Sleep(250 * time.Millisecond)
}
