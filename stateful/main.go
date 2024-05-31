package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Структура "операция чтения"
type readOp struct {
	key  int
	resp chan int
}

// Структура "операция записи"
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	fmt.Println(" \n[ STATEFUL-ГОРУТИНЫ ]\n ")

	// Счетчики и каналы
	var readOps, writeOps uint64
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Чтение или запись
	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Планирование чтения
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Планирование записи
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Ожидание
	fmt.Println("Ожидание...")
	fmt.Println()
	time.Sleep(time.Second)

	// Вывод
	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("Операций чтения:", readOpsFinal)
	fmt.Println("Операций записи:", writeOpsFinal)
}
