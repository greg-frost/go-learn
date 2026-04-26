package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

// Эхо-функция
func Echo(from io.Reader, to io.Writer) {
	for {
		data := make([]byte, 1024)
		n, _ := from.Read(data)
		data = bytes.ToUpper(data[:n])
		to.Write(data)
	}
}

// Чтение источника
func Read(pipe chan<- []byte) {
	for {
		data := make([]byte, 1024)
		if n, _ := os.Stdin.Read(data); n > 0 {
			pipe <- data[:n]
		}
	}
}

// Запись в источник
func Write(pipe <-chan []byte, done <-chan time.Time) {
	for {
		select {
		case data := <-pipe:
			os.Stdout.Write(bytes.ToUpper(data))
		case <-done:
			fmt.Println("Время истекло...")
			return
		}
	}
}

func main() {
	fmt.Println(" \n[ ЭХО-СЕРВИС ]\n ")

	timeout := 10 * time.Second
	fmt.Println("Вводите строки, пока не истечет время:")

	// Без каналов
	// go Echo(os.Stdin, os.Stdout)
	// time.Sleep(timeout)
	// fmt.Println("Время истекло...")

	// С каналами
	done := time.After(timeout)
	pipe := make(chan []byte)
	go Read(pipe)
	Write(pipe, done)
}
