package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Хранилище пар ключ/значение
var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

// Ошибка поиска ключа
var ErrKeyNotFound = errors.New("ключ не найден")

// Добавление значения по ключу
func Put(key, value string) error {
	store.Lock()
	defer store.Unlock()

	store.m[key] = value

	return nil
}

// Получение значения по ключу
func Get(key string) (string, error) {
	store.RLock()
	defer store.RUnlock()

	value, ok := store.m[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return value, nil
}

// Удаление значения по ключу
func Delete(key string) error {
	store.Lock()
	defer store.Unlock()

	delete(store.m, key)

	return nil
}

// Обработчик добавления значения
func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Обработчик получения значения
func keyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := Get(key)
	if errors.Is(err, ErrKeyNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

// Обработчик удаления значения
func keyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println(" \n[ GO CLOUD ]\n ")

	// Новый роутер
	r := mux.NewRouter()

	// Обработчики
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	r.HandleFunc("/v1/{key}", keyValueGetHandler).Methods("GET")
	r.HandleFunc("/v1/{key}", keyValueDeleteHandler).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
