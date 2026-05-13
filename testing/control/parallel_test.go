package control_test

import (
	"testing"
	"time"
)

func TestParallelA(t *testing.T) {
	t.Parallel()
	t.Log("Параллельный тест А...")
	time.Sleep(3 * time.Second)
}

func TestParallelB(t *testing.T) {
	t.Parallel()
	t.Log("Параллельный тест B...")
	time.Sleep(3 * time.Second)
}

func TestParallelC(t *testing.T) {
	t.Parallel()
	t.Log("Параллельный тест C...")
	time.Sleep(3 * time.Second)
}
