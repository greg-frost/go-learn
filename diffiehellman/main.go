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

// Возведение в степень
func pow(x, n int) int {
	return int(math.Pow(
		float64(x),
		float64(n),
	))
}

// Быстрое возведение в степень
func fastPow(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		n /= 2
		x *= x
	}
	return res
}

// Быстрое возведение в степень по модулю
func fastPowMod(x, n, p int) int {
	res := 1
	x %= p
	for n > 0 {
		if n&1 == 1 {
			res = res * x % p
		}
		n >>= 1
		x = x * x % p
	}
	return res
}

func main() {
	fmt.Println(" \n[ DIFFIE-HELLMAN ]\n ")

	// Генерация открытых чисел
	fmt.Println("Выбор двух открытых чисел:")
	var p int
	for !isPrime(p) {
		p = random(1e6, 1e9) // Простое число p
	}
	fmt.Printf("p = %d\n", p)

	g := random(2, p-2) // Число g (от 2 до p-2)
	fmt.Printf("g = %d\n\n", g)

	fmt.Println("...")
	fmt.Println()

	// Действия Алисы
	fmt.Println("Алиса придумывает секрет:")
	a := random(1, p-1) // Число a (от 1 до p-1)
	fmt.Printf("a = %d\n\n", a)

	fmt.Println("Алиса рассчитывает открытое число на основе секрета:")
	A := fastPowMod(g, a, p) // Число A (g^a mod p)
	fmt.Printf("A = %d\n\n", A)

	fmt.Println("Алиса отправляет свое открытое число Бобу.")
	fmt.Println()

	fmt.Println("...")
	fmt.Println()

	// Действия Боба
	fmt.Println("Боб придумывает секрет:")
	b := random(1, p-1) // Число b (от 1 до p-1)
	fmt.Printf("b = %d\n\n", b)

	fmt.Println("Боб рассчитывает открытое число на основе секрета:")
	B := fastPowMod(g, b, p) // Число B (g^b mod p)
	fmt.Printf("B = %d\n\n", B)

	fmt.Println("Боб отправляет свое открытое число Алисе.")
	fmt.Println()

	fmt.Println("...")
	fmt.Println()

	// Генерация ключей
	fmt.Println("Алиса получает ключ из числа Боба и своего секрета:")
	aliceKey := fastPowMod(B, a, p) // Ключ Алисы (B^a mod p)
	fmt.Printf("aliceKey = %d\n\n", aliceKey)

	fmt.Println("Боб получает ключ из числа Алисы и своего секрета:")
	bobKey := fastPowMod(A, b, p) // Ключ Боба (A^b mod p)
	fmt.Printf("bobKey = %d\n\n", bobKey)

	fmt.Println("Ключи идентичны:", aliceKey == bobKey)
}
