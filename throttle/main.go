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
func Throttle(worker Worker, max, refill uint, period time.Duration) Worker {
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

// Мульти-дроссельная заслонка
type MultiThrottled func(context.Context, string) (bool, string, error)

// Структура "бакет"
type bucket struct {
	tokens uint
	time   time.Time
}

func MultiThrottle(worker Worker, max, refill uint, d time.Duration) MultiThrottled {
	// Свой лимит для каждого пользователя
	buckets := make(map[string]*bucket)

	return func(ctx context.Context, uid string) (bool, string, error) {
		b := buckets[uid]

		// Новый бакет
		if b == nil {
			buckets[uid] = &bucket{
				tokens: max - 1,
				time:   time.Now(),
			}

			// Вызов функции
			res, err := worker(ctx)
			return true, res, err
		}

		// Количество токенов, которые можно добавить,
		// учитывая время, прошедшее с последнего запроса
		refillInterval := uint(time.Since(b.time) / d)
		newTokens := refill * refillInterval
		currentTokens := b.tokens + newTokens

		// Недостаточно токенов
		if currentTokens < 1 {
			return false, "", errors.New("Слишком много запросов")
		}

		// Если бакет пополнился, запоминается текущее время
		if currentTokens > max {
			b.time = time.Now()
			b.tokens = max - 1
		} else { // Иначе - время последнего добавления жетонов
			deltaTokens := currentTokens - b.tokens
			deltaRefill := deltaTokens / refill
			deltaTime := time.Duration(deltaRefill) * d

			b.time = b.time.Add(deltaTime)
			b.tokens = currentTokens - 1
		}

		// Вызов функции
		res, err := worker(ctx)
		return true, res, err
	}
}

func main() {
	fmt.Println(" \n[ ДРОССЕЛЬНАЯ ЗАСЛОНКА ]\n ")

	// Дроссельная заслонка
	fmt.Println("Таймеры:")
	counter := Counter()
	timers := Throttle(counter, 5, 3, time.Second)

	// Работа
	for i := 0; i < 15; i++ {
		res, err := timers(context.Background())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()

	// Мульти-дроссельная заслонка
	fmt.Println("Интервалы (мульти):")
	counter = Counter()
	intervals := MultiThrottle(counter, 5, 3, time.Second)

	// Работа
	for i := 0; i < 15; i++ {
		ok, res, err := intervals(context.Background(), "user")
		if !ok {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
