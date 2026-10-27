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

// Локальное событие
func (c *LamportClock) Tick() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	return c.value
}

// Получение события
func (c *LamportClock) Update(received int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	if received > c.value {
		c.value = received
	}
	c.value++
	return c.value
}

// Структура "сообщение Лэмпорта"
type LamportMessage struct {
	SenderID int
	Clock    int
	Payload  string
}

// Запуск общения часов Лэмпорта
func RunLamportClocks(id int, send, recv chan LamportMessage, wg *sync.WaitGroup, d time.Duration) {
	clock := new(LamportClock)

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
