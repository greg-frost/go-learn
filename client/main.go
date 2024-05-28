package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println(" \n[ HTTP-КЛИЕНТ ]\n ")

	// Запрос
	resp, err := http.Get("https://go.dev")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Статус ответа
	fmt.Println("Статус ответа:", resp.Status)
	fmt.Println()

	// Чтение построчно
	fmt.Println("Первые строки:")
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 3; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("...")
	fmt.Println()

	// Чтение полностью
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Осталось символов:", len(body))
}
