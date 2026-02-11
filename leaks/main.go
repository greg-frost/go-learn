package main

import (
	"fmt"
	"runtime"
	"strings"
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

// Тип "карта значений"
type vMap map[int][128]byte

// Тип "карта указателей"
// (уменьшит размер утечки)
type pMap map[int]*[128]byte

// Функция с утечкой карты
func leakMap(m vMap) vMap {
	size := len(m)
	for i := 1000; i < size; i++ {
		delete(m, i)
	}
	return m
}

// Функция без утечки карты
func noLeakMap(m vMap) vMap {
	size := len(m)
	for i := 1000; i < size; i++ {
		delete(m, i)
	}

	nm := make(vMap, len(m))
	for k, v := range m {
		nm[k] = v
	}
	return nm
}

// Функция с утечкой строки
func leakString(s string) string {
	return s[:100]
}

// Функция без утечки строки
func noLeakString(s string) string {
	// Вариант 1
	return strings.Clone(s[:100])

	// Вариант 2
	// return string([]byte(s[:100]))
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
	m := make(vMap)
	for i := 0; i < size; i++ {
		m[i] = [128]byte{}
	}
	m = leakMap(m)
	runtime.GC()
	runtime.KeepAlive(m)
	printAlloc("Память после")
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Без утечки карты
	fmt.Println("(без утечки)")
	printAlloc("Память до")
	m = make(vMap)
	for i := 0; i < size; i++ {
		m[i] = [128]byte{}
	}
	nm := noLeakMap(m)
	runtime.GC()
	runtime.KeepAlive(nm)
	printAlloc("Память после")
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Утечка строки
	fmt.Println("Утечка строки:")
	printAlloc("Память до")
	s := strings.Repeat("#", count*size)
	ns := leakString(s)
	runtime.GC()
	runtime.KeepAlive(ns)
	printAlloc("Память после")
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Без утечки строки
	fmt.Println("(без утечки)")
	printAlloc("Память до")
	s = strings.Repeat("#", count*size)
	ns = noLeakString(s)
	runtime.GC()
	runtime.KeepAlive(ns)
	printAlloc("Память после")
}
