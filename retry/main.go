package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Сигнатура функции
type Worker func(context.Context) (string, error)

// Нестабильная функция
func Unstable(threshold int) Worker {
	return func(context.Context) (string, error) {
		if chance := rand.Intn(100); chance >= threshold {
			return "OK", nil
		}
		return "", errors.New("Не работает")
	}
}

// Повтор
func Retry(worker Worker, retries uint, delay time.Duration) Worker {
	return func(ctx context.Context) (string, error) {
		for r := 0; ; r++ {
			// Вызов функции
			res, err := worker(ctx)

			// Если нет ошибок или превышен лимит повторов
			if err == nil || r >= int(retries) {
				return res, err
			}

			fmt.Printf("Попытка №%d не удалась (повтор через %v)\n", r+1, delay)

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return "", ctx.Err()
			}
		}
	}
}

func main() {
	fmt.Println(" \n[ ПОВТОР ]\n ")

	// Настройка
	unstable := Unstable(85)
	worker := Retry(unstable, 5, 2*time.Second)

	// Работа
	res, err := worker(context.Background())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
