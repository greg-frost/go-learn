package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
	scanner := bufio.NewScanner(get.Body)
	for i := 0; scanner.Scan() && i < 3; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("...")
	fmt.Println()

	// Чтение полностью
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Осталось символов:", len(body))

	// POST-запрос

	fmt.Println()
	fmt.Println("POST:", addr)

	data := strings.NewReader("payload data")
	req, err := http.NewRequest("POST", addr, data)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Статус ответа:", res.Status)
}
