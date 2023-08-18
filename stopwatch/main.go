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

	sw := Stopwatch{}
	sw.Start()
	fmt.Println("Подождите...")

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println("Замеры времени:", sw.GetResults())
}
