package main

import (
	"testing"
)

// Структура "моковое сообщение"
type MockMessage struct {
	email   string
	subject string
	body    []byte
}

// Отправка мокового сообщения
func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

// Тестирование сообщения
func TestAlert(t *testing.T) {
	msgr := &MockMessage{}
	subject := "Абстракция"
	body := []byte("Тестовое сообщение")

	Alert(msgr, body)

	if msgr.subject != subject {
		t.Errorf("Тема - ожидается: %s, получено: %s", subject, msgr.subject)
	}
	if string(msgr.body) != string(body) {
		t.Errorf("Сообщение - ожидается: %s, получено: %s", body, msgr.body)
	}
}
