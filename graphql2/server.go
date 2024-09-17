package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golearn/graphql2/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

// Порт по умолчанию
const defaultPort = "8080"

func main() {
	fmt.Println(" \n[ GRAPHQL 2 (GQLGEN) ]\n ")

	// Порт
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Сервер
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))

	// Web-socket
	srv.AddTransport(&transport.Websocket{})

	// Обработчики
	http.Handle("/", playground.Handler("GraphQL-сервер", "/query"))
	http.Handle("/query", srv)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Printf("(на http://localhost:%s)\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
