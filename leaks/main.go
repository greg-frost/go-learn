package main

import (
	"fmt"
	"runtime"
)

// Печать потребления памяти
func printAlloc(caption string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: %d KB\n", caption, m.Alloc/1024)
}

// Функция с утечкой памяти
func memoryLeak(msg []byte) []byte {
	return msg[:5]
}

// Функция с утечкой памяти
func memoryOk(msg []byte) []byte {
	// Вариант 1
	header := make([]byte, 5)
	copy(header, msg)
	return header

	// Вариант 2
	// return msg[:5:5]
}

func main() {
	fmt.Println(" \n[ УТЕЧКИ ]\n ")

	// Утечка памяти
	fmt.Println("Утечка памяти:")
	var caps int
	count := 1_000
	size := 1_000_000
	printAlloc("Память до")
	for i := 0; i < count; i++ {
		msg := make([]byte, size)
		header := memoryLeak(msg)
		caps += cap(header)
	}
	printAlloc("Память после")
	fmt.Println("Общая емкость:", caps)
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Без утечки памяти
	fmt.Println("Без утечки:")
	caps = 0
	printAlloc("Память до")
	for i := 0; i < count; i++ {
		msg := make([]byte, size)
		header := memoryOk(msg)
		caps += cap(header)
	}
	printAlloc("Память после")
	fmt.Println("Общая емкость:", caps)
}
