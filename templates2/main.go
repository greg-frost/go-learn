package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

// Структура "страница"
type Page struct {
	Title   string
	Content string
}

// Обработчик шаблона страницы
func handlePage(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		Title:   "HTML-шаблон",
		Content: "Данная страница была сгенерирована в Go!",
	}

	path := os.Getenv("GOPATH") + "/src/golearn/templates2/"
	t := template.Must(template.ParseFiles(path + "simple.html"))
	t.Execute(w, page)
}

func main() {
	fmt.Println(" \n[ HTML-ШАБЛОНЫ ]\n ")

	// Обработчик
	http.HandleFunc("/", handlePage)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
