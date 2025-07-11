package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

// Создание и получение альбомов
func processAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/* Создание альбома */

	if r.Method == "POST" {
		var newAlbum album
		if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
			return
		}
		albums = append(albums, newAlbum)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(albums)
		return
	}

	/* Получение альбома по ID */

	var id string
	id = strings.Split(r.URL.Path, "/")[2]
	if id == "" {
		id = r.URL.Query().Get("id")
	}

	// Go 1.22
	// id := r.PathValue("id")

	if id != "" {
		for _, a := range albums {
			if a.ID == id {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(a)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"message": "альбом не найден"}`)
		return
	}

	/* Получение всех альбомов */

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

func main() {
	fmt.Println(" \n[ REST 1 (GO DOC) ]\n ")

	// Обработчик
	http.HandleFunc("/albums/", processAlbums)

	// Go 1.22
	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /albums", getAlbums)
	// mux.HandleFunc("GET /albums/{id}", getAlbumById)
	// mux.HandleFunc("POST /albums", createAlbum)
	// http.ListenAndServe("localhost:8080", mux)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
