package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Отправка
func Send(msg chan<- int, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Отправитель: канал закрыт")
			close(msg) // ... отправитель закрывает канал
			return
		default:
			msg <- rand.Intn(1e6)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Получение
func Recieve(msg <-chan int, done chan<- bool, until <-chan time.Time) {
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			fmt.Println()
			fmt.Println("Получатель: прием закончен")
			done <- true // Когда получатель заканчивает прием, ...
			return
		}
	}
}

func main() {
	fmt.Println(" \n[ ЗАКРЫТИЕ ]\n ")

	// Каналы
	msg := make(chan int)
	done := make(chan bool)
	until := time.After(5 * time.Second)

	// Отправка и получение
	fmt.Println("Отправка в канал:")
	fmt.Println()
	go Send(msg, done)
	Recieve(msg, done, until)
}
