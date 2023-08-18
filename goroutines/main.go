package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Текстовая рутина
func routineText(txt string, size, ms int) {
	for i := 1; i <= size; i++ {
		fmt.Print(txt, i, " ")
		waitMs(ms)
	}
	fmt.Println("Поток", txt, "завершен!")
}

// Символьная рутина
func routineVisual(txt string, size, ms int) {
	for i := 1; i <= size; i++ {
		fmt.Print(txt)
		waitMs(ms)
	}
}

// Рутина со случайной паузой
func routineRand(n, ms int) {
	for i := 0; i < n; i++ {
		fmt.Print(n, ":", i, " ")
		waitRandMs(ms)
	}
}

// Пауза в мс
func waitMs(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

// Пауза в сек
func waitSec(sec int) {
	waitMs(sec * 1000)
}

// Случайная пауза в мс
func waitRandMs(ms int) {
	waitMs(rand.Intn(ms))
}

func main() {
	fmt.Println(" \n[ ГОРУТИНЫ ]\n ")

	/* Текстовые рутины */

	go routineText("A", 100, 1)
	go routineText("B", 50, 10)
	go routineText("C", 10, 100)

	waitSec(2)

	fmt.Println()

	/* Символьные рутины */

	go routineVisual("-", 100, 1)
	go routineVisual("+", 50, 10)
	go routineVisual("#", 10, 100)

	waitSec(2)

	fmt.Println(" \n ")

	/* Рутины со случайной паузой */

	for i := 0; i < 10; i++ {
		go routineRand(i, 250)
	}

	/* Ожидание ввода */

	var input string
	fmt.Scanln(&input)
}
