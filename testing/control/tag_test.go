//go:build integration

package control_test

import (
	"testing"
	"time"
)

func TestIntegrationTag(t *testing.T) {
	t.Log("Интеграционный тест:")
	t.Log("(выполнение...)")
	time.Sleep(3 * time.Second)
}
