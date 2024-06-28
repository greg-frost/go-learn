package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Путь
var path = os.Getenv("GOPATH") + "/src/golearn/"

// Обработчик главной страницы
func handleMain(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path+"server/main.go")
}

func main() {
	fmt.Println(" \n[ ФАЙЛОВЫЙ СЕРВЕР ]\n ")

	// Обработчики
	http.HandleFunc("/", handleMain)
	http.Handle(
		"/test/",
		http.StripPrefix(
			"/test/",
			http.FileServer(http.Dir(path+"test")),
		),
	)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
