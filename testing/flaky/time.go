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

// Тип "функция текущего времени"
type now func() time.Time

// Структура "кэш"
type Cache struct {
	mu     sync.RWMutex
	events []Event
	now    now
}

// Конструктор
func NewCache() *Cache {
	return &Cache{
		events: make([]Event, 0),
		now:    time.Now,
	}
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

// Обрезка старых событий (с зависимостью)
func (c *Cache) TrimOlderThanDep(since time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	t := c.now().Add(-since)
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}

// Обрезка старых событий (с параметром)
func (c *Cache) TrimOlderThanParam(now time.Time, since time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	t := now.Add(-since)
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}

// Добавление событий
func (c *Cache) Add(events []Event) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.events = append(c.events, events...)
}

// Получение всех событий
func (c *Cache) Events() []Event {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.events
}
