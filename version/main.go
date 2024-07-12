package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Структура "сообщение версии 1"
type MessageV1 struct {
	Info string `json:"info"`
}

// Обработчик версии по URL
func handleByUrl(w http.ResponseWriter, r *http.Request) {
	data := MessageV1{
		Info: "Сообщение приложения, V1",
	}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

// Получение страницы
func getPage(url, contentType string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", contentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if contentType != "" &&
		res.Header.Get("Content-Type") != contentType {
		return "", errors.New("неизвестный тип содержимого")
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(b), nil
}

func main() {
	fmt.Println(" \n[ ВЕРСИОНИРОВАНИЕ ]\n ")

	/* Сервер */

	fmt.Println("Сервер:")
	go func() {
		http.HandleFunc("/api/v1/endpoint", handleByUrl)

		fmt.Println("Ожидаю обновлений...")
		fmt.Println("(на http://localhost:8080)")
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}()

	time.Sleep(250 * time.Millisecond)
	fmt.Println()

	/* Клиенты */

	fmt.Println("Клиенты:")
	res, err := getPage("http://localhost:8080/api/v1/endpoint", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("V1 по URL:", res)
}
