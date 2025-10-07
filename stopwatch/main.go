package main

import (
	"fmt"
	"time"
)

// Структура "секундомер"
type Stopwatch struct {
	start  time.Time
	splits []time.Time
}

// Старт
func (s *Stopwatch) Start() {
	s.start = time.Now()
	s.splits = nil
}

// Сохранение промежуточного результата
func (s *Stopwatch) SaveSplit() {
	s.splits = append(s.splits, time.Now())
}

// Все результаты
func (s Stopwatch) GetResults() (results []time.Duration) {
	for _, v := range s.splits {
		results = append(results, v.Sub(s.start))
	}
	return
}

func main() {
	fmt.Println(" \n[ СЕКУНДОМЕР ]\n ")

	var sw Stopwatch
	sw.Start()
	fmt.Println("Нажимайте Enter, чтобы делать замеры времени...")

	times := 3

	for i := 0; i < times; i++ {
		fmt.Scanln()
		fmt.Printf("Замер сделан, осталось: %d", times-i-1)
		sw.SaveSplit()
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Замеры времени:", sw.GetResults())
}
