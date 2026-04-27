package main

import (
	"fmt"
	"time"
)

// Структура "пожертвования"
type Donation struct {
	balance int
	ch      chan int
}

func main() {
	fmt.Println(" \n[ SYNC-COND ]\n ")

	donation := &Donation{ch: make(chan int)}

	// Горутина-слушатель
	listener := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d цель достигнута\n", donation.balance)
				return
			}
		}
	}

	// Прослушивание пожертвований
	go listener(10)
	go listener(15)

	// Осуществление пожертвований
	go func() {
		fmt.Println("Осуществление пожертвований...")
		for {
			time.Sleep(250 * time.Millisecond)
			donation.balance++
			donation.ch <- donation.balance
		}
	}()

	// Ожидание ввода
	var input string
	fmt.Scanln(&input)
}
