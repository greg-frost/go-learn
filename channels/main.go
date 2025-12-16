package main

import (
	"fmt"
	"time"
)

// Рутина
func routine(from, to, ms int) {
	for i := from; i <= to; i++ {
		fmt.Print(i, " ")
		waitMs(ms)
	}
}

// Канал
func channel(from, to, ms int) {
	c := make(chan int)
	go readChannel(c, ms)

	for i := from; i <= to; i++ {
		c <- i
	}
	close(c)
}

// Чтение канала
func readChannel(c chan int, ms int) {
	for i := range c {
		fmt.Print(i, " ")
		waitMs(ms)
	}
}

// Сумма (через каналы)
func chanSum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// Ряд Фибоначчи (через каналы)
func chanFibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// IN-OUT каналы
func chanInOut(in <-chan int, out chan<- int) {
	go func() {
		for val := range in {
			var res int

			switch {
			case val%2 == 0:
				res = triple(val)
			default:
				res = double(val)
			}

			out <- res
		}
	}()
}

// Удвоение
func double(val int) int {
	return val * 2
}

// Утроение
func triple(val int) int {
	return val * 3
}

// Пинг
func ping(c chan string) {
	for {
		c <- "Ping"
	}
}

// Понг
func pong(c chan<- string) {
	for {
		c <- "Pong"
	}
}

// Печать
func print(c <-chan string) {
	for {
		fmt.Print(<-c, " ")
		waitMs(1000)
	}
}

// Пауза в мс
func waitMs(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {
	fmt.Println(" \n[ КАНАЛЫ ]\n ")

	// Рутины
	fmt.Println("Без каналов:")
	go routine(1, 3, 1)
	go routine(4, 7, 10)
	go routine(8, 12, 100)
	waitMs(500)
	fmt.Println()
	fmt.Println()

	// Каналы
	fmt.Println("С каналами:")
	channel(1, 3, 1)
	channel(4, 7, 10)
	channel(8, 12, 100)
	waitMs(250)
	fmt.Println()
	fmt.Println()

	// Сумма (через каналы)
	c := make(chan int)
	s := []int{7, 2, 8, -9, 4, 0}
	go chanSum(s[:len(s)/2], c)
	go chanSum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println("Сумма (через каналы):", x, y, x+y)
	fmt.Println()

	// Фибоначчи (через каналы)
	f := make(chan int, 10)
	go chanFibonacci(cap(f), f)
	fmt.Println("Ряд Фибоначчи (через каналы):")
	for {
		if v, ok := <-f; ok {
			fmt.Print(v, " ")
		} else {
			break
		}
	}
	fmt.Println()
	fmt.Println()

	// IN-OUT каналы
	in := make(chan int)
	out := make(chan int)
	chanInOut(in, out)
	fmt.Println("IN-OUT каналы:")
	i := 1
	for i <= 10 {
		select {
		case in <- i:
			i++
		case v := <-out:
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
	fmt.Println()

	// Пинг-понг
	p := make(chan string)
	go ping(p)
	go pong(p)
	go print(p)

	// Ожидание ввода
	var input string
	fmt.Scanln(&input)
}
