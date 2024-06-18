package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println(" \n[ REDIS ]\n ")

	// Создание клиента
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()

	// Подключение (пинг)
	ping, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PING:", ping)

	/* Простое значение */

	// Запись
	err = client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// Чтение
	value, err := client.Get(ctx, "key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("key = %v\n", value)

	/* Составное значение */

	// Структура
	type Composite struct {
		Int    int    `json:"int"`
		String string `json:"string"`
		Bool   bool   `json:"bool"`
	}

	// Маршаллизация
	jsonBytes, err := json.Marshal(Composite{
		Int:    1,
		String: "str",
		Bool:   true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Запись
	err = client.Set(ctx, "composite", jsonBytes, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// Чтение
	value, err = client.Get(ctx, "composite").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("composite = %v\n", value)
}
