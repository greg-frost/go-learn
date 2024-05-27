package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" \n[ ЛИМИТЕР ]\n ")

	const (
		count = 5
		limit = 3
	)

	/* Равномерный лимитер */

	fmt.Println("Равномерный:")
	fmt.Println()

	limiter := time.Tick(200 * time.Millisecond)

	requests := make(chan int, count)
	for i := 1; i <= count; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-limiter
		fmt.Println("Запрос", req, time.Now())
	}

	/* Всплесковый лимитер */

	fmt.Println()
	fmt.Println("Всплесковый:")
	fmt.Println()

	burstyLimiter := make(chan time.Time, limit)
	for i := 1; i <= limit; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, count)
	for i := 1; i <= count; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Запрос", req, time.Now())
	}
}
