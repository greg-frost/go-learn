package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-learn/graphql2/auth"
	"go-learn/graphql2/graph"
	"go-learn/graphql2/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
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

	// Конфигурация
	c := graph.Config{
		Resolvers: &graph.Resolver{},
	}

	// Сложность запросов
	countComplexity := func(childComplexity int, genre *model.Genre, limit, offset *int) int {
		if limit == nil {
			return childComplexity
		}
		return *limit * childComplexity
	}
	c.Complexity.Query.Videos = countComplexity
	c.Complexity.Video.Related = countComplexity

	// Роутер
	router := chi.NewRouter()

	// Промежуточный слой (аутентификация)
	router.Use(auth.Middleware(nil))

	// Сервер
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	// Web-socket
	srv.AddTransport(&transport.Websocket{})

	// Лимит сложности запросов
	srv.Use(extension.FixedComplexityLimit(1000))

	// Обработчики
	router.Handle("/", playground.Handler("GraphQL-сервер", "/query"))
	router.Handle("/query", srv)

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Printf("(на http://localhost:%s)\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
