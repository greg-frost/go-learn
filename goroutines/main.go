package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Текстовая рутина
func TextRoutine(msg string, size, ms int) {
	for i := 1; i <= size; i++ {
		fmt.Print(msg, i, " ")
		waitMs(ms)
	}
	fmt.Println("Поток", msg, "завершен!")
}

// Символьная рутина
func SymbolRoutine(msg string, size, ms int) {
	for i := 1; i <= size; i++ {
		fmt.Print(msg)
		waitMs(ms)
	}
}

// Рутина со случайной паузой
func RandomRoutine(n, ms int) {
	for i := 0; i < n; i++ {
		fmt.Print(n, ":", i, " ")
		waitRandMs(ms)
	}
}

// Пауза в мс
func waitMs(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
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

	// Текстовые рутины
	go TextRoutine("A", 100, 1)
	go TextRoutine("B", 50, 10)
	go TextRoutine("C", 10, 100)
	waitSec(2)
	fmt.Println()

	// Символьные рутины
	go SymbolRoutine("-", 100, 1)
	go SymbolRoutine("+", 50, 10)
	go SymbolRoutine("#", 10, 100)
	waitSec(2)
	fmt.Println()
	fmt.Println()

	// Рутины со случайной паузой
	for i := 0; i < 10; i++ {
		go RandomRoutine(i, 250)
	}

	// Ожидание ввода
	var input string
	fmt.Scanln(&input)
}
