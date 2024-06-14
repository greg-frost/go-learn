package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// TCP-сервер
func tcpServer(protocol, address string) {
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
		go tcpConnection(c)
	}
}

// TCP-соединение
func tcpConnection(c net.Conn) {
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

// UDP-сервер
func udpServer(protocol, address string) {
	conn, err := net.ListenPacket(protocol, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		go udpConnection(conn)
	}
}

// UDP-соединение
func udpConnection(c net.PacketConn) {
	msg := make([]byte, 1024)
	n, _, err := c.ReadFrom(msg)
	if err != nil {
		return
	}
	fmt.Print(string(msg[:n]))
}

func main() {
	fmt.Println(" \n[ СЕТЕВОЕ ЛОГИРОВАНИЕ ]\n ")

	// Запуск серверов
	go tcpServer("tcp", "localhost:8080")
	go udpServer("udp", "localhost:9090")
	time.Sleep(100 * time.Millisecond)

	timeout := 30 * time.Second
	flags := log.LstdFlags | log.Lshortfile

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

	time.Sleep(250 * time.Millisecond)
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

	time.Sleep(250 * time.Millisecond)
}
