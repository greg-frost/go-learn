package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	go handleSignals(cancel)

	if err := startServer(ctx); err != nil {
		log.Fatal(err)
	}
}
