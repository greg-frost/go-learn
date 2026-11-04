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
	mu     sync.RWMutex
}

// Конструктор
func NewVectorClock(size int) *VectorClock {
	return &VectorClock{
		values: make([]int, size),
	}
}

// Клонирование часов
func (vc *VectorClock) Clone() *VectorClock {
	vc.mu.RLock()
	defer vc.mu.RUnlock()
	res := make([]int, len(vc.values))
	copy(res, vc.values)
	return &VectorClock{
		values: res,
	}
}

// Локальное событие
func (vc *VectorClock) Tick(id int) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	vc.values[id]++
}

// Слияние часов
func (vc *VectorClock) Merge(other *VectorClock) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	for id := range vc.values {
		if other.values[id] > vc.values[id] {
			vc.values[id] = other.values[id]
		}
	}
}

// Сравнение часов
func (vc *VectorClock) Before(other *VectorClock) bool {
	vc.mu.RLock()
	defer vc.mu.RUnlock()
	var less bool
	for id := range vc.values {
		if vc.values[id] > other.values[id] {
			return false
		}
		if vc.values[id] < other.values[id] {
			less = true
		}
	}
	return less
}

// Структура "векторное сообщение"
type VectorMessage struct {
	SenderID int
	Clock    *VectorClock
	Payload  string
}

// Размер векторных часов
const vectorSize = 3

// Запуск общения векторных часов
func RunVectorClocks(id int, r, s1, s2 chan VectorMessage, wg *sync.WaitGroup, d time.Duration) {
	clock := NewVectorClock(vectorSize)

	// Локальное событие
	clock.Tick(id)
	fmt.Printf("[P%d] локальное событие #1, время=%v\n", id, clock.values)

	time.Sleep(d)

	// Отправляем сообщение
	// (перед отправкой увеличиваем счетчик)
	clock.Tick(id)
	msg := VectorMessage{
		SenderID: id,
		Clock:    clock.Clone(),
		Payload:  "Vector",
	}
	s1 <- msg
	s2 <- msg
	fmt.Printf("[P%d] отправил сообщение, время=%v\n", id, clock.values)

	// Ждем входящее сообщение
	incoming := <-r
	clock.Merge(incoming.Clock)
	fmt.Printf("[P%d] получил сообщение от P%d (время=%v), новое время=%v, payload=%s\n",
		id, incoming.SenderID, incoming.Clock.values, clock.values, incoming.Payload)

	// Еще одно локальное событие
	clock.Tick(id)
	fmt.Printf("[P%d] локальное событие #2, время=%d\n", id, clock.values)

	wg.Done()
}

func main() {
	fmt.Println(" \n[ ЛОГИЧЕСКИЕ ЧАСЫ ]\n ")

	var wg sync.WaitGroup

	// Часы Лэмпорта
	fmt.Println("Часы Лэмпорта:")
	l1 := make(chan LamportMessage, 1)
	l2 := make(chan LamportMessage, 1)
	wg.Add(2)
	go RunLamportClocks(0, l1, l2, &wg, 200*time.Millisecond)
	go RunLamportClocks(1, l2, l1, &wg, 300*time.Millisecond)
	wg.Wait()
	fmt.Println()

	// Векторные часы
	fmt.Println("Векторные часы:")
	v1 := make(chan VectorMessage, 2)
	v2 := make(chan VectorMessage, 2)
	v3 := make(chan VectorMessage, 2)
	wg.Add(3)
	go RunVectorClocks(0, v1, v2, v3, &wg, 200*time.Millisecond)
	go RunVectorClocks(1, v2, v1, v3, &wg, 250*time.Millisecond)
	go RunVectorClocks(2, v3, v2, v1, &wg, 300*time.Millisecond)
	wg.Wait()
}
