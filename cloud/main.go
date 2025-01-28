package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"go-learn/base"

	"github.com/gorilla/mux"
)

// Хранилище пар ключ/значение
var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

// Ошибка поиска ключа
var ErrKeyNotFound = errors.New("ключ не найден")

// Добавление значения по ключу
func Put(key, value string) error {
	store.Lock()
	defer store.Unlock()

	store.m[key] = value

	return nil
}

// Получение значения по ключу
func Get(key string) (string, error) {
	store.RLock()
	defer store.RUnlock()

	value, ok := store.m[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

// Удаление значения по ключу
func Delete(key string) error {
	store.Lock()
	defer store.Unlock()

	delete(store.m, key)

	return nil
}

// Структура "событие"
type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

// Тип события
type EventType byte

// Виды типов событий
const (
	_                  = iota
	EventPut EventType = iota
	EventDelete
)

// Интерфейс "регистратор транзакций"
type TransactionLogger interface {
	WritePut(key, value string)
	WriteDelete(key string)
	Err() <-chan error

	Run()
	Read() (<-chan Event, <-chan error)
}

// Структура "регистратор транзакций в файл"
type FileTransactionLogger struct {
	events       chan<- Event
	errors       <-chan error
	lastSequence uint64
	file         *os.File
}

// Запись транзакции добавления
func (l *FileTransactionLogger) WritePut(key, value string) {
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

// Запись транзакции удаления
func (l *FileTransactionLogger) WriteDelete(key string) {
	l.events <- Event{EventType: EventDelete, Key: key}
}

// Получение канала ошибок
func (l *FileTransactionLogger) Err() <-chan error {
	return l.errors
}

// Параметры событий
const (
	EventsCapacity   = 16               // Размер буфера
	EventsFileFormat = "%d\t%d\t%s\t%s" // Формат файла
)

// Запуск регистратора
func (l *FileTransactionLogger) Run() {
	events := make(chan Event, EventsCapacity)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		for e := range events {
			l.lastSequence++

			_, err := fmt.Fprintf(l.file, EventsFileFormat+"\n",
				l.lastSequence, e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- err
				return
			}
		}
	}()
}

// Чтение регистратора
func (l *FileTransactionLogger) Read() (<-chan Event, <-chan error) {
	scanner := bufio.NewScanner(l.file)
	events := make(chan Event)
	errors := make(chan error, 1)

	go func() {
		var e Event
		defer close(events)
		defer close(errors)

		for scanner.Scan() {
			line := scanner.Text()

			_, err := fmt.Sscanf(line, EventsFileFormat,
				&e.Sequence, &e.EventType, &e.Key, &e.Value)
			if err != nil {
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
func NewFileTransactionLogger(filename string) (TransactionLogger, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл регистратора транзакций: %w", err)
	}

	return &FileTransactionLogger{file: file}, nil
}

// Регистратор транзакций
var logger TransactionLogger

// Инициализация регистрации транзакций
func initializeTransactionLog() error {
	path := base.Dir("cloud")
	filename := filepath.Join(path, "transaction.log")
	var err error

	logger, err = NewFileTransactionLogger(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать регистратор транзакций: %w", err)
	}

	events, errors := logger.Read()
	e, ok := Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-errors:
		case e, ok = <-events:
			switch e.EventType {
			case EventPut:
				err = Put(e.Key, e.Value)
			case EventDelete:
				err = Delete(e.Key)
			}
		}
	}

	logger.Run()

	return err
}

// Обработчик добавления значения
func handlePut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Обработчик получения значения
func handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := Get(key)
	if errors.Is(err, ErrKeyNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

// Обработчик удаления значения
func handleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println(" \n[ GO CLOUD ]\n ")

	// Новый роутер
	r := mux.NewRouter()

	// Обработчики
	r.HandleFunc("/v1/{key}", handlePut).Methods("PUT")
	r.HandleFunc("/v1/{key}", handleGet).Methods("GET")
	r.HandleFunc("/v1/{key}", handleDelete).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
