package main

import (
	"fmt"
)

// Часы Лэмпорта
type LamportClock struct {
	Value uint64
}

// Локальное событие
func (c *LamportClock) LocalEvent() {
	c.Value++
}

// Получение события
func (c *LamportClock) Receive(remote uint64) {
	if remote > c.Value {
		c.Value = remote
	}
	c.Value++
}

func main() {
	fmt.Println(" \n[ ЛОГИЧЕСКИЕ ЧАСЫ ]\n ")
}
