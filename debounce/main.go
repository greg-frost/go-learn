package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Сигнатура функции
type Worker func(context.Context) (string, error)

// Функция-счетчик
func Counter() Worker {
	count := 0
	return func(context.Context) (string, error) {
		count++
		return fmt.Sprintf("Запрос №%d", count), nil
	}
}

// Антидребезг (первый запрос)
func DebounceFirst(worker Worker, d time.Duration) Worker {
	var threshold time.Time
	var res string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()

		// Увеличение периода
		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		// Если запрос попадает в период
		if time.Now().Before(threshold) {
			return res, err
		}

		// Вызов функции
		res, err = worker(ctx)

		return res, err
	}
}

// Антидребезг (последний запрос)
func DebounceLast(worker Worker, period time.Duration) Worker {
	var threshold time.Time = time.Now()
	var ticker *time.Ticker
	var res string
	var err error
	var once sync.Once
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer m.Unlock()

		// Увеличение периода
		threshold = time.Now().Add(period)

		once.Do(func() {
			// Запуск таймера
			ticker = time.NewTicker(10 * time.Millisecond)

			go func() {
				// Остановка таймера
				defer func() {
					m.Lock()
					ticker.Stop()
					once = sync.Once{}
					m.Unlock()
				}()

				for {
					select {
					// Получение последнего вызова
					case <-ticker.C:
						m.Lock()
						if time.Now().After(threshold) {
							// Вызов функции
							res, err = worker(ctx)
							m.Unlock()
							return
						}
						m.Unlock()

					// Отмена
					case <-ctx.Done():
						m.Lock()
						res, err = "", ctx.Err()
						m.Unlock()
						return
					}
				}
			}()
		})

		return res, err
	}
}

func main() {
	fmt.Println(" \n[ АНТИДРЕБЕЗГ ]\n ")

	// Отброс всех запросов, кроме первого
	counter := Counter()
	fmt.Println("DebounceFirst:")
	for i := 0; i < 5; i++ {
		worker := DebounceFirst(counter, 50*time.Millisecond)
		fmt.Printf("\nПачка №%d\n--------\n", i+1)

		for j := 0; j < 10; j++ {
			go func() {
				time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
				res, _ := worker(context.Background())
				fmt.Println(res)
			}()
		}
		time.Sleep(200 * time.Millisecond)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// Отброс всех запросов, кроме последнего
	counter = Counter()
	fmt.Println("DebounceLast:")
	for i := 0; i < 5; i++ {
		worker := DebounceLast(counter, 50*time.Millisecond)
		fmt.Printf("\nПачка №%d\n--------\n", i+1)

		for j := 0; j < 10; j++ {
			go func() {
				time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
				res, _ := worker(context.Background())
				if res != "" {
					fmt.Println(res)
				}
			}()
		}
		time.Sleep(200 * time.Millisecond)
	}
	time.Sleep(500 * time.Millisecond)
}
