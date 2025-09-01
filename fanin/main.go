package main

import (
	"fmt"
	"sync"
	"time"
)

// Мультиплексор
func Funnel(sources ...<-chan string) <-chan string {
	dest := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(sources))

	// Обход каналов-источников
	for _, ch := range sources {
		go func(c <-chan string) {
			defer wg.Done()

			// Обход значений канала
			for val := range c {
				// Запись в канал-приемник
				dest <- val
			}
		}(ch)
	}

	// Ожидание и закрытие
	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest
}

func main() {
	fmt.Println(" \n[ МУЛЬТИПЛЕКСОР ]\n ")

	// Подготовка
	sources := make([]<-chan string, 0)

	// Запись в каналы-источники
	for i := 1; i <= 3; i++ {
		ch := make(chan string)
		sources = append(sources, ch)

		go func(i int) {
			defer close(ch)

			for j := 1; j <= 5; j++ {
				ch <- fmt.Sprintf("#%v -> %v", i, j)
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	// Чтение из канала-приемника
	dest := Funnel(sources...)
	for d := range dest {
		fmt.Println(d)
	}
}
