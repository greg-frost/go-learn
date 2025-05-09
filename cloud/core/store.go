package core

import (
	"errors"
	"log"
	"sync"
)

// Структура "хранилище пар ключ/значение"
type KeyValueStore struct {
	m        map[string]string
	mu       sync.RWMutex
	transact TransactionLogger
}

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

// Восстановление хранилища из транзакций
func (store *KeyValueStore) Restore() error {
	events, errors := store.transact.Read()
	e, ok := Event{}, true
	var count int
	var err error

	for ok && err == nil {
		select {
		case err, ok = <-errors:
		case e, ok = <-events:
			switch e.EventType {
			case EventPut:
				err = store.Put(e.Key, e.Value)
				count++
			case EventDelete:
				err = store.Delete(e.Key)
				count++
			}
		}
	}

	log.Printf("Восстановлено событий: %d", count)

	store.transact.Run()

	go func() {
		for err := range store.transact.Err() {
			log.Print(err)
		}
	}()

	return err
}
