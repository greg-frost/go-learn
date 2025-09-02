package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Сигнатура функции
type Worker func(context.Context) (string, error)

// Функция-счетчик
func Counter() Worker {
	var count int
	return func(context.Context) (string, error) {
		count++
		return fmt.Sprintf("Запрос №%d", count), nil
	}
}

// Дроссельная заслонка
func Throttle(worker Worker, max uint, refill uint, period time.Duration) Worker {
	var tokens = max
	var once sync.Once

	return func(ctx context.Context) (string, error) {
		// Отмена
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		once.Do(func() {
			// Запуск таймера
			ticker := time.NewTicker(period)

			go func() {
				// Остановка таймера
				defer ticker.Stop()

				for {
					select {
					// Отмена
					case <-ctx.Done():
						return

					// Восполнение токенов
					case <-ticker.C:
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})

		// Ограничение
		if tokens <= 0 {
			return "", errors.New("Слишком много запросов")
		}
		tokens--

		// Вызов функции
		return worker(ctx)
	}
}

func main() {
	fmt.Println(" \n[ ДРОССЕЛЬНАЯ ЗАСЛОНКА ]\n ")

	// Настройка
	counter := Counter()
	worker := Throttle(counter, 5, 3, time.Second)

	// Работа
	for i := 0; i < 15; i++ {
		res, err := worker(context.Background())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
