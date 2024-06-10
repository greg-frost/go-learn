package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Тип "функция"
type function func()

// Безопасное выполнение
func safelyGo(do function) {
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
func processMessage() {
	fmt.Println("Подготовка сообщения...")
	panic(errors.New("Сообщение испорчено!"))
}

func main() {
	fmt.Println(" \n[ БЕЗОПАСНОЕ ВЫПОЛНЕНИЕ ]\n ")

	// Опасный вызов горутины
	// go processMessage()

	// Безопасный вызов горутины
	safelyGo(processMessage)

	// Ожидание завершения
	time.Sleep(100 * time.Millisecond)
}
