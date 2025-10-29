package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Обработчик
func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println(" \n[ ПОСЕЩЕНИЕ МАРШРУТОВ ]\n ")

	// Роутер
	router := mux.NewRouter()

	// GET-обработчики
	getSub := router.Methods(http.MethodGet).Subrouter()
	getSub.HandleFunc("/albums", handler)
	getSub.HandleFunc("/albums/{id:\\d+}", handler)

	// POST-обработчики
	postSub := router.Methods(http.MethodPost).Subrouter()
	postSub.HandleFunc("/albums", handler)

	// PUT-обработчики
	putSub := router.Methods(http.MethodPut).Subrouter()
	putSub.HandleFunc("/albums/{id:[0-9]+}", handler)

	// DELETE-обработчики
	deleteSub := router.Methods(http.MethodDelete).Subrouter()
	deleteSub.HandleFunc("/albums/{id:\\d+}", handler)

	// Посещение всех маршрутов
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		} else {
			// Подмаршрутизатор
			fmt.Println("SUBROUTER") // Печать
			// return nil // Пропуск
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queryTemplates, err := route.GetQueriesTemplates()
		if err == nil && len(queryTemplates) > 0 {
			fmt.Println("Queries templates:", strings.Join(queryTemplates, ","))
		}
		queryRegexps, err := route.GetQueriesRegexp()
		if err == nil && len(queryRegexps) > 0 {
			fmt.Println("Queries regexps:", strings.Join(queryRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil && len(methods) > 0 {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
