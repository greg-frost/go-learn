package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"go-learn/base"

	"github.com/graphql-go/graphql"
)

// Имя аргумента
const arg = "key"

// Получение поля из контекста
func fieldFromContext(p graphql.ResolveParams) (interface{}, error) {
	return p.Context.Value(p.Args[arg]), nil
}

// Список пользователей
var users []*User

// Структура "пользователь"
type User struct {
	ID     int
	Name   string
	Active bool
}

// Получение пользователей
func ListUsers() ([]*User, error) {
	return users, nil
}

// Получение пользователя по ID
func GetUsersByID(id int) (*User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("пользователь не найден")
}

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

	/* Режим сервера */

	fmt.Println("Режим сервера:")
	fmt.Println()

	path := base.Dir("graphql3")
	jsonUsers, err := ioutil.ReadFile(filepath.Join(path, "users.json"))
	if err != nil {
		log.Fatalf("не удалось загрузить файл с пользователями: %v", err)

	}
	err = json.Unmarshal(jsonUsers, &users)
	if err != nil {
		log.Fatalf("не удалось обработать json-файл: %v", err)
	}

	// Поля
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.Name, nil
					}
					return nil, nil
				},
			},
			"active": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.Active, nil
					}
					return nil, nil
				},
			},
		},
	})

	// Схема
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(userType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return ListUsers()
				},
			},
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetUsersByID(p.Args["id"].(int))
				},
			},
		},
	})

}
