package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Структура "альбом"
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Данные альбомов
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Получение всех альбомов
func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

// Получение альбома по ID
func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	for _, a := range albums {
		if a.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(a)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `{"message": "альбом не найден"}`)
}

// Создание альбома
func postAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newAlbum album
	if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(albums)
}

func main() {
	fmt.Println(" \n[ GORILLA MUX ]\n ")

	// Роутер
	router := mux.NewRouter().StrictSlash(true)

	// Обработчики
	router.HandleFunc("/albums", getAlbums).Methods("GET")
	router.HandleFunc("/albums/{id}", getAlbumByID).Methods("GET")
	router.HandleFunc("/albums", postAlbum).Methods("POST")

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
