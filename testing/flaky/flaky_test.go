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
