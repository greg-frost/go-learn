package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура "пожертвования"
type Donation struct {
	balance int
	mu      sync.RWMutex
}

func main() {
	fmt.Println(" \n[ SYNC-COND ]\n ")

	donation := new(Donation)

	// Горутина-слушатель
	listener := func(goal int) {
		donation.mu.RLock()
		// Освобождение и захват мьютекса
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d цель достигнута\n", donation.balance)
		donation.mu.RUnlock()
	}

	// Прослушивание пожертвований
	go listener(10)
	go listener(15)

	// Осуществление пожертвований
	go func() {
		fmt.Println("Осуществление пожертвований...")
		for {
			time.Sleep(250 * time.Millisecond)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()

	// Ожидание ввода
	var input string
	fmt.Scanln(&input)
}
