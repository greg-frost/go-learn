package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

// Число воркеров
var workers = 3

// Семафор
var sem = semaphore.NewWeighted(int64(workers))

// Обработка значения
func process(n int) int {
	time.Sleep(250 * time.Millisecond)
	return n * n
}

func main() {
	fmt.Println(" \n[ СЕМАФОР ]\n ")

	jobs := 10
	res := make([]int, jobs)
	ctx := context.Background()

	fmt.Printf("Выполнение %d задач в %d потоках:\n", jobs, workers)
	for i := range res {
		// Получение семафора
		err := sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("не удалось получить семафор:", err)
			break
		}

		// Выполнение работы
		go func(i int) {
			defer sem.Release(1)
			res[i] = process(i)
		}(i)
	}

	// Блокировка до завершения (аналог wg.Wait)
	err := sem.Acquire(ctx, int64(workers))
	if err != nil {
		log.Fatal("не удалось получить семафор:", err)
	}

	// Результаты
	for k, v := range res {
		fmt.Println(k, "->", v)
	}
}
