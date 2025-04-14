package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// TCP-сервер
func tcpServer(protocol, address string, print bool) {
	conn, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		c, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go tcpConnection(c, print)
	}
}

// TCP-соединение
func tcpConnection(c net.Conn, print bool) {
	defer c.Close()
	for {
		msg := make([]byte, 1024)
		n, err := c.Read(msg)
		if err != nil {
			return
		}
		if print {
			fmt.Print(string(msg[:n]))
		}
	}
}

// UDP-сервер
func udpServer(protocol, address string, print bool) {
	conn, err := net.ListenPacket(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		udpConnection(conn, print)
	}
}

// UDP-соединение
func udpConnection(c net.PacketConn, print bool) {
	msg := make([]byte, 1024)
	n, _, err := c.ReadFrom(msg)
	if err != nil {
		return
	}
	if print {
		fmt.Print(string(msg[:n]))
	}
}

func main() {
	fmt.Println(" \n[ СЕТЕВОЕ ЛОГИРОВАНИЕ ]\n ")

	timeout := 30 * time.Second
	flags := log.LstdFlags | log.Lshortfile

	/* Использование */

	// Запуск серверов
	go tcpServer("tcp", "localhost:8080", true)
	go udpServer("udp", "localhost:9090", true)
	time.Sleep(250 * time.Millisecond)

	// Логирование по TCP
	tcp, err := net.DialTimeout("tcp", "localhost:8080", timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer tcp.Close()

	fmt.Println("Протокол TCP:")
	logger := log.New(tcp, "[tcp] ", flags)
	logger.Print("Hello")
	logger.Printf("Cruel\n")
	logger.Println("World")

	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// Логирование по UDP
	udp, err := net.DialTimeout("udp", "localhost:9090", timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer udp.Close()

	fmt.Println("Протокол UDP:")
	logger = log.New(udp, "[udp] ", flags)
	logger.Print("Hello")
	logger.Printf("Cruel\n")
	logger.Println("World")

	time.Sleep(500 * time.Millisecond)

	/* Сравнение скорости */

	times := 100000

	// Запуск серверов
	go tcpServer("tcp", "localhost:8085", false)
	go udpServer("udp", "localhost:9095", false)
	time.Sleep(250 * time.Millisecond)

	fmt.Println()
	fmt.Println("Сравнение скорости")
	fmt.Println("------------------")
	fmt.Println()

	// Логирование по TCP
	tcp, err = net.DialTimeout("tcp", "localhost:8085", timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer tcp.Close()

	fmt.Println("Протокол TCP:")
	logger = log.New(tcp, "[tcp] ", flags)
	start := time.Now()
	for i := 1; i <= times; i++ {
		logger.Print(i)
		if i == 1000 || i == 10000 || i == 100000 || i == 1000000 {
			fmt.Printf("%7d: %v\n", i, time.Since(start))
		}
	}
	fmt.Println()

	// Логирование по UDP
	udp, err = net.DialTimeout("udp", "localhost:9095", timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer udp.Close()

	fmt.Println("Протокол UDP:")
	logger = log.New(udp, "[udp] ", flags)
	start = time.Now()
	for i := 1; i <= times; i++ {
		logger.Print(i)
		if i == 1000 || i == 10000 || i == 100000 || i == 1000000 {
			fmt.Printf("%7d: %v\n", i, time.Since(start))
		}
	}
}
