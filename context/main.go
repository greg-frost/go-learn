package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

// Запуск сервера
func startServer(ctx context.Context) error {
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		return err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	defer l.Close()

	log.Print("Сервер запущен")

	for {
		select {
		case <-ctx.Done():
			log.Print("Сервер остановлен")
			return nil
		default:
			log.Print("Ожидание подключений")

			if err := l.SetDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}

			_, err := l.Accept()
			if err != nil {
				if os.IsTimeout(err) {
					continue
				}

				return err
			}

			log.Print("Новое подключение")
		}
	}
}

// Обработка сигналов
func handleSignals(cancel context.CancelFunc) {
	chSig := make(chan os.Signal)
	signal.Notify(chSig, os.Interrupt)
	for {
		sig := <-chSig
		switch sig {
		case os.Interrupt:
			cancel()
			return
		}
	}
}

func main() {
	fmt.Println(" \n[ КОНТЕКСТ ]\n ")

	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на localhost:8080)")
	fmt.Println("или отмены (Ctrl+C)")

	fmt.Println()

	ctx, cancel := context.WithCancel(context.Background())
	go handleSignals(cancel)

	if err := startServer(ctx); err != nil {
		log.Fatal(err)
	}
}
