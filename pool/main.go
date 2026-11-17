package main

import (
	"fmt"
	"strings"
	"sync"
)

// Число воркеров
const workersCount = 3

// Структура "задача"
type Job struct {
	id   int
	word string
}

// Структура "результат"
type Result struct {
	job    Job
	worker int
	res    string
}

// Обработка задач
func Process(worker int, jobs chan Job, ress chan Result) {
	for j := range jobs {
		ress <- Result{
			job:    j,
			worker: worker,
			res:    strings.ToUpper(j.word),
		}
	}
}

func main() {
	fmt.Println(" \n[ ПУЛ ВОРКЕРОВ ]\n ")

	// Исходные данные
	words := []string{"Hello", "My", "Darling", "Seniorita", "Goodbye", "Cruel", "World"}
	fmt.Println(words)
	fmt.Println()

	// Каналы для задач и результатов
	jobs := make(chan Job, workersCount)
	ress := make(chan Result, workersCount)

	// Постановка задач в очередь
	go func() {
		for i, w := range words {
			jobs <- Job{
				id:   i,
				word: w,
			}
		}
		fmt.Println("Задачи добавлены")
		close(jobs)
	}()

	// Запуск воркеров для обработки задач
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Process(i+1, jobs, ress)
		}(i)
	}
	fmt.Println("Воркеры запущены")

	// Ожидание завершения Воркеров
	go func() {
		wg.Wait()
		fmt.Println("Воркеры завершены")
		close(ress)
	}()

	// Получение результатов
	res := make([]string, len(words))
	for r := range ress {
		res[r.job.id] = r.res
	}
	fmt.Println("Результаты получены")
	fmt.Println()
	fmt.Println(res)
}
