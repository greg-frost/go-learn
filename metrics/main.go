package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
	"time"
)

// Метрики
const numRoutines = "/sched/goroutines:goroutines"   // Количество горутин
const freeMemory = "/memory/classes/heap/free:bytes" // Освобожденная память

func main() {
	fmt.Println(" \n[ МЕТРИКИ ]\n ")

	// Срез для метрик
	ms := []metrics.Sample{
		{Name: numRoutines},
		{Name: freeMemory},
	}

	fmt.Println("Подождите...")

	// Создание горутин и выделение памяти
	for i := 0; i < 3; i++ {
		go func() {
			_ = make([]int, 1_000_000)
			time.Sleep(3 * time.Second)
		}()
	}

	time.Sleep(time.Second)

	// Сборка мусора
	runtime.GC()

	// Чтение метрик
	fmt.Println()
	metrics.Read(ms)
	for _, m := range ms {
		if m.Value.Kind() == metrics.KindBad {
			fmt.Printf("Метрика %q более не поддерживается", m.Name)
			continue
		}
		fmt.Println(m.Name, "-", m.Value.Uint64())
	}
	fmt.Println()

	// Список доступных метрик
	fmt.Println("Все метрики:")
	for _, m := range metrics.All() {
		fmt.Println()
		fmt.Println(m.Name, "-", m.Description)
	}
}
