package main

import (
	"fmt"
	"time"
)

// Структура "сообщение"
type Message struct {
	sendedAt time.Time
}

// Отправка сообщения (через метод)
func (m *Message) Send(email, subject string, body []byte) error {
	fmt.Printf("Адрес: %q\nТема:  %q\nТекст: %q\n", email, subject, body)
	m.sendedAt = time.Now()
	return nil
}

// Интерфейс "сообщенец"
type Messager interface {
	Send(email, subject string, body []byte) error
}

// Отправка сообщения (через интерфейс)
func Alert(m Messager, text []byte) error {
	return m.Send("admin@example.com", "Абстракция", text)
}

func main() {
	fmt.Println(" \n[ MOCK-ТЕСТИРОВАНИЕ ]\n ")

	// Провайдер сообщений
	msgr := &Message{}

	// Без интерфейса
	fmt.Println("Обычное сообщение:")
	msgr.Send("noreply@example.com", "Конкретика", []byte("Использование метода структуры"))
	fmt.Println()

	// С интерфейсом
	fmt.Println("Использование интерфейса:")
	Alert(msgr, []byte("Использование функции с интерфейсом"))
}
