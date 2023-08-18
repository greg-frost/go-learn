package main

import (
	"fmt"
	"sync"
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

func main() {
	fmt.Println(" \n[ СОСТОЯНИЕ ГОНКИ ]\n ")

	times := 1000

	/* Без синхронизации */

	for i := 0; i < times; i++ {
		go increment()
	}

	time.Sleep(time.Millisecond * 50)

	fmt.Println("Есть гонка: ", counter)

	/* Мьютекс */

	counter = 0

	for i := 0; i < times; i++ {
		go mutexIncrement()
	}

	time.Sleep(time.Millisecond * 50)

	fmt.Println("С мьютексом:", counter)

	/* Канал */

	counter = 0

	for i := 0; i < times; i++ {
		go channelIncrement()
	}

	time.Sleep(time.Millisecond * 50)

	fmt.Println("Через канал:", counter)
}
