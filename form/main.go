package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"go-learn/base"
)

// Путь и шаблон
var path = base.Dir("form")
var t = template.Must(template.ParseFiles(filepath.Join(path, "form.html")))

// Структура "страница"
type Page struct {
	Title string
	Name  string
	Langs []string
}

// Обработчик формы
func handleForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := r.FormValue("name")
	if name == "" {
		name = r.PostFormValue("name")
	}

	var langs []string
	for _, v := range r.Form["langs"] {
		langs = append(langs, v)
	}

	page := &Page{
		Title: "HTML-форма",
		Name:  name,
		Langs: langs,
	}
	t.Execute(w, page)
}

func main() {
	fmt.Println(" \n[ HTML-ФОРМА ]\n ")

	// Обработчик
	http.HandleFunc("/", handleForm)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
