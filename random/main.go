package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Набор случайных чисел
func randomInts(n, from, to int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = from + rand.Intn(to-from+1)
	}
	return res
}

func main() {
	fmt.Println(" \n[ РАНДОМ ]\n ")

	// Seed не задан
	fmt.Println("Случайные числа:")
	fmt.Println(randomInts(10, 1, 10))
	fmt.Println(randomInts(10, 1, 10))
	fmt.Println(randomInts(10, 1, 10))
	fmt.Println()

	// Seed задан определенно
	fmt.Println("Неслучайные числа:")
	rand.Seed(0)
	fmt.Println(randomInts(10, 1, 10))
	fmt.Println(randomInts(10, 1, 10))
	fmt.Println(randomInts(10, 1, 10))
	fmt.Println()

	// Seed задан произвольно
	fmt.Println("Снова случайные числа:")
	rand.Seed(time.Now().Unix())
	fmt.Println(randomInts(10, 1, 10))
	fmt.Println(randomInts(10, 1, 10))
	fmt.Println(randomInts(10, 1, 10))
}
