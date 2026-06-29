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
func Primes1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// Простые числа (оптимизированный вариант)
func Primes2(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// Числа Фибоначчи
func Fib1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)
	return int64(Fib2(n-1)) + int64(Fib2(n-2))
}

// Числа Фибоначчи (оптимизированный вариант)
func Fib2(n int) int {
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

	// Файл профиля CPU
	cpuProfile := filepath.Join(os.TempDir(), "cpuProfile.out")
	cpuFile, err := os.Create(cpuProfile)
	if err != nil {
		log.Fatal(err)
	}
	defer cpuFile.Close()

	fmt.Println("Процессор:")
	fmt.Println(cpuProfile)
	fmt.Println()

	// Запуск профилирования CPU
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	// Простые числа
	var total int
	for i := 2; i < 100000; i++ {
		isPrime := Primes1(i)
		if isPrime {
			total++
		}
	}
	fmt.Println("Простых чисел (1):", total)

	// Простые числа (оптимизированный вариант)
	total = 0
	for i := 2; i < 100000; i++ {
		isPrime := Primes2(i)
		if isPrime {
			total++
		}
	}
	fmt.Println("Простых чисел (2):", total)
	fmt.Println()

	// Числа Фибоначии
	fmt.Println("Числа Фибоначчи (1):")
	for i := 1; i < 90; i++ {
		fibo := Fib1(i)
		fmt.Print(fibo, " ")
	}
	fmt.Println()
	fmt.Println()

	// Числа Фибоначии (оптимизированный вариант)
	fmt.Println("Числа Фибоначчи (2):")
	for i := 1; i < 90; i++ {
		fibo := Fib2(i)
		fmt.Print(fibo, " ")
	}
	fmt.Println()
	fmt.Println()

	// Сборка мусора
	runtime.GC()

	// Файл профиля памяти
	memoryProfile := filepath.Join(os.TempDir(), "memoryProfile.out")
	memoryFile, err := os.Create(memoryProfile)
	if err != nil {
		log.Fatal(err)
	}
	defer memoryFile.Close()

	fmt.Println("Память:")
	fmt.Println(memoryProfile)
	fmt.Println()

	for i := 0; i < 10; i++ {
		s := make([]byte, 50_000_000)
		if s == nil {
			fmt.Println("Ошибка выделения памяти!")
		} else {
			fmt.Println("Выделено 50 МБ...")
		}
		time.Sleep(50 * time.Millisecond)
	}

	// Профилирование памяти
	err = pprof.WriteHeapProfile(memoryFile)
	if err != nil {
		log.Fatal(err)
	}
}
