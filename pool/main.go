package main

import (
	"fmt"
	"strings"
	"sync"
)

// Число воркеров
const workers = 3

// Структура "задача"
type Job struct {
	id   int
	word string
}

// Структура "результат"
type Res struct {
	job    Job
	worker int
	res    string
}

// Обработка задач
func process(worker int, jobs chan Job, done chan Res) {
	for j := range jobs {
		done <- Res{
			job:    j,
			worker: worker,
			res:    strings.ToUpper(j.word),
		}
	}
}

func main() {
	fmt.Println(" \n[ ПУЛ ]\n ")

	// Исходные данные
	words := []string{"Hello", "Cruel", "World", "Goodbye", "My", "Darling", "Seniorita"}
	n := len(words)
	fmt.Println(words)
	fmt.Println()

	// Каналы для задач и результатов
	jobs := make(chan Job, n)
	done := make(chan Res, n)

	// Постановка задач в очередь
	for i, w := range words {
		jobs <- Job{
			id:   i,
			word: w,
		}
	}
	fmt.Println("Задачи добавлены")
	close(jobs)

	// Запуск воркеров для обработки задач
	for i := 0; i < workers; i++ {
		go process(i+1, jobs, done)
	}
	fmt.Println("Обработчики запущены")

	// Получение результатов
	res := make([]string, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			r := <-done
			res[r.job.id] = r.res
		}()
	}
	wg.Wait()
	fmt.Println("Результаты получены")
	fmt.Println()

	fmt.Println(res)
}
