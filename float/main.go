package main

import (
	"fmt"
)

// Менее точная функция
func f1(n int) float64 {
	res := 10_000.
	for i := 0; i < n; i++ {
		res += 1.0001
	}
	return res
}

// Более точная функция
func f2(n int) float64 {
	var res float64
	for i := 0; i < n; i++ {
		res += 1.0001
	}
	return res + 10_000.
}

func main() {
	fmt.Println(" \n[ ЧИСЛА С ПЛАВАЮЩЕЙ ТОЧКОЙ ]\n ")

	// Сравнение точности
	fmt.Println("Сравнение точности:")
	size := 1000
	precise := 11000.1
	n1, n2 := f1(size), f2(size)
	fmt.Printf("f1 = %v (%v)\n", n1, precise-n1)
	fmt.Printf("f2 = %v (%v)\n", n2, precise-n2)
}
