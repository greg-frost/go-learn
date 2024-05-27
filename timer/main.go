package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" \n[ ТАЙМЕР ]\n ")

	// Ожидание таймера
	t1 := time.NewTimer(time.Second)
	<-t1.C
	fmt.Println("Таймер №1 сработал")

	// Остановка таймера
	t2 := time.NewTimer(time.Second)
	go func() {
		<-t2.C
		fmt.Println("Таймер №2 сработал")
	}()
	stop := t2.Stop()
	if stop {
		fmt.Println("Таймер №2 остановлен")
	}
}
