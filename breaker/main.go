package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
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

// Размыкатель цепи
func Breaker(worker Worker, threshold uint) Worker {
	var failures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock()

		// Число оставшихся попыток
		f := failures - int(threshold)

		// Сервис недоступен

		if f >= 0 {
			delay := 2 << f * time.Second
			shouldRetryAt := lastAttempt.Add(delay)

			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", fmt.Errorf("Сервис недоступен (повтор через %v)", delay)
			}
		}

		m.RUnlock()

		// Вызов функции
		res, err := worker(ctx)

		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()

		// Ошибка
		if err != nil {
			failures++
			return res, err
		}

		// Сброс
		failures = 0

		return res, nil
	}
}

func main() {
	fmt.Println(" \n[ РАЗМЫКАТЕЛЬ ЦЕПИ ]\n ")

	// Настройка
	unstable := Unstable(85)
	worker := Breaker(unstable, 5)

	// Работа
	for {
		res, err := worker(context.Background())

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			break
		}

		time.Sleep(time.Second)
	}
}
