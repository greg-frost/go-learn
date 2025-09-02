package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Сигнатура функции
type Worker func(string) (string, error)

// Сигнатура функции с контекстом
type WithContext func(context.Context, string) (string, error)

// Медленная функция
func Slow(max time.Duration) Worker {
	return func(arg string) (string, error) {
		time.Sleep(time.Duration(rand.Intn(int(max))))
		return arg, nil
	}
}

// Таймаут
func Timeout(worker Worker) WithContext {
	return func(ctx context.Context, arg string) (string, error) {
		chRes := make(chan string)
		chErr := make(chan error)

		go func() {
			// Вызов функции
			res, err := worker(arg)
			chRes <- res
			chErr <- err
		}()

		select {
		case res := <-chRes:
			return res, <-chErr
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

func main() {
	fmt.Println(" \n[ Таймаут ]\n ")

	// Настройка
	ctx := context.Background()
	ctxt, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Работа
	slow := Slow(2 * time.Second)
	timeout := Timeout(slow)
	for i := 0; i < 3; i++ {
		res, err := timeout(ctxt, "Go!")
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println(res)
		}
	}
}
