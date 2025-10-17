package main

import (
	"fmt"
	"log"
	"runtime"
)

// Печать статистики памяти
func printMemoryStats(m runtime.MemStats) {
	runtime.ReadMemStats(&m)
	fmt.Println("Alloc:", m.Alloc)
	fmt.Println("HeapAlloc:", m.HeapAlloc)
	fmt.Println("TotalAlloc:", m.TotalAlloc)
	fmt.Println("NumGC:", m.NumGC)
}

func main() {
	fmt.Println(" \n[ СБОРЩИК МУСОРА ]\n ")

	var m runtime.MemStats

	// Память не выделена
	fmt.Println("Начальное состояние")
	printMemoryStats(m)
	fmt.Println()

	// Выделено 500 МБ
	fmt.Println("Выделение памяти")
	for i := 0; i < 10; i++ {
		s := make([]byte, 50_000_000)
		if s == nil {
			log.Fatal("Ошибка выделения памяти!")
		}
	}
	printMemoryStats(m)
	fmt.Println()

	// Выделено 1 ГБ
	fmt.Println("Еще выделение памяти")
	for i := 0; i < 10; i++ {
		s := make([]byte, 100_000_000)
		if s == nil {
			log.Fatal("Ошибка выделения памяти!")
		}
	}
	printMemoryStats(m)
}
