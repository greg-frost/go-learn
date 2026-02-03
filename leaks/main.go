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

// Функция с утечкой емкости
func leakCap(msg []byte) []byte {
	return msg[:5]
}

// Функция без утечки емкости
func noLeakCap(msg []byte) []byte {
	// Вариант 1
	header := make([]byte, 5)
	copy(header, msg)
	return header

	// Вариант 2
	// return msg[:5:5]
}

// Структура "срезы"
type Slice struct {
	v []byte
}

// Функция с утечкой среза
func leakSlice(slices []Slice) []Slice {
	return slices[:2]
}

func main() {
	fmt.Println(" \n[ УТЕЧКИ ]\n ")

	count := 1_000
	size := 100_000

	// Утечка емкости
	fmt.Println("Утечка емкости:")
	printAlloc("Память до")
	var caps int
	for i := 0; i < count; i++ {
		msg := make([]byte, size)
		five := leakCap(msg)
		caps += cap(five)
	}
	printAlloc("Память после")
	fmt.Println("Общая емкость:", caps)
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Без утечки емкости
	fmt.Println("(без утечки)")
	caps = 0
	printAlloc("Память до")
	for i := 0; i < count; i++ {
		msg := make([]byte, size)
		five := noLeakCap(msg)
		caps += cap(five)
	}
	printAlloc("Память после")
	fmt.Println("Общая емкость:", caps)
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Утечка среза
	fmt.Println("Утечка среза:")
	slices := make([]Slice, count)
	printAlloc("Память до")
	for i := 0; i < count; i++ {
		slices[i] = Slice{
			v: make([]byte, size),
		}
	}
	printAlloc("Память после")
	two := leakSlice(slices)
	runtime.GC() // Сборка мусора
	printAlloc("Память в итоге")
	runtime.KeepAlive(two)
}
