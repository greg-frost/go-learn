package main

import (
	"fmt"
	"sync"
	"time"
)

// Выполнение между (единственное)
func Between(wait, done chan struct{}, msg string) {
	<-wait
	fmt.Println(msg)
	time.Sleep(250 * time.Millisecond)
	close(done)
}

// Выполнение после (многократное)
func After(wait chan struct{}, wg *sync.WaitGroup, msg string) {
	<-wait
	fmt.Println(msg)
	time.Sleep(250 * time.Millisecond)
	wg.Done()
}

func main() {
	fmt.Println(" \n[ ПОРЯДОК ВЫПОЛНЕНИЯ ]\n ")

	// Этапы работы
	var wg sync.WaitGroup
	a := make(chan struct{})
	b := make(chan struct{})
	c := make(chan struct{})
	d := make(chan struct{})

	// После третьего этапа
	wg.Add(1)
	go After(d, &wg, "Логирование")

	// Первый этап
	go Between(a, b, "Инициализация")

	// После третьего этапа
	wg.Add(1)
	go After(d, &wg, "Очистка")

	// Второй этап
	go Between(b, c, "Сбор данных")

	// Третий этап
	go Between(c, d, "Выполнение")

	// После третьего этапа
	wg.Add(1)
	go After(d, &wg, "Завершение")

	close(a) // Запуск первого этапа
	wg.Wait()
}
