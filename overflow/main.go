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

func main() {
	fmt.Println(" \n[ ПЕРЕПОЛНЕНИЕ ]\n ")
}
