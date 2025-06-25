package main

import (
	"fmt"
)

// Структура "синглтон"
type singleton struct {
	store map[string]string
}

// Экземпляр синглтона
var instance *singleton

// Получение единственного экземпляра
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{
			store: make(map[string]string),
		}
	}
	return instance
}

// Чтение значения
func (s *singleton) Get(key string) (string, bool) {
	value, ok := s.store[key]
	return value, ok
}

// Запись значения
func (s *singleton) Put(key, value string) {
	s.store[key] = value
}

// Удаление значения
func (s *singleton) Delete(key string) {
	delete(s.store, key)
}

func main() {
	fmt.Println(" \n[ СИНГЛТОН ]\n ")

	// Первый синглтон
	fmt.Println("Первый экземпляр")
	singleton := GetInstance()

	fmt.Println("Запись значения:")
	singleton.Put("key", "value")
	fmt.Println("key -> value")

	fmt.Println("Чтение значения:")
	value, ok := singleton.Get("key")
	fmt.Printf("key = %s (%t)\n", value, ok)
	fmt.Println()

	// Второй синглтон
	fmt.Println("Первый экземпляр")
	other := GetInstance()

	fmt.Println("Чтение значения:")
	value, ok = other.Get("key")
	fmt.Printf("key = %s (%t)\n", value, ok)

	fmt.Println("Удаление значения:")
	other.Delete("key")
	fmt.Println("key -- value")
	fmt.Println()

	// Первый синглтон
	fmt.Println("Первый экземпляр")
	fmt.Println("Чтение значения:")
	value, ok = singleton.Get("key")
	fmt.Printf("key = %q (%t)\n", value, ok)
}
