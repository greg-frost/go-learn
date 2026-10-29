package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура "часы Лэмпорта"
type LamportClock struct {
	value int
	mu    sync.Mutex
}

// Конструктор
func NewLamportClock() *LamportClock {
	return new(LamportClock)
}

// Локальное событие
func (lc *LamportClock) Tick() int {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.value++
	return lc.value
}

// Получение события
func (lc *LamportClock) Update(received int) int {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	if received > lc.value {
		lc.value = received
	}
	lc.value++
	return lc.value
}

// Структура "сообщение Лэмпорта"
type LamportMessage struct {
	SenderID int
	Clock    int
	Payload  string
}

// Запуск общения часов Лэмпорта
func RunLamportClocks(id int, send, recv chan LamportMessage, wg *sync.WaitGroup, d time.Duration) {
	clock := NewLamportClock()

	// Локальное событие
	t1 := clock.Tick()
	fmt.Printf("[P%d] локальное событие #1, время=%d\n", id, t1)

	time.Sleep(d)

	// Отправляем сообщение
	// (перед отправкой увеличиваем счетчик)
	t2 := clock.Tick()
	msg := LamportMessage{
		SenderID: id,
		Clock:    t2,
		Payload:  "Lamport",
	}
	send <- msg
	fmt.Printf("[P%d] отправил сообщение, время=%d\n", id, t2)

	// Ждем входящее сообщение
	incoming := <-recv
	t3 := clock.Update(incoming.Clock)
	fmt.Printf("[P%d] получил сообщение от P%d (время=%d), новое время=%d, payload=%s\n",
		id, incoming.SenderID, incoming.Clock, t3, incoming.Payload)

	// Еще одно локальное событие
	t4 := clock.Tick()
	fmt.Printf("[P%d] локальное событие #2, время=%d\n", id, t4)

	wg.Done()
}

// Структура "векторные часы"
type VectorClock struct {
	values []int
	mu     sync.Mutex
}

// Конструктор
func NewVectorClock(size int) *VectorClock {
	return &VectorClock{
		values: make([]int, size),
	}
}

// Клонирование часов
func (vc *VectorClock) Clone() *VectorClock {
	res := make([]int, len(vc.values))
	copy(res, vc.values)
	return &VectorClock{
		values: res,
	}
}

// Локальное событие
func (vc *VectorClock) Tick(id int) {
	vc.values[id]++
}

// Слияние часов
func (vc *VectorClock) Merge(other *VectorClock) {
	for id := range vc.values {
		if other.values[id] > vc.values[id] {
			vc.values[id] = other.values[id]
		}
	}
}

// Сравнение часов
func (vc1 *VectorClock) Before(vc2 *VectorClock) bool {
	var less bool
	for id := range vc1.values {
		if vc1.values[id] > vc2.values[id] {
			return false
		}
		if vc1.values[id] < vc2.values[id] {
			less = true
		}
	}
	return less
}

func main() {
	fmt.Println(" \n[ ЛОГИЧЕСКИЕ ЧАСЫ ]\n ")

	fmt.Println("Часы Лэмпорта:")
	var wg sync.WaitGroup
	p1 := make(chan LamportMessage, 1)
	p2 := make(chan LamportMessage, 1)
	wg.Add(2)
	go RunLamportClocks(0, p1, p2, &wg, 200*time.Millisecond)
	go RunLamportClocks(1, p2, p1, &wg, 300*time.Millisecond)
	wg.Wait()
}
