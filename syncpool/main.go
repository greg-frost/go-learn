package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Пул объектов
var pool = sync.Pool{
	New: func() any {
		fmt.Println("(инициализация буфера)")
		return make([]byte, 50)
	},
}

// Заполнение случайными значениями
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
	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			buf := pool.Get().([]byte) // Получение буфера из пула
			buf = buf[:0]              // Очистка буфера
			defer pool.Put(buf)        // Возврат буфера в пул
			FillRandom(&buf)           // Использование буфера
			fmt.Println(string(buf))
		}()
		if i%5 == 0 {
			time.Sleep(50 * time.Millisecond)
		}
	}

	wg.Wait()
}
