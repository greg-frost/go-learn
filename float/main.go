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

	// Сравнение точности при сложении
	fmt.Println("Точность сложения:")
	size := 1000
	precise := 11000.1
	n1, n2 := f1(size), f2(size)
	fmt.Printf("f1 = %v (%v)\n", n1, precise-n1)
	fmt.Printf("f2 = %v (%v) [ точнее ]\n", n2, precise-n2)
	fmt.Println()

	// Сравнение точности при умножении
	fmt.Println("Точность умножения:")
	a := 100000.001
	b := 1.0001
	c := 1.0002
	precise = 200030.002
	n1 = a * (b + c)
	n2 = a*b + a*c
	fmt.Printf("a * (b+c) = %v (%v)\n", n1, n1-precise)
	fmt.Printf("a*b + a*c = %v (%v) [ точнее ]\n", n2, n2-precise)
}
