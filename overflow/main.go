package main

import (
	"fmt"
	"math"
)

// Инкремент без переполнения (int)
func IncInt(counter int) int {
	if counter == math.MaxInt {
		panic("переполнение int")
	}
	return counter + 1
}

// Инкремент без переполнения (uint)
func IncUint(counter uint) uint {
	if counter == math.MaxUint {
		panic("переполнение uint")
	}
	return counter + 1
}

// Инкремент без переполнения (int32)
func IncInt32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("переполнение int32")
	}
	return counter + 1
}

// Инкремент без переполнения (int64)
func IncInt64(counter int64) int64 {
	if counter == math.MaxInt64 {
		panic("переполнение int64")
	}
	return counter + 1
}

// Сложение без переполнения
func AddInt(a, b int) int {
	if a > math.MaxInt-b {
		panic("переполнение int")
	}
	return a + b
}

// Умножение без переполнения
func MultInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	res := a * b
	if a == 1 || b == 1 {
		return res
	}
	if a == math.MinInt || b == math.MinInt {
		panic("переполнение int")
	}
	if res/b != a {
		panic("переполнение int")
	}
	return res
}

// Попытка инкремента
func TryIncInt(counter int) {
	fmt.Printf("%d + 1 = ", counter)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	res := IncInt(counter)
	fmt.Println(res)
}

// Попытка сложения
func TryAddInt(a, b int) {
	fmt.Printf("%d + %d = ", a, b)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	res := AddInt(a, b)
	fmt.Println(res)
}

// Попытка умножения
func TryMultInt(a, b int) {
	fmt.Printf("%d * %d = ", a, b)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	res := MultInt(a, b)
	fmt.Println(res)
}

func main() {
	fmt.Println(" \n[ ПЕРЕПОЛНЕНИЕ ]\n ")

	// Числа
	full := math.MaxInt
	half := full / 2
	quart := half / 2

	// Переполнение при инкременте
	fmt.Println("Инкремент:")
	TryIncInt(quart)
	TryIncInt(half)
	TryIncInt(full - 1)
	TryIncInt(full)
	fmt.Println()

	// Переполнение при сложении
	fmt.Println("Сложение:")
	TryAddInt(quart, quart)
	TryAddInt(half, quart)
	TryAddInt(half, half)
	TryAddInt(half+1, half+1)
	TryAddInt(full, half)
	TryAddInt(full, full)
	TryAddInt(full, 1)
	fmt.Println()

	// Переполнение при умножении
	fmt.Println("Умножение:")
	TryMultInt(quart, 0)
	TryMultInt(quart, 1)
	TryMultInt(quart, 2)
	TryMultInt(quart, 3)
	TryMultInt(quart, 4)
	TryMultInt(quart+1, 4)
	TryMultInt(quart, 5)
	TryMultInt(half, 2)
	TryMultInt(half+1, 2)
	TryMultInt(half, 3)
	TryMultInt(full, 1)
	TryMultInt(full, 2)
	TryMultInt(full+1, 1)
	TryMultInt(full+1, 2)
}
