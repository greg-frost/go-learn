package main

import (
	"fmt"
	"log"
	"net/http"

	"go-learn/clean/internal/composites"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println(" \n[ ЧИСТАЯ АРХИТЕКТУРА ]\n ")

	// Роутер
	router := httprouter.New()

	// Авторы
	authors, err := composites.NewAuthorComposite()
	if err != nil {
		log.Fatalf("композит авторов: %v", err)
	}
	authors.Handler.Register(router)

	// Книги
	books, err := composites.NewBookComposite(authors)
	if err != nil {
		log.Fatalf("композит книг: %v", err)
	}
	books.Handler.Register(router)

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
