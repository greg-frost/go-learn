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

// Структура "срез"
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

// Тип "карта"
type Map map[int]bool

// Функция с утечкой карты
func leakMap(m Map) Map {
	for i := 1000; i < len(m); i++ {
		delete(m, i)
	}
	return m
}

func main() {
	fmt.Println(" \n[ УТЕЧКИ ]\n ")

	count := 1_000
	size := 100_000

	// Утечка емкости
	fmt.Println("Утечка емкости:")
	printAlloc("Память до")
	msg := make([]byte, count*size)
	five := leakCap(msg)
	runtime.GC()
	runtime.KeepAlive(five)
	printAlloc("Память после")
	fmt.Println("Емкость:", cap(five))
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Без утечки емкости
	fmt.Println("(без утечки)")
	printAlloc("Память до")
	msg = make([]byte, size)
	five = noLeakCap(msg)
	runtime.GC()
	runtime.KeepAlive(five)
	printAlloc("Память после")
	fmt.Println("Емкость:", cap(five))
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Утечка среза
	fmt.Println("Утечка среза:")
	printAlloc("Память до")
	slices := make([]Slice, count)
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
	runtime.GC()

	// Без утечки среза
	fmt.Println("(без утечки)")
	printAlloc("Память до")
	slices = make([]Slice, count)
	for i := 0; i < count; i++ {
		slices[i] = Slice{
			v: make([]byte, size),
		}
	}
	two = noLeakSlice(slices)
	runtime.GC()
	runtime.KeepAlive(two)
	printAlloc("Память после")
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Утечка карты
	fmt.Println("Утечка карты:")
	printAlloc("Память до")
	m := make(Map)
	for i := 0; i < count*count; i++ {
		m[i] = true
	}
	m = leakMap(m)
	runtime.GC()
	runtime.KeepAlive(m)
	printAlloc("Память после")
}
