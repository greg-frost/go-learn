package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

// Кэш
var cache *lru.Cache

// Инициализация
func init() {
	var err error
	cache, err = lru.NewWithEvict(2,
		func(key, value interface{}) {
			fmt.Printf("[ вытеснено: key=%v value=%v ]\n", key, value)
		},
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(" \n[ LRU-кэш ]\n ")

	// Добавление
	fmt.Println("Добавление значений")
	cache.Add(1, "a")
	fmt.Println("1 -> a")
	cache.Add(2, "b")
	fmt.Println("2 -> b")
	fmt.Println()

	// Получение
	fmt.Println("Получение с обновлением")
	value, ok := cache.Get(1)
	fmt.Println("1:", value, ok)
	fmt.Println()

	// Просмотр
	fmt.Println("Просмотр без обновления")
	fmt.Println("2:", cache.Contains(2))
	fmt.Println()

	// Добавление с вытеснением
	fmt.Println("Добавление с вытеснением")
	cache.Add(3, "c")
	fmt.Println("3 -> c")
	fmt.Println()

	// Получение вытесненного значения
	fmt.Println("Получение вытесненного значения")
	value, ok = cache.Get(2)
	fmt.Println("2:", value, ok)
}
