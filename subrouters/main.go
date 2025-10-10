package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Получение всех альбомов
func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Получение всех альбомов")
}

// Получение альбома по ID
func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Получение альбома по ID (%s)", id)
}

// Создание альбома
func postAlbum(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Создание альбома")
}

// Изменение альбома по ID
func putAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Изменение альбома по ID (%s)", id)
}

// Удаление альбома по ID
func deleteAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Удаление альбома по ID (%s)", id)
}

func main() {
	fmt.Println(" \n[ ПОДМАРШТУРИЗАТОРЫ ]\n ")

	// Роутер
	router := mux.NewRouter()

	// GET-обработчики
	getSub := router.Methods(http.MethodGet).Subrouter()
	getSub.HandleFunc("/albums", getAlbums)
	getSub.HandleFunc("/albums/{id}", getAlbumByID)

	// POST-обработчики
	postSub := router.Methods(http.MethodPost).Subrouter()
	postSub.HandleFunc("/albums", postAlbum)

	// PUT-обработчики
	putSub := router.Methods(http.MethodPut).Subrouter()
	putSub.HandleFunc("/albums/{id}", putAlbumByID)

	// DELETE-обработчики
	deleteSub := router.Methods(http.MethodDelete).Subrouter()
	deleteSub.HandleFunc("/albums/{id}", deleteAlbumByID)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
