package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Счетчик
var counter int

// Инкремент без синхронизации
func Increment() {
	counter++
}

// Мьютекс
var mu sync.Mutex

// Инкремент с мьютексом
func IncrementMutex() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

// Канал
var busy = make(chan int, 1)

// Инкремент через канал
func IncrementChannel() {
	busy <- 1
	counter++
	<-busy
}

// Atomic-счетчик
var acounter atomic.Uint64

// Инкремент через atomic
func IncrementAtomic() {
	acounter.Add(1)
}

func main() {
	fmt.Println(" \n[ ГОНКА ]\n ")

	// Гонка данных
	fmt.Println("Гонка данных")
	times := 1000

	// Без синхронизации
	for i := 0; i < times; i++ {
		go Increment()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Счетчик:  ", counter)

	// Мьютекс
	counter = 0
	for i := 0; i < times; i++ {
		go IncrementMutex()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Мьютекс: ", counter)

	// Канал
	counter = 0
	for i := 0; i < times; i++ {
		go IncrementChannel()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Канал:   ", counter)

	// Atomic
	for i := 0; i < times; i++ {
		go IncrementAtomic()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Atomic:  ", acounter.Load())
	fmt.Println()

	// Состояние гонки
	fmt.Println("Состояние гонки")
	var state string
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()
		state = "День"
	}()
	go func() {
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()
		state = "Ночь"
	}()
	wg.Wait()
	fmt.Println("Значение:", state)
	fmt.Println()

	// Гонка в append
	fmt.Println("Append")
	fmt.Print("Нет гонки: ")
	s := make([]int, 0)
	go func() {
		s1 := append(s, 1)
		fmt.Print(s1, " ")
	}()
	go func() {
		s2 := append(s, 2)
		fmt.Print(s2, " ")
	}()
	time.Sleep(50 * time.Millisecond)
	fmt.Println()
	fmt.Print("Есть гонка: ")
	s = make([]int, 0, 1)
	go func() {
		// c := make([]int, len(s), cap(s))
		// copy(c, s)
		s1 := append(s, 1)
		fmt.Print(s1, " ")
	}()
	go func() {
		// c := make([]int, len(s), cap(s))
		// copy(c, s)
		s2 := append(s, 2)
		fmt.Print(s2, " ")
	}()
	time.Sleep(50 * time.Millisecond)
	fmt.Println()
}
