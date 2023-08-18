package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Структура "Сервер"
type Server struct{}

// Смена знака
func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

// Удвоение
func (this *Server) Double(i int64, reply *int64) error {
	*reply = i * 2
	return nil
}

// Сервер
func server() {
	rpc.Register(new(Server))

	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Ошибка инициализации сервера:", err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Ошибка приема соединения:", err)
			continue
		}

		go rpc.ServeConn(c)
	}
}

// Клиент
func client() {
	c, err := rpc.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Ошибка инициализации клиента:", err)
		return
	}

	var input, result int64
	input = 128

	/* Смена знака */

	err = c.Call("Server.Negate", input, &result)
	if err != nil {
		fmt.Println("Ошибка выполнения RPC:", err)
	} else {
		fmt.Print("Server.Negate(", input, ") = ", result)
	}

	fmt.Println()

	/* Удвоение */

	err = c.Call("Server.Double", input, &result)
	if err != nil {
		fmt.Println("Ошибка выполнения RPC:", err)
	} else {
		fmt.Print("Server.Double(", input, ") = ", result)
	}

	c.Close()
}

func main() {
	fmt.Println(" \n[ RPC ]\n ")

	/* Запуск клиента и сервера */

	go server()
	go client()

	/* Ожидание ввода */

	var input string
	fmt.Scanln(&input)
}
