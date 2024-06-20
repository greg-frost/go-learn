package main

import (
	"testing"
	"testing/quick"
)

// Обычное тестирование
func TestPad(t *testing.T) {
	size := 6
	if r := Pad("test", uint(size)); len(r) != size {
		t.Errorf("Длина - ожидается: %d, получено: %d", size, len(r))
	}
}

// Порождающее тестирование
func TestPadGenerative(t *testing.T) {
	fn := func(s string, max uint8) bool {
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}
	if err := quick.Check(fn, &quick.Config{MaxCount: 200}); err != nil {
		t.Error(err)
	}
}
