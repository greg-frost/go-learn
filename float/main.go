package main

import (
	"fmt"
	"math"
)

// Менее точная функция сложения
func sumLessPrecise(n int) float64 {
	res := 10_000.
	for i := 0; i < n; i++ {
		res += 1.0001
	}
	return res
}

// Более точная функция сложения
func sumMorePrecise(n int) float64 {
	var res float64
	for i := 0; i < n; i++ {
		res += 1.0001
	}
	return res + 10_000.
}

// Менее точная функция умножения
func prodLessPrecise(a, b, c float64) float64 {
	return a * (b + c)
}

// Более точная функция умножения
func prodMorePrecise(a, b, c float64) float64 {
	return a*b + a*c
}

// Сравнение чисел float
func compareFloat(a, b, precise float64) bool {
	return math.Abs(a-b) < precise
}

func main() {
	fmt.Println(" \n[ ЧИСЛА С ПЛАВАЮЩЕЙ ТОЧКОЙ ]\n ")

	// Сравнение точности при сложении
	fmt.Println("Точность сложения:")
	size := 1000
	preciseSum := 11000.1
	sum1 := sumLessPrecise(size)
	sum2 := sumMorePrecise(size)
	fmt.Printf("10000 + sum = %v (%v)\n", sum1, preciseSum-sum1)
	fmt.Printf("sum + 10000 = %v (%v) [ точнее ]\n", sum2, preciseSum-sum2)
	fmt.Println("Равенство:", sum1 == sum2)
	fmt.Println("Сравнение:", compareFloat(sum1, sum2, 1e-6))
	fmt.Println()

	// Сравнение точности при умножении
	fmt.Println("Точность умножения:")
	a := 100000.001
	b := 1.0001
	c := 1.0002
	preciseProd := 200030.002
	prod1 := prodLessPrecise(a, b, c)
	prod2 := prodMorePrecise(a, b, c)
	fmt.Printf("a * (b+c) = %v (%v)\n", prod1, prod1-preciseProd)
	fmt.Printf("a*b + a*c = %v (%v) [ точнее ]\n", prod2, prod2-preciseProd)
	fmt.Println("Равенство:", prod1 == prod2)
	fmt.Println("Сравнение:", compareFloat(prod1, prod2, 1e-6))
}
