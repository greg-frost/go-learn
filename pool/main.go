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
func process(worker int, jobs chan Job, done chan Res, wg *sync.WaitGroup) {
	for j := range jobs {
		done <- Res{
			job:    j,
			worker: worker,
			res:    strings.ToUpper(j.word),
		}
	}
	wg.Done()
}

func main() {
	fmt.Println(" \n[ ПУЛ ]\n ")

	// Исходные данные
	words := []string{"Hello", "Cruel", "World", "Goodbye", "My", "Darling", "Seniorita"}
	n := len(words)
	fmt.Println(words)
	fmt.Println()

	// Каналы для задач, результатов и завершения
	jobs := make(chan Job, workers)
	ress := make(chan Res, workers)
	done := make(chan bool)

	// Постановка задач в очередь
	go func() {
		for i, w := range words {
			jobs <- Job{
				id:   i,
				word: w,
			}
		}
		close(jobs)
		fmt.Println("Задачи добавлены")
	}()

	// Получение результатов
	res := make([]string, n)
	go func() {
		for r := range ress {
			res[r.job.id] = r.res
		}
		done <- true
		fmt.Println("Результаты получены")
	}()

	// Запуск воркеров для обработки задач
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go process(i+1, jobs, ress, &wg)
	}
	wg.Wait()
	fmt.Println("Обработчики запущены")
	close(ress)

	<-done
	fmt.Println()
	fmt.Println(res)
}
