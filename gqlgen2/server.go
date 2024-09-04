package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golearn/gqlgen2/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	fmt.Println(" \n[ GQLGEN 2 ]\n ")

	// Порт
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Сервер
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))

	// Обработчики
	http.Handle("/", playground.Handler("GraphQL-сервер", "/query"))
	http.Handle("/query", srv)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Printf("(на http://localhost:%s)\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
