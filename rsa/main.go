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

// Взаимно простые ли числа
func isCoprime(n, m int) bool {
	for i := 2; i*i <= n && i*i <= m; i++ {
		if n%i == 0 && m%i == 0 {
			return false
		}
	}
	return true
}

// Алгоритм Евклида (наименьший общий делитель)
func euclid(a, b int) int {
	if b == 0 {
		return a
	}
	return euclid(b, a%b)
}

// Расширенный алгоритм Евклида (разложение на сумму множителей)
func extendedEuclid(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	r, x, y := extendedEuclid(b, a%b)
	return r, y, x - a/b*y
}

// Модульная (мультипликативная) инверсия
func modularInverse(a, n int) int {
	r, x, _ := extendedEuclid(a, n)
	if r != 1 {
		return 0
	}
	if x < 0 {
		return x + n
	}
	return x
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

// Минимальный и максимальный пороги
const (
	min = 1e3
	max = 1e4
)

func main() {
	fmt.Println(" \n[ RSA ]\n ")

	/* Генерация простых чисел */

	fmt.Println("Выбор двух простых чисел:")

	// Простое число p
	var p int
	for !isPrime(p) {
		p = random(min, max)
	}
	fmt.Printf("p = %d\n", p)

	// Простое число q
	var q, diff int
	for !isPrime(q) || q == p || diff < min {
		q = random(min, max)
		diff = int(math.Abs(float64(p - q)))
	}
	fmt.Printf("q = %d\n\n", q)

	// Произведение простых чисел
	fmt.Println("Произведение p и q:")

	n := p * q
	fi := (p - 1) * (q - 1)
	fmt.Printf("n = %d\n\n", n)

	fmt.Println("...")
	fmt.Println()

	/* Генерация ключа шифрования */

	fmt.Println("Выбор числа для шифрования:")

	// Число e
	var e int
	for !isPrime(e) || !isCoprime(e, fi) {
		e = random(min, max)
	}
	//e = 65537
	fmt.Printf("e = %d\n\n", e)

	/* Генерация ключа дешифрования */

	fmt.Println("Выбор числа для дешифрования:")

	// Число d
	d := modularInverse(e, fi)
	fmt.Printf("d = %d\n\n", d)

	fmt.Println("...")
	fmt.Println()

	/* Шифрование и дешифрование */

	fmt.Println("Оригинальное сообщение:")

	// Генерация сообщения (не длиннее n)
	M := random(1, n)
	fmt.Printf("M = %d\n\n", M)

	fmt.Println("Зашифрованное сообщение:")

	// Шифрование (M^e mod n)
	Me := fastPowMod(M, e, n)
	fmt.Printf("M(e) = %d\n\n", Me)

	fmt.Println("Расшифрованное сообщение:")

	// Дешифрование (Me^d mod n)
	Md := fastPowMod(Me, d, n)
	fmt.Printf("M(d) = %d\n\n", Md)

	fmt.Println("Сообщения идентичны:", M == Md)
}
