package transact

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"go-learn/cloud/core"
)

// Структура "регистратор транзакций в файл"
type fileTransactionLogger struct {
	events       chan<- core.Event
	errors       <-chan error
	lastSequence uint64
	file         *os.File
}

// Параметры событий
const (
	FileEventsCapacity = 16               // Размер буфера
	FileEventsFormat   = "%d\t%d\t%s\t%q" // Формат файла
)

// Запись транзакции добавления
func (l *fileTransactionLogger) WritePut(key, value string) {
	l.events <- core.Event{EventType: core.EventPut, Key: key, Value: value}
}

// Запись транзакции удаления
func (l *fileTransactionLogger) WriteDelete(key string) {
	l.events <- core.Event{EventType: core.EventDelete, Key: key}
}

// Получение канала ошибок
func (l *fileTransactionLogger) Err() <-chan error {
	return l.errors
}

// Запуск регистратора
func (l *fileTransactionLogger) Run() {
	events := make(chan core.Event, FileEventsCapacity)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		for e := range events {
			l.lastSequence++

			_, err := fmt.Fprintf(l.file, FileEventsFormat+"\n",
				l.lastSequence, e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- err
				return
			}
		}
	}()
}

// Чтение событий
func (l *fileTransactionLogger) Read() (<-chan core.Event, <-chan error) {
	scanner := bufio.NewScanner(l.file)
	events := make(chan core.Event)
	errors := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errors)

		var e core.Event

		for scanner.Scan() {
			line := scanner.Text()

			_, err := fmt.Sscanf(line, FileEventsFormat,
				&e.Sequence, &e.EventType, &e.Key, &e.Value)
			if err != nil && err != io.EOF {
				errors <- fmt.Errorf("ошибка парсинга файла: %w", err)
				return
			}

			if l.lastSequence >= e.Sequence {
				errors <- fmt.Errorf("последовательность транзакций нарушена")
				return
			}
			l.lastSequence = e.Sequence

			events <- e
		}

		if err := scanner.Err(); err != nil {
			errors <- fmt.Errorf("ошибка сканирования файла: %w", err)
			return
		}
	}()

	return events, errors
}

// Конструктор регистратора
func newFileTransactionLogger(filename string) (core.TransactionLogger, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл регистратора транзакций: %w", err)
	}

	return &fileTransactionLogger{file: file}, nil
}
