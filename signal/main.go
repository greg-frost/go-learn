package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println(" \n[ СИГНАЛЫ ]\n ")

	// Каналы сигналов и завершения
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Прослушивание определенных сигналов
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Перехват сигнала
	go func() {
		sig := <-sigs
		fmt.Println("Получен сигнал:", sig)
		done <- true
	}()

	// Ожидание прерывания
	fmt.Println("Ожидание сигнала...")
	fmt.Println("(нажмите Ctrl+C)")

	<-done

	fmt.Println("Завершение работы")
}
