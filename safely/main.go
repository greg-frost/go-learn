package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Тип "функция"
type Function func()

// Безопасное выполнение
func SafelyGo(do Function) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Перехват паники: %s", err)
			}
		}()
		do()
	}()
}

// Обработка сообщения
func ProcessMessage() {
	fmt.Println("Подготовка сообщения...")
	panic(errors.New("Сообщение испорчено!"))
}

func main() {
	fmt.Println(" \n[ БЕЗОПАСНОЕ ВЫПОЛНЕНИЕ ]\n ")

	// Опасный вызов горутины
	// go processMessage()

	// Безопасный вызов горутины
	SafelyGo(ProcessMessage)

	// Ожидание завершения
	time.Sleep(100 * time.Millisecond)
}
