package flaky

import (
	"sync"
	"testing"
	"time"
)

// Структура "заглушка публикатора"
type publisherMock struct {
	mu     sync.RWMutex
	values []Value
}

// Публикация
func (p *publisherMock) Publish(values []Value) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.values = values
}

// Получение
func (p *publisherMock) Get() []Value {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.values
}

// Нестабильный тест (flaky)
func TestGetBestValueFlaky(t *testing.T) {
	var mock publisherMock
	n, m := 3, 5
	h := Handler{
		publisher: &mock,
		n:         n,
	}

	value := h.GetBestValue(m)
	time.Sleep(10 * time.Millisecond) // Пауза, порождающая нестабильность
	published := mock.Get()

	if len(published) != n {
		t.Fatal("Количество: опубликовано:", len(published), "ожидается", n)
	}
	if value != published[0] {
		t.Error("Значение: получено", value, ", опубликовано", published[0])
	}
}

// Функция проверки утверждения
func assert(t *testing.T, assertion func() bool,
	maxRetry int, waitTime time.Duration, errorMsg string) {
	for i := 0; i < maxRetry; i++ {
		if assertion() {
			return
		}
		time.Sleep(waitTime)
	}
	t.Fatal(errorMsg)
}

// Тест с повторами
func TestGetBestValueRetries(t *testing.T) {
	var mock publisherMock
	n, m := 3, 5
	h := Handler{
		publisher: &mock,
		n:         n,
	}

	value := h.GetBestValue(m)

	assert(t, func() bool {
		return len(mock.Get()) == n
	}, 30, time.Millisecond, "Неверное количество")

	assert(t, func() bool {
		return mock.Get()[0] == value
	}, 30, time.Millisecond, "Неверное значение")
}

// Структура "заглушка публикатора (с каналом)"
type publisherChanMock struct {
	ch chan []Value
}

// Публикация
func (p *publisherChanMock) Publish(values []Value) {
	p.ch <- values
}

// Получение
func (p *publisherChanMock) Get() []Value {
	return <-p.ch
}

// Тест с каналом
func TestGetBestValueChan(t *testing.T) {
	mock := publisherChanMock{
		ch: make(chan []Value),
	}
	defer close(mock.ch)
	n, m := 3, 5
	h := Handler{
		publisher: &mock,
		n:         n,
	}

	value := h.GetBestValue(m)
	published := mock.Get()

	if len(published) != n {
		t.Fatal("Количество: опубликовано:", len(published), "ожидается", n)
	}
	if value != published[0] {
		t.Error("Значение: получено", value, ", опубликовано", published[0])
	}
}
