package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Случайное число в диапазоне
func random(from, to int) int {
	return from + rand.Intn(to-from+1)
}

// Простое ли число
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Быстрое возведение в степень по модулю
func fastPowMod(x, n, p int) int {
	c, d, r := x%p, n, 1
	for d > 0 {
		if d&1 == 1 {
			r = r * c % p
		}
		d >>= 1
		c = c * c % p
	}
	return r
}

func main() {
	fmt.Println(" \n[ RSA ]\n ")

	/* Генерация простых чисел */

	fmt.Println("Выбор двух простых чисел:")

	// Простое числа p
	var p int
	for !isPrime(p) {
		p = random(1e6, 1e9)
	}
	fmt.Printf("p = %d\n", p)

	// Простое число q
	var q, diff int
	for !isPrime(q) || q == p || diff < 1e6 {
		q = random(1e6, 1e9)
		diff = int(math.Abs(float64(p - q)))
	}
	fmt.Printf("q = %d\n\n", q)

	// Произведение этих чисел
	fmt.Println("Произведение p и q:")

	n := p * q
	fmt.Printf("n = %d\n", n)
}
