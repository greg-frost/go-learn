package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"go-learn/base"
	// fs "github.com/Masterminds/go-fileserver"
)

// Путь
var path = base.Dir("fileserver/..")

// Обработчик главной страницы
func handleMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(path, "server", "main.go"))
}

func main() {
	fmt.Println(" \n[ ФАЙЛОВЫЙ СЕРВЕР ]\n ")

	// Обработчики
	http.HandleFunc("/", handleMain)
	http.Handle(
		"/testing/",
		http.StripPrefix(
			"/testing/",
			http.FileServer(http.Dir(filepath.Join(path, "testing"))),
		),
	)

	// Пользовательский файл-сервер
	// с настраиваемой страницей ошибок
	// fs.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// 	fmt.Fprintln(w, "Ошибка 404 - Страница не найдена!")
	// }

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	// log.Fatal(http.ListenAndServe("localhost:8080",
	// 	fs.FileServer(http.Dir(filepath.Join(path, "testing"))),
	// ))
}
