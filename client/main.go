package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Структура "клиент с контекстом"
type ClientContext struct {
	http.Client
}

// Конструктор клиента с контекстом
func NewClientContext() *ClientContext {
	return new(ClientContext)
}

// Get-запрос клиента с контекстом
func (c *ClientContext) GetContext(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func main() {
	fmt.Println(" \n[ HTTP-КЛИЕНТ ]\n ")

	// GET-запрос
	addr := "https://go.dev"
	fmt.Println("GET:", addr)
	get, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer get.Body.Close()

	// Статус ответа
	fmt.Println("Статус ответа:", get.Status)
	fmt.Println()

	// Чтение построчно
	fmt.Println("Первые строки:")
	var count int
	scanner := bufio.NewScanner(get.Body)
	for i := 0; scanner.Scan() && i < 3; i++ {
		line := scanner.Text()
		count += len(line)
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("...")

	// Чтение полностью
	body, err := io.ReadAll(get.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("(осталось символов: %d)\n", len(body)-count)
	fmt.Println()

	// POST-запрос
	fmt.Println("POST:", addr)
	data := strings.NewReader("payload")
	req, err := http.NewRequest("POST", addr, data)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Статус ответа:", res.Status)
	fmt.Println()

	// Свой клиент
	fmt.Println("HEAD:", addr)
	client := &http.Client{Timeout: 200 * time.Millisecond}
	res, err = client.Head(addr)
	if err != nil {
		fmt.Println("Ошибка: таймаут!")
	} else {
		fmt.Println("Статус ответа:", res.Status)
	}
	fmt.Println()

	// Контекст
	fmt.Println("CONTEXT:", addr)
	clientContext := NewClientContext()
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	res, err = clientContext.GetContext(ctx, addr)
	if err != nil && errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("Ошибка: таймаут!")
	} else {
		fmt.Println("Статус ответа:", res.Status)
	}
}
