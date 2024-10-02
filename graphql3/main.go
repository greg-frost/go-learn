package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	fmt.Println(" \n[ GRAPHQL 3 (GRAPHQL-GO) ]\n ")

	/* Простая схема */

	fmt.Println("Простая схема:")
	fmt.Println()

	// Поля
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return `Привет, Мир!`, nil
			},
		},
	}

	// Объект Query
	Query := graphql.ObjectConfig{Name: "Query", Fields: fields}
	// Конфигурация
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(Query)}
	// Схема
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("не удалось создать новую схему: %v", err)
	}

	// Запрос
	query := `{ hello }`
	fmt.Printf("Запрос:\n%s\n\n", query)

	// Выполнение
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("не удалось выполнить операцию: %+v", r.Errors)
	}

	// Ответ
	response, _ := json.MarshalIndent(r, "", "   ")
	fmt.Printf("Ответ:\n%s\n\n", response)

	/* Использование контекста */

	fmt.Println("Использование контекста:")
	fmt.Println()

	// Имя аргумента
	const arg = "key"

	// Функция получения поля
	fieldFromContext := func(p graphql.ResolveParams) (interface{}, error) {
		return p.Context.Value(p.Args[arg]), nil
	}

	// Объект Query, конфигурация и схема
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"value": &graphql.Field{
					Type: graphql.String,
					Args: graphql.FieldConfigArgument{
						arg: &graphql.ArgumentConfig{Type: graphql.String},
					},
					Resolve: fieldFromContext,
				},
			},
		}),
	})

	// Запрос
	query = `{ value(` + arg + `: "Username") }`
	fmt.Printf("Запрос:\n%s\n\n", query)

	// Выполнение
	r = graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       context.WithValue(context.TODO(), "Username", "Greg Frost"),
	})
	if len(r.Errors) > 0 {
		log.Fatalf("не удалось выполнить операцию: %+v", r.Errors)
	}

	// Ответ
	response, _ = json.MarshalIndent(r, "", "   ")
	fmt.Printf("Ответ:\n%s\n\n", response)
}
