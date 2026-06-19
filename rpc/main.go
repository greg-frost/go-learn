package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// Структура "сервер"
type server struct{}

// Смена знака
func (*server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

// Удвоение
func (*server) Double(i int64, reply *int64) error {
	*reply = i * 2
	return nil
}

// Сервер
func Server() {
	// Регистрация RPC-сервера
	rpc.Register(new(server))

	// Прослушивание TCP
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Ошибка инициализации сервера:", err)
		return
	}

	// Ожидание соединений
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Ошибка приема соединения:", err)
			continue
		}

		// Обработка запроса
		go rpc.ServeConn(c)
	}
}

// Клиент
func Client() {
	// Инициализация RPC-клиента
	c, err := rpc.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Ошибка инициализации клиента:", err)
		return
	}
	defer c.Close()

	var input int64 = 128
	var result int64

	// Смена знака
	err = c.Call("Server.Negate", input, &result)
	if err != nil {
		fmt.Println("Ошибка выполнения RPC:", err)
	} else {
		fmt.Printf("Server.Negate(%d) = %d\n", input, result)
	}

	// Удвоение
	err = c.Call("Server.Double", input, &result)
	if err != nil {
		fmt.Println("Ошибка выполнения RPC:", err)
	} else {
		fmt.Printf("Server.Double(%d) = %d\n", input, result)
	}
}

func main() {
	fmt.Println(" \n[ RPC ]\n ")

	fmt.Println("Server:")

	// Запуск клиента и сервера
	go Server()
	go Client()

	// Ожидание
	time.Sleep(time.Second)
}
