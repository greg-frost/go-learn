package main

import (
	"fmt"
	"strings"
	"sync"
)

// Число воркеров
const workersCount = 3

// Размер буфера
const bufferSize = 5

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
	fmt.Println(" \n[ ПУЛ 2 ]\n ")

	// Исходные данные
	words := []string{"Hello", "My", "Darling", "Seniorita", "Goodbye", "Cruel", "World"}
	n := len(words)
	fmt.Println(words)
	fmt.Println()

	// Каналы для задач и результатов
	jobs := make(chan Job, bufferSize)
	ress := make(chan Result, bufferSize)

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
			Process(i+1, jobs, ress)
			wg.Done()
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
	res := make([]string, n)
	for r := range ress {
		res[r.job.id] = r.res
	}
	fmt.Println("Результаты получены")
	fmt.Println()
	fmt.Println(res)
}
