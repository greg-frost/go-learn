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
	author, err := composites.NewAuthorComposite()
	if err != nil {
		log.Fatal(err)
	}
	author.Handler.Register(router)

	// Книги
	book, err := composites.NewBookComposite(author)
	if err != nil {
		log.Fatal(err)
	}
	book.Handler.Register(router)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
