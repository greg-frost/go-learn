package flaky

import (
	"sync"
	"time"
)

// Структура "событие"
type Event struct {
	Timestamp time.Time
	Data      string
}

// Структура "кэш"
type Cache struct {
	mu     sync.RWMutex
	events []Event
}

// Обрезка старых событий
func (c *Cache) TrimOlderThan(since time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	t := time.Now().Add(-since)
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}
