package main

import (
	"fmt"
	"time"
)

// Рутина
func Routine(from, to, ms int) {
	for i := from; i <= to; i++ {
		fmt.Print(i, " ")
		waitMs(ms)
	}
}

// Канал
func Channel(from, to, ms int) {
	c := make(chan int)
	go ReadChannel(c, ms)

	for i := from; i <= to; i++ {
		c <- i
	}
	close(c)
}

// Чтение канала
func ReadChannel(c chan int, ms int) {
	for i := range c {
		fmt.Print(i, " ")
		waitMs(ms)
	}
}

// Сумма (через каналы)
func ChanSum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// Ряд Фибоначчи (через каналы)
func ChanFibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// IN-OUT каналы
func ChanInOut(in <-chan int, out chan<- int) {
	go func() {
		for val := range in {
			var res int
			switch {
			case val%2 == 0:
				res = Double(val)
			default:
				res = Triple(val)
			}
			out <- res
		}
	}()
}

// Удвоение
func Double(val int) int {
	return val * 2
}

// Утроение
func Triple(val int) int {
	return val * 3
}

// Нулевые каналы
func ChanNil(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					ch1 = nil
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					break
				}
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

// Пинг
func Ping(c chan string) {
	for {
		c <- "Ping"
	}
}

// Понг
func Pong(c chan<- string) {
	for {
		c <- "Pong"
	}
}

// Печать
func Print(c <-chan string) {
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
	go Routine(1, 3, 1)
	go Routine(4, 7, 10)
	go Routine(8, 12, 100)
	waitMs(500)
	fmt.Println()
	fmt.Println()

	// Каналы
	fmt.Println("С каналами:")
	Channel(1, 3, 1)
	Channel(4, 7, 10)
	Channel(8, 12, 100)
	waitMs(250)
	fmt.Println()
	fmt.Println()

	// Сумма (через каналы)
	c := make(chan int)
	s := []int{7, 2, 8, -9, 4, 0}
	go ChanSum(s[:len(s)/2], c)
	go ChanSum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println("Сумма (через каналы):", x, y, x+y)
	fmt.Println()

	// Фибоначчи (через каналы)
	f := make(chan int, 10)
	go ChanFibonacci(cap(f), f)
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
	ChanInOut(in, out)
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

	// Нулевые каналы
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch := ChanNil(ch1, ch2)
	fmt.Println("Нулевые каналы:")
	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	for v := range ch {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println()

	// Пинг-понг
	p := make(chan string)
	go Ping(p)
	go Pong(p)
	go Print(p)

	// Ожидание ввода
	var input string
	fmt.Scanln(&input)
}
