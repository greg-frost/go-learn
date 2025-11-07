package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"
)

// Простые числа
func primes1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// Простые числа (оптимизированный вариант)
func primes2(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// Числа Фибоначчи
func fibo1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)
	return int64(fibo2(n-1)) + int64(fibo2(n-2))
}

// Числа Фибоначчи (оптимизированный вариант)
func fibo2(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(50 * time.Millisecond)
	return fn[n]
}

func main() {
	fmt.Println(" \n[ ПРОФИЛИРОВАНИЕ ]\n ")

	// Файл профиля
	profile := filepath.Join(os.TempDir(), "profile.out")
	file, err := os.Create(profile)
	if err != nil {
		log.Fatal(err)
	}

	// Запуск профилирования
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	// Простые числа
	var total int
	for i := 2; i < 100000; i++ {
		isPrime := primes1(i)
		if isPrime {
			total++
		}
	}
	fmt.Println("Простых чисел (1):", total)

	// Простые числа (оптимизированный вариант)
	total = 0
	for i := 2; i < 100000; i++ {
		isPrime := primes2(i)
		if isPrime {
			total++
		}
	}
	fmt.Println("Простых чисел (2):", total)
	fmt.Println()

	// Числа Фибоначии
	fmt.Println("Числа Фибоначчи (1):")
	for i := 1; i < 90; i++ {
		fibo := fibo1(i)
		fmt.Print(fibo, " ")
	}
	fmt.Println()
	fmt.Println()

	// Числа Фибоначии (оптимизированный вариант)
	fmt.Println("Числа Фибоначчи (2):")
	for i := 1; i < 90; i++ {
		fibo := fibo2(i)
		fmt.Print(fibo, " ")
	}
	fmt.Println()

	runtime.GC() // Запуск сборщика мусора
}
