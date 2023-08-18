package main

import (
	"fmt"
	"time"
)

// Пауза (сон)
func sleep(sec int) {
	select {
	case <-time.After(time.Second * time.Duration(sec)):
		return
	}
}

func main() {
	fmt.Println(" \n[ ВЫБОР КАНАЛА ]\n ")

	ch1 := make(chan string)
	ch2 := make(chan string)

	/* Первый канал */

	go func() {
		count := 1
		for {
			ch1 <- "From 1"
			time.Sleep(time.Second * 2)

			// Закрытие после пары записей
			if count == 3 {
				close(ch1)
				break
			}
			count++
		}
	}()

	/* Второй канал */

	go func() {
		for {
			ch2 <- "From 2"
			time.Sleep(time.Second * 3)
		}
	}()

	/* Выбор канала */

	go func() {
		for {
			select {
			case msg1, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				fmt.Print(msg1, " ")
			case msg2, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				fmt.Print(msg2, " ")
			case <-time.After(time.Second):
				fmt.Print("(tic) ")
			}
		}
	}()

	/* Сон */

	s := 5
	fmt.Print("*** Засыпаем на ", s, " сек... *** ")
	sleep(s)
	fmt.Print("*** Проснулись! *** ")

	/* Ожидание ввода */

	var input string
	fmt.Scanln(&input)
}
