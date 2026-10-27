package main

import (
	"fmt"
	"sync"
	"time"
)

// Часы Лэмпорта
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

// Структура "сообщение"
type Message struct {
	SenderID int
	Clock    int
	Payload  string
}

func runProcess(id int, sendCh, recvCh chan Message, wg *sync.WaitGroup, delay time.Duration) {
	clock := new(LamportClock)

	// Локальное событие
	t1 := clock.Tick()
	fmt.Printf("[P%d] локальное событие #1, время=%d\n", id, t1)

	time.Sleep(delay)

	// Отправляем сообщение
	// (перед отправкой увеличиваем счетчик)
	t2 := clock.Tick()
	msg := Message{
		SenderID: id,
		Clock:    t2,
		Payload:  "Hello",
	}
	sendCh <- msg
	fmt.Printf("[P%d] отправил сообщение, время=%d\n", id, t2)

	// Ждем входящее сообщение
	incoming := <-recvCh
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
	chAtoB := make(chan Message, 1)
	chBtoA := make(chan Message, 1)
	wg.Add(2)
	go runProcess(0, chAtoB, chBtoA, &wg, 200*time.Millisecond)
	go runProcess(1, chBtoA, chAtoB, &wg, 300*time.Millisecond)
	wg.Wait()
}
