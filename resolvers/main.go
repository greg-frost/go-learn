package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"regexp"
	"strings"
)

// Структура "обработчик пути"
type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

// Получение нового обработчика пути
func newPathResolver() *pathResolver {
	return &pathResolver{
		handlers: make(map[string]http.HandlerFunc),
	}
}

// Добавление обработчика пути
func (pr *pathResolver) Add(path string, handler http.HandlerFunc) {
	pr.handlers[path] = handler
}

// Обработка HTTP-запроса по пути
func (pr *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path
	for pattern, handlerFunc := range pr.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(w, r)
			return
		} else if err != nil {
			fmt.Fprint(w, err)
		}
	}
	http.NotFound(w, r)
}

// Структура "обработчик выражения"
type regexpResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

// Получение нового обработчика выражения
func newRegexpResolver() *regexpResolver {
	return &regexpResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

// Добавление обработчика выражения
func (rr *regexpResolver) Add(reg string, handler http.HandlerFunc) {
	rr.handlers[reg] = handler
	cache := regexp.MustCompile(reg)
	rr.cache[reg] = cache
}

// Обработка HTTP-запроса по выражению
func (rr *regexpResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path
	for pattern, handlerFunc := range rr.handlers {
		if rr.cache[pattern].MatchString(check) {
			handlerFunc(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

// Приветствие
func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Мир"
	}
	fmt.Fprintf(w, "Привет, %s!", name)
}

// Прощание
func goodbye(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	var name string
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Мир"
	}
	fmt.Fprintf(w, "Прощай, %s...", name)
}

func main() {
	fmt.Println(" \n[ HTTP-ОБРАБОТЧИКИ ]\n ")

	// Обработчик пути
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)

	// Обработчик выражений
	rr := newRegexpResolver()
	rr.Add("GET /hello/?", hello)
	rr.Add("(GET|HEAD) /goodbye(/[A-Za-zА-Яа-яё0-9]*)?", goodbye)

	// Выбор типа обработчика
	useRegexpResolver := flag.Bool("regexp", false, "использовать обработчик регулярных выражений")
	flag.Parse()
	var handler http.Handler = pr
	if *useRegexpResolver {
		fmt.Println("Используется обработчик выражений")
		handler = rr
	} else {
		fmt.Println("Используется обработчик пути")
	}
	fmt.Println()

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
