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
func increment() {
	counter++
}

// Мьютекс
var mtx sync.Mutex

// Инкремент с мьютексом
func mutexIncrement() {
	mtx.Lock()
	defer mtx.Unlock()
	counter++
}

// Канал
var busy = make(chan int, 1)

// Инкремент через канал
func channelIncrement() {
	busy <- 1
	counter++
	<-busy
}

// Atomic-счетчик
var acounter atomic.Uint64

// Инкремент через atomic
func atomicIncrement() {
	acounter.Add(1)
}

func main() {
	fmt.Println(" \n[ ГОНКА ДАННЫХ ]\n ")

	times := 1000

	// Без синхронизации
	for i := 0; i < times; i++ {
		go increment()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Гонка:  ", counter)

	// Мьютекс
	counter = 0
	for i := 0; i < times; i++ {
		go mutexIncrement()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Mutex:  ", counter)

	// Atomic
	for i := 0; i < times; i++ {
		go atomicIncrement()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Atomic: ", acounter.Load())

	// Канал
	counter = 0
	for i := 0; i < times; i++ {
		go channelIncrement()
	}
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Channel:", counter)
}
