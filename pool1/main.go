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
func Process(worker int, jobs chan Job, ress chan Result, wg *sync.WaitGroup) {
	for j := range jobs {
		ress <- Result{
			job:    j,
			worker: worker,
			res:    strings.ToUpper(j.word),
		}
	}
	wg.Done()
}

func main() {
	fmt.Println(" \n[ ПУЛ 1 ]\n ")

	// Исходные данные
	words := []string{"Hello", "Cruel", "World", "Goodbye", "My", "Darling", "Seniorita"}
	n := len(words)
	fmt.Println(words)
	fmt.Println()

	// Каналы для задач, результатов и завершения
	jobs := make(chan Job, workersCount)
	ress := make(chan Result, workersCount)
	done := make(chan bool)

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

	// Получение результатов
	res := make([]string, n)
	go func() {
		for r := range ress {
			res[r.job.id] = r.res
		}
		fmt.Println("Результаты получены")
		done <- true
	}()

	// Запуск воркеров для обработки задач
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go Process(i+1, jobs, ress, &wg)
	}
	fmt.Println("Воркеры запущены")
	wg.Wait()
	fmt.Println("Воркеры завершены")
	close(ress)

	<-done
	fmt.Println()
	fmt.Println(res)
}
