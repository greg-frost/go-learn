package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	fmt.Println(" \n[ ПОВТОРНОЕ ИСПОЛЬЗОВАНИЕ ]\n ")

	// Транспорт
	tr := &http.Transport{
		// DisableKeepAlives: true,
		// DisableCompression: true,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
	}

	// Клиент
	client := &http.Client{Transport: tr}

	times := 25
	fmt.Println("Число повторов:", times)
	fmt.Println()

	// Запрос 1
	var recieved int
	start := time.Now()
	for i := 0; i < times; i++ {
		r, err := client.Get("https://go.dev")
		if err != nil {
			log.Fatal(err)
		}
		o, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		r.Body.Close()
		recieved += len(o)
	}
	fmt.Println("Запрос 1:")
	fmt.Println("Получено", recieved, "байт")
	fmt.Println("Заняло", time.Since(start))
	fmt.Println()

	// Запрос 2
	recieved = 0
	start = time.Now()
	for i := 0; i < times; i++ {
		r, err := client.Get("http://example.com")
		if err != nil {
			log.Fatal(err)
		}
		o, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		r.Body.Close()
		recieved += len(o)
	}
	fmt.Println("Запрос 1:")
	fmt.Println("Получено", recieved, "байт")
	fmt.Println("Заняло", time.Since(start))
}
