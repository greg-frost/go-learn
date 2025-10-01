package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Каналы для чтения и записи
var read = make(chan int)
var write = make(chan int)

// Запись значения
func setValue(value int) {
	write <- value
}

// Чтение значения
func getValue() int {
	return <-read
}

// Конвейер
func pipeline() {
	var value int
	for {
		select {
		case newValue := <-write:
			value = newValue
			fmt.Printf("%d ", value)
		case read <- value:
		}
	}
}

func main() {
	fmt.Println(" \n[ КОНВЕЙЕР ]\n ")

	n := 10
	fmt.Printf("Генерация %d случайных чисел:\n", n)
	var wg sync.WaitGroup
	go pipeline()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			setValue(rand.Intn(10 * n))
		}()
	}
	wg.Wait()
	fmt.Printf("\nПоследнее значение: %d\n", getValue())
}
