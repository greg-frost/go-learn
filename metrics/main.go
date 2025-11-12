package main

import (
	"fmt"
	"runtime"
	"runtime/metrics"
	"sync"
	"time"
)

func main() {
	fmt.Println(" \n[ МЕТРИКИ ]\n ")

	// Метрики
	routinesCount := "/sched/goroutines:goroutines" // Число горутин
	freeMemory := "/memory/classes/heap/free:bytes" // Освобожденная память

	// Срез для метрик
	ms := []metrics.Sample{
		{Name: routinesCount},
		{Name: freeMemory},
	}

	fmt.Println("Подождите...")

	// Создание горутин и выделение памяти
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = make([]int, 1_000_000)
			time.Sleep(3 * time.Second)
		}()
	}

	wg.Wait()

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
		fmt.Print(m.Name)
		fmt.Println(" -", m.Description)
		fmt.Println()
	}
}
