package main

import (
	"fmt"
	"sync"
	"time"
)

// Привет
func doHello() {
	time.Sleep(time.Second)
	fmt.Print("Hello ")
}

// Жестокий
func doCruel() {
	time.Sleep(2 * time.Second)
	fmt.Print("Cruel ")
}

// Мир
func doWorld() {
	time.Sleep(3 * time.Second)
	fmt.Print("World ")
}

// Чтение из канала и пропуск через фукнцию-процессор
func GatherAndProcess(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}

	return result
}

// Функция-процессор (квадратный корень)
func sqrt(x int) int {
	return x * x
}

func main() {
	fmt.Println(" \n[ СИНХРОНИЗАЦИЯ ]\n ")

	// WaitGroup 1
	fmt.Println("WaitGroup 1:")
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		doHello()
	}()
	go func() {
		defer wg.Done()
		doCruel()
	}()
	go func() {
		defer wg.Done()
		doWorld()
	}()
	wg.Wait()
	fmt.Println(" \n ")

	// WaitGroup 2
	fmt.Println("WaitGroup 2:")
	const size = 10
	res := make([]int, size)
	in := make(chan int, size)
	for i := 0; i < size; i++ {
		in <- i + 1
	}
	close(in)
	copy(res, GatherAndProcess(in, sqrt, size))
	fmt.Println(res)
	fmt.Println()

	// Once
	fmt.Println("Однократно:")
	var once sync.Once
	for i := 0; i < 2; i++ {
		once.Do(func() {
			fmt.Print("Once")
		})
		fmt.Print(" upon a time")
	}
	fmt.Println()
}
