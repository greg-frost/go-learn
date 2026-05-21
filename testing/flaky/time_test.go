package flaky

import (
	"testing"
	"time"
)

// Хрупкий тест (brittle)
func TestTrimOlderThanBrittle(t *testing.T) {
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

// Тест с зависимостью
func TestTrimOlderThanDep(t *testing.T) {
	events := []Event{
		{Timestamp: parseTime(t, "2026-05-21T06:00:00.04Z")},
		{Timestamp: parseTime(t, "2026-05-21T06:00:00.05Z")},
		{Timestamp: parseTime(t, "2026-05-21T06:00:00.06Z")},
	}
	cache := &Cache{now: func() time.Time {
		return parseTime(t, "2026-05-21T06:00:00.06Z")
	}}
	cache.Add(events)

	cache.TrimOlderThanDep(15 * time.Millisecond)
	trimmed := cache.Events()
	n := 2

	if len(trimmed) != n {
		t.Fatal("Количество: опубликовано:", len(trimmed), "ожидается", n)
	}
}

// Тест с параметром
func TestTrimOlderThanParam(t *testing.T) {
	events := []Event{
		{Timestamp: parseTime(t, "2026-05-21T06:00:00.04Z")},
		{Timestamp: parseTime(t, "2026-05-21T06:00:00.05Z")},
		{Timestamp: parseTime(t, "2026-05-21T06:00:00.06Z")},
	}
	cache := new(Cache)
	cache.Add(events)

	cache.TrimOlderThanParam(
		parseTime(t, "2026-05-21T06:00:00.06Z"),
		15*time.Millisecond,
	)
	trimmed := cache.Events()
	n := 2

	if len(trimmed) != n {
		t.Fatal("Количество: опубликовано:", len(trimmed), "ожидается", n)
	}
}

// Парсинг времени
func parseTime(t *testing.T, timestamp string) time.Time {
	res, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		t.Fail()
	}
	return res
}
