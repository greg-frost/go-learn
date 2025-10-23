package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Структура "ошибка"
type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

// Структура "обертка ошибки"
type ErrorWrap struct {
	Err Error `json:"error"`
}

// Формирование JSON-ошибки
func JSONError(w http.ResponseWriter, e Error) {
	data := ErrorWrap{e}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

// Обработчик страницы с ошибкой
func handleError(w http.ResponseWriter, r *http.Request) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     40300,
		Message:  "Доступ запрещен навсегда",
	}
	JSONError(w, e)
}

// Форматирование ошибки
func (e Error) Error() string {
	msg := "HTTP: %d, Код: %d, Сообщение: %s"
	return fmt.Sprintf(msg, e.HTTPCode, e.Code, e.Message)
}

// Получение страницы
func get(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if res.Header.Get("Content-Type") != "application/json" {
			msg := "Неизвестная ошибка, HTTP-статус: %s"
			return res, fmt.Errorf(msg, res.Status)
		}

		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		var data ErrorWrap
		err = json.Unmarshal(b, &data)
		if err != nil {
			msg := "Не удалось прочитать JSON: %s, HTTP-статус: %s"
			return res, fmt.Errorf(msg, err, res.Status)
		}
		data.Err.HTTPCode = res.StatusCode

		return res, data.Err
	}
	return res, nil
}

func main() {
	fmt.Println(" \n[ HTTP-ОШИБКИ ]\n ")

	/* Сервер */

	fmt.Println("Сервер:")
	go func() {
		http.HandleFunc("/", handleError)

		fmt.Println("Ожидаю соединений...")
		fmt.Println("(на http://localhost:8080)")
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}()

	time.Sleep(250 * time.Millisecond)
	fmt.Println()

	/* Клиент */

	fmt.Println("Клиент:")
	res, err := get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(b))
}
