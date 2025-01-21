package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Хранилище пар ключ/значение
var store = make(map[string]string)

// Ошибка поиска ключа
var ErrKeyNotFound = errors.New("ключ не найден")

// Добавление значения по ключу
func Put(key, value string) error {
	store[key] = value
	return nil
}

// Получение значения по ключу
func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

// Удаление значения по ключу
func Delete(key string) error {
	delete(store, key)
	return nil
}

// Обработчик добавления значения
func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	fmt.Println(" \n[ GO CLOUD ]\n ")

	// Новый роутер
	r := mux.NewRouter()

	// Обработчики
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
