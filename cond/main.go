package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура "пожертвования"
type Donation struct {
	balance int
	cond    *sync.Cond
}

func main() {
	fmt.Println(" \n[ SYNC-COND ]\n ")

	// Пожертвования
	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Горутина-слушатель
	listener := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			// Ожидание выполнения условия
			// (уведомления broadcast)
			donation.cond.Wait()
		}
		fmt.Printf("$%d цель достигнута\n", donation.balance)
		donation.cond.L.Unlock()
	}

	// Прослушивание пожертвований
	go listener(10)
	go listener(15)

	// Осуществление пожертвований
	go func() {
		fmt.Println("Осуществление пожертвований...")
		for {
			time.Sleep(250 * time.Millisecond)

			// Обновление баланса
			donation.cond.L.Lock()
			donation.balance++
			donation.cond.L.Unlock()

			// Уведомление всех ожидающих горутин
			donation.cond.Broadcast()
		}
	}()

	// Ожидание ввода
	var input string
	fmt.Scanln(&input)
}
