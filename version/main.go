package main

import (
	"encoding/json"
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

// Структура "сообщение версии 2"
type MessageV2 struct {
	Version int    `json:"ver"`
	Message string `json:"msg"`
}

// Обработчик версии 1 по URL
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

// Обработчик версии по Content-Type
func handleByContentType(w http.ResponseWriter, r *http.Request) {
	var err error
	var b []byte
	var contentType string
	acceptType := r.Header.Get("Accept")

	switch acceptType {
	case "application/vnd.myapi.json; version=2.0":
		data := MessageV2{
			Version: 2,
			Message: "Сообщение приложения",
		}
		b, err = json.Marshal(data)
		contentType = "application/vnd.myapi.json; version=2.0"
	case "application/vnd.myapi.json; version=1.0":
		fallthrough
	default:
		data := MessageV1{
			Info: "Сообщение приложения, V1",
		}
		b, err = json.Marshal(data)
		contentType = "application/vnd.myapi.json; version=1.0"
	}
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", contentType)
	fmt.Fprint(w, string(b))
}

// Получение страницы
func getPage(url, acceptType string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", acceptType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	contentType := res.Header.Get("Content-Type")
	if acceptType != "" &&
		acceptType != contentType {
		return "", fmt.Errorf("неизвестный тип содержимого - %q", contentType)
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
		http.HandleFunc("/endpoint", handleByContentType)

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

	res, err = getPage("http://localhost:8080/endpoint", "application/vnd.myapi.json; version=1.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("V1 по Content-Type:", res)

	res, err = getPage("http://localhost:8080/endpoint", "application/vnd.myapi.json; version=2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("V2 по Content-Type:", res)
}
