package control_test

import (
	"testing"
	"time"
)

func TestShortLong(t *testing.T) {
	t.Log("Длительный тест:")
	if testing.Short() {
		t.Skip("(пропуск)")
	}
	t.Log("(выполнение...)")
	time.Sleep(3 * time.Second)
}
