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
	five := make([]byte, 5)
	copy(five, msg)
	return five

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

// Функция без утечки среза
func noLeakSlice(slices []Slice) []Slice {
	// Вариант 1
	two := make([]Slice, 2)
	copy(two, slices)
	return two

	// Вариант 2
	// for i := 2; i < len(slices); i++ {
	// 	slices[i].v = nil
	// }
	// return slices[:2]
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
	two := leakSlice(slices)
	runtime.GC()
	runtime.KeepAlive(two)
	printAlloc("Память после")
	fmt.Println()

	// Сборка мусора
	slices = nil
	runtime.GC()

	// Без утечки среза
	fmt.Println("(без утечки)")
	slices = make([]Slice, count)
	printAlloc("Память до")
	for i := 0; i < count; i++ {
		slices[i] = Slice{
			v: make([]byte, size),
		}
	}
	two = noLeakSlice(slices)
	runtime.GC()
	runtime.KeepAlive(two)
	printAlloc("Память после")
}
