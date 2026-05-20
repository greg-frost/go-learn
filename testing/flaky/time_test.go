package flaky

import (
	"testing"
	"time"
)

// Хрупкий тест (brittle)
func TestTrimOlderThan(t *testing.T) {
	events := []Event{
		{Timestamp: time.Now().Add(-20 * time.Millisecond)},
		{Timestamp: time.Now().Add(-10 * time.Millisecond)},
		{Timestamp: time.Now().Add(10 * time.Millisecond)},
	}
	cache := new(Cache)
	cache.Add(events)

	cache.TrimOlderThan(15 * time.Millisecond)
	trimmed := cache.Events()
	n := 2

	if len(trimmed) != n {
		t.Fatal("Количество: опубликовано:", len(trimmed), "ожидается", n)
	}
}
