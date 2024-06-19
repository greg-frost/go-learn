package main

import (
	"errors"
	"fmt"
	"io"
)

// Структура "мой райтер"
type MyWriter struct {
	buf []byte
}

// Правильная сигнатура метода
func (m *MyWriter) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, errors.New("Нечего записывать!")
	}
	m.buf = b
	return len(b), nil
}

// Неправильная сигнатура метода
// func (m *MyWriter) Write(b []byte) error {
// 	if len(b) == 0 {
// 		return errors.New("Нечего записывать!")
// 	}
// 	m.buf = b
// 	return nil
// }

func main() {
	fmt.Println(" \n[ КАНАРЕЕЧНЫЕ ТЕСТЫ ]\n ")

	// Канарейка прямо в коде
	var _ io.Writer = &MyWriter{}

	fmt.Println("MyWriter соответствует интерфейсу io.Writer!")
	fmt.Println("(а иначе бы код даже не скомпилировался...)")
	fmt.Println()

	// Утверждение интерфейса (возможна паника)
	var myWriter interface{} = &MyWriter{}
	writer := myWriter.(io.Writer)
	_ = writer

	fmt.Println("Утверждение интерфейса также прошло успешно.")
}
