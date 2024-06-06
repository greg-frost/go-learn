package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(" \n[ НАБЛЮДАТЕЛЬ ]\n ")

	// Каналы
	msg := make(chan int)
	done := make(chan bool)
	until := time.After(5 * time.Second)

	// Отправка и получение
	fmt.Println("Отправка в канал:")
	fmt.Println()
	go send(msg, done)
	recieve(msg, done, until)
}

// Отправка
func send(msg chan<- int, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Отправитель: канал закрыт")
			close(msg) // ... Отправитель закрывает канал
			return
		default:
			msg <- rand.Intn(1_000_000)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Получение
func recieve(msg <-chan int, done chan<- bool, until <-chan time.Time) {
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			fmt.Println()
			fmt.Println("Получатель: прием закончен")
			done <- true // Когда получатель заканчивает прием ...
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}
