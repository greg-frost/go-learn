package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Структура "противодавление"
type PressureGauge struct {
	ch chan struct{}
}

// Конструктор противодавления
func NewPG(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

// Обработка процесса
func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("нет места")
	}
}

// Медленная функция
func doSlowThings() string {
	time.Sleep(3 * time.Second)
	return "Done"
}

func main() {
	fmt.Println(" \n[ ПРОТИВОДАВЛЕНИЕ ]\n ")

	/* Настройка обработчиков */

	pg := NewPG(5)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doSlowThings()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})

	/* Запуск сервера */

	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на localhost:8080)")

	http.ListenAndServe("localhost:8080", nil)
}
