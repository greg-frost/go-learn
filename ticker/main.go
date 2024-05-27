package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" \n[ ТИКЕР ]\n ")

	// Тикер и канал завершения
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// Тик
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println(t)
			}
		}
	}()

	// Ожидание и остановка
	time.Sleep(1500 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Тикер остановлен")
}
