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
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("не удалось создать новую схему: %v", err)
	}

	// Запрос
	query := "{\n\thello\n}"
	fmt.Printf("Запрос:\n%s\n\n", query)

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("не удалось выполнить операцию: %+v", r.Errors)
	}

	// Ответ
	response, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("Ответ:\n%s\n", response)
}
