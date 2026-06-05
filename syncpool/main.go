package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		fmt.Println("(Инициализация буфера)")
		fmt.Println()
		return make([]byte, 50)
	},
}

func FillRandom(buf *[]byte) {
	start := 20
	n := start + rand.Intn(cap(*buf)-start)
	for i := 0; i < n; i++ {
		m := rand.Intn(26)
		start := 'a'
		if rand.Intn(2) == 1 {
			start = 'A'
		}
		*buf = append(*buf, byte(int(start)+m))
	}
}

func main() {
	fmt.Println(" \n[ SYNC-POOL ]\n ")

	times := 25
	for i := 0; i < times; i++ {
		func() {
			buf := pool.Get().([]byte) // Получение буфера из пула
			buf = buf[:0]              // Очистка буфера
			defer pool.Put(buf)        // Возврат буфера в пул

			// Использование буфера
			FillRandom(&buf)
			fmt.Println(string(buf))
		}()
	}
}
