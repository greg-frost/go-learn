package core

import (
	"errors"
	"sync"
)

// Структура "хранилище пар ключ/значение"
type KeyValueStore struct {
	m        map[string]string
	mu       sync.RWMutex
	transact TransactionLogger
}

// Интерфейс "регистратор транзакций"
type TransactionLogger interface {
	WritePut(key, value string)
	WriteDelete(key string)
	Err() <-chan error

	Run()
	Read() (<-chan Event, <-chan error)
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

// Ошибка поиска ключа
var ErrKeyNotFound = errors.New("ключ не найден")

// Конструктор хранилища
func NewKeyValueStore(tl TransactionLogger) *KeyValueStore {
	return &KeyValueStore{
		m:        make(map[string]string),
		transact: tl,
	}
}

// Получение значения по ключу
func (store *KeyValueStore) Get(key string) (string, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	value, ok := store.m[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

// Добавление значения по ключу
func (store *KeyValueStore) Put(key, value string) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.m[key] = value
	store.transact.WritePut(key, value)

	return nil
}

// Удаление значения по ключу
func (store *KeyValueStore) Delete(key string) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	delete(store.m, key)
	store.transact.WriteDelete(key)

	return nil
}
