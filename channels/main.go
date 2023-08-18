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
	go channelRoutine(c, ms)

	for i := from; i <= to; i++ {
		c <- i
	}
}

// Рутина внутри канала
func channelRoutine(c chan int, ms int) {
	i := 0
	for i >= 0 {
		i = <-c
		fmt.Print(i, " ")
		waitMs(ms)
	}
}

// Сумма (через каналы)
func chanSum(s []int, c chan int) {
	sum := 0
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

// Удвоение
func double(val int) int {
	return val * 2
}

// Утроение
func triple(val int) int {
	return val * 3
}

// IN-OUT каналы
func runInOut(in <-chan int, out chan<- int) {
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

// Пинг
func pinger(c chan string) {
	for {
		c <- "Ping"
	}
}

// Понг
func ponger(c chan<- string) {
	for {
		c <- "Pong"
	}
}

// Принтер
func printer(c <-chan string) {
	for {
		fmt.Print(<-c, " ")
		waitSec(1)
	}
}

// Пауза в мс
func waitMs(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

// Пауза в сек
func waitSec(sec int) {
	waitMs(sec * 1000)
}

func main() {
	fmt.Println(" \n[ КАНАЛЫ ]\n ")

	/* Каналы */

	fmt.Println("С каналами:")
	channel(1, 3, 1)
	channel(4, 7, 10)
	channel(8, 12, 100)

	fmt.Println(" \n ")

	/* Рутины */

	fmt.Println("Без каналов:")
	go routine(1, 3, 1)
	go routine(4, 7, 10)
	go routine(8, 12, 100)

	waitSec(1)
	fmt.Println(" \n ")

	/* Сумма (через каналы) */

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go chanSum(s[:len(s)/2], c)
	go chanSum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println("Сумма (через каналы):", x, y, x+y)

	/* Фибоначчи (через каналы) */

	fmt.Println()

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

	fmt.Println(" \n ")

	/* IN-OUT каналы */

	in := make(chan int)
	out := make(chan int)

	runInOut(in, out)

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

	/* Пинг-понг */

	fmt.Println(" \n ")

	p := make(chan string)

	go pinger(p)
	go ponger(p)
	go printer(p)

	/* Ожидание ввода */

	var input string
	fmt.Scanln(&input)
}
