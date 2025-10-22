package frontend

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"go-learn/cloud/core"

	"github.com/gorilla/mux"
)

// Структура "REST-фронтэнд"
type restFrontEnd struct {
	store *core.KeyValueStore
}

// Конструктор REST-фронтэнда
func newRestFrontEnd() *restFrontEnd {
	return new(restFrontEnd)
}

// Получение значения
func (f *restFrontEnd) handleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := f.store.Get(key)
	if errors.Is(err, core.ErrKeyNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

// Добавление значения
func (f *restFrontEnd) handlePut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = f.store.Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Удаление значения
func (f *restFrontEnd) handleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	err := f.store.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Запуск REST-сервера
func (f *restFrontEnd) Start(kvs *core.KeyValueStore) error {
	f.store = kvs

	// Новый роутер
	r := mux.NewRouter()

	// Обработчики
	r.HandleFunc("/v1/{key}", f.handleGet).Methods("GET")
	r.HandleFunc("/v1/{key}", f.handlePut).Methods("PUT")
	r.HandleFunc("/v1/{key}", f.handleDelete).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")

	// HTTP
	err := http.ListenAndServe("localhost:8080", r)

	// HTTPS
	// path := base.Dir("cloud")
	// err := http.ListenAndServeTLS("localhost:8080",
	// 	filepath.Join(path, "cert.pem"),
	// 	filepath.Join(path, "key.pem"),
	// 	r)

	return err
}
