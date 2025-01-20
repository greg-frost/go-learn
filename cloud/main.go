package main

import (
	"errors"
	"fmt"
)

// Хранилище пар-значений
var store map[string]string

// Ошибка поиска ключа
var ErrKeyNotFound = errors.New("ключ не найден")

// Добавление значения по ключу
func Put(key, value string) error {
	store[key] = value
	return nil
}

// Получение значения по ключу
func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

// Удаление значения по ключу
func Delete(key string) error {
	delete(store, key)
	return nil
}

func main() {
	fmt.Println(" \n[ GO CLOUD ]\n ")
}
