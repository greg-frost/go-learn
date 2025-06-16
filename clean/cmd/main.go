package main

import (
	"fmt"
	"log"
	"net/http"

	ahandler "go-learn/clean/internal/adapters/api/author"
	bhandler "go-learn/clean/internal/adapters/api/book"
	astorage "go-learn/clean/internal/adapters/db/author"
	bstorage "go-learn/clean/internal/adapters/db/book"
	aservice "go-learn/clean/internal/domain/author"
	bservice "go-learn/clean/internal/domain/book"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println(" \n[ ЧИСТАЯ АРХИТЕКТУРА ]\n ")

	// Роутер
	router := httprouter.New()

	// Авторы
	authorStorage := astorage.NewStorage()
	authorService := aservice.NewService(authorStorage)
	authorHandler := ahandler.NewHandler(authorService)
	authorHandler.Register(router)

	// Книги
	bookStorage := bstorage.NewStorage()
	bookService := bservice.NewService(bookStorage, authorService)
	bookHandler := bhandler.NewHandler(bookService)
	bookHandler.Register(router)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
