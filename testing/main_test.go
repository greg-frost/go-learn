package main

import (
	"fmt"
	"os"
	"testing"
)

// Первый тест
func TestFirst(t *testing.T) {
	t.Log("Запуск первого теста")
}

// Второй тест
func TestSecond(t *testing.T) {
	t.Log("Запуск второго теста")
	t.Cleanup(func() {
		t.Log("Уборка после второго теста")
	})
}

// Третий тест
func TestThird(t *testing.T) {
	t.Log("Запуск третьего теста")
	t.Cleanup(func() {
		t.Log("Вторая уборка после третьего теста")
	})
	t.Cleanup(func() {
		t.Log("Первая уборка после третьего теста")
	})
}

// Действия до тестов
func beforeTests() {
	fmt.Println("Действия до тестов")
}

// Действия после тестов
func afterTests() {
	fmt.Println("Действия после тестов")
}

func TestMain(m *testing.M) {
	beforeTests()   // Действия до тестов
	code := m.Run() // Выполнение всех тестов
	afterTests()    // Действия после тестов
	os.Exit(code)
}
