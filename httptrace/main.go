package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	fmt.Println(" \n[ HTTP-ТРАССИРОВКА ]\n ")

	// Клиент и запрос
	addr := "https://go.dev"
	client := &http.Client{}
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Трассировка
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("Получен первый байт ответа")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Соединение: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Printf("Начало соединения: %v (%v)\n", addr, network)
		},
		ConnectDone: func(network, addr string, err error) {
			if err != nil {
				fmt.Println("Ошибка соединения:", err)
				return
			}
			fmt.Printf("Завершение соединения: %v (%v)\n", addr, network)
		},
		WroteHeaders: func() {
			fmt.Println("Добавлены заголовки")
		},
	}

	// Отслеживание запроса
	fmt.Println("Отслеживание запроса")
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err = http.DefaultTransport.RoundTrip(req)
	fmt.Println()

	// Запрос
	fmt.Println("Выполнение запроса")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Статус ответа:", res.Status)
}
