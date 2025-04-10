package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"
)

// Счетчик
var counter int

// Канал
var ch = make(chan byte)

// Параметр
var n = flag.Int("n", 1e5, "Количество горутин, которые нужно создать")

// Горутина
func f() {
	counter++
	<-ch // Блокировка
}

func main() {
	fmt.Println(" \n[ ЧИСЛО ГОРУТИН ]\n ")

	flag.Parse()
	if *n <= 0 {
		log.Fatal("Неверное количество горутин")
	}

	// Ограничение свободных потоков ОС до 1
	runtime.GOMAXPROCS(1)

	// Копия MemStats
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)

	t0 := time.Now().UnixNano()
	for i := 0; i < *n; i++ {
		go f()
	}
	runtime.Gosched() // Вызов планировщика
	t1 := time.Now().UnixNano()
	runtime.GC() // Вызов сборщика мусора

	// Копия MemStats
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	if counter != *n {
		log.Fatal("Не удалось запустить все горутины")
	}

	fmt.Printf("Горутин: %d\n", *n)
	fmt.Printf("Для каждой:\n")
	fmt.Printf("  Память: %.2f байт\n", float64(m1.Sys-m0.Sys)/float64(*n))
	fmt.Printf("  Время:  %f µs\n", float64(t1-t0)/float64(*n)/1000)
}
