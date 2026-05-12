package control_test

import (
	"os"
	"testing"
	"time"
)

func TestIntegrationEnv(t *testing.T) {
	t.Log("Интеграционный тест:")
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("(пропуск)")
	}
	t.Log("(выполнение...)")
	time.Sleep(3 * time.Second)
}
