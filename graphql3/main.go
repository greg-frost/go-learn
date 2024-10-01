package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	fmt.Println(" \n[ GRAPHQL 3 (GRAPHQL-GO) ]\n ")

	// Схема
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
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(Query)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("не удалось создать новую схему: %v", err)
	}

	// Запрос
	query := "{\n    hello\n}"
	fmt.Printf("Запрос:\n%s\n\n", query)

	// Выполнение
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("не удалось выполнить операцию: %+v", r.Errors)
	}

	// Ответ
	response, _ := json.MarshalIndent(r, "", "    ")
	fmt.Printf("Ответ:\n%s\n", response)
}
