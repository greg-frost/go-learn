package main

import (
	"fmt"
	"sync"
	"time"
)

// Демультиплексор
func Split(source <-chan string, n int) []<-chan string {
	dests := make([]<-chan string, 0)

	// Создание каналов-приемников
	for i := 1; i <= n; i++ {
		ch := make(chan string)
		dests = append(dests, ch)

		go func(i int) {
			defer close(ch)

			// Обход значений канала-источника
			for n := range source {
				// Запись в канал-приемник
				ch <- fmt.Sprintf("#%v <- %v", i, n)
			}
		}(i)
	}

	return dests
}

func main() {
	fmt.Println(" \n[ ДЕМУЛЬТИПЛЕКСОР ]\n ")

	// Подготовка
	source := make(chan string)
	dests := Split(source, 5)

	// Запись в канал-источник
	go func() {
		for i := 1; i <= 15; i++ {
			source <- fmt.Sprint(i)
			time.Sleep(100 * time.Millisecond)
		}

		close(source)
	}()

	var wg sync.WaitGroup
	wg.Add(len(dests))

	// Чтение из каналов-приемников
	for _, ch := range dests {
		go func(d <-chan string) {
			defer wg.Done()

			// Обход значений канала
			for val := range d {
				fmt.Println(val)
			}
		}(ch)
	}

	wg.Wait()
}
