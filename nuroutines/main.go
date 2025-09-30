package main

import (
	"fmt"
)

// Генерация значений в канал
func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()

	return ch, cancel
}

func main() {
	fmt.Println(" \n[ НЮАНСЫ ГОРУТИН ]\n ")

	// Взаимоблокировка (deadlock)
	fmt.Println("Взаимоблокировка:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		v1 := "Lock"
		ch1 <- v1 // Ожидаем чтения из ch1
		v2 := <-ch2
		fmt.Println(v1, v2)
	}()

	v1 := "Dead"
	// ch2 <- v1   // Ожидаем чтения из ch2
	// v2 := <-ch1 // Будет блокировка
	var v2 string
	select {
	case ch2 <- v1:
	case v2 = <-ch1:
	}
	fmt.Println(v1, v2)
	fmt.Println()

	// Замыкание
	fmt.Println("Замыкание:")
	a := []int{2, 4, 6, 8, 10}
	ch3 := make(chan int, len(a))
	for _, v := range a {
		//v := v // Можно так ...
		go func(v int) {
			ch3 <- v * 2
		}(v) // ... или так
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch3, " ")
	}
	fmt.Println()
	fmt.Println()

	// Отмена
	fmt.Println("Отмена:")
	ch4, cancel := countTo(10)
	for i := range ch4 {
		if i > 5 {
			break
		}
		fmt.Print(i, " ")
	}
	cancel()
	fmt.Println()
	fmt.Println()

	// Закрытие
	fmt.Println("Закрытие:")
	n := 5
	ch5 := make(chan int, n)
	for i := 0; i < n; i++ {
		ch5 <- i * 10
	}
	// Чтение нескольких значений
	for i := 0; i < 3; i++ {
		fmt.Print(<-ch5, " ")
	}
	close(ch5) // Закрытие на середине
	fmt.Print("(канал закрыт) ")
	// Range дочитает закрытый канал
	for i := range ch5 {
		fmt.Print(i, " ")
	}
	// Comma-Ok тоже дочитает
	for {
		i, ok := <-ch5
		if !ok {
			break
		}
		fmt.Print(i, ok, " ")
	}
	fmt.Println()
}
