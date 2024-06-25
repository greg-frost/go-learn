package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

// Структура "страница"
type Page struct {
	Title   string
	Content string
	Date    time.Time
}

// Путь и шаблон
var path = os.Getenv("GOPATH") + "/src/golearn/templates2/"
var t *template.Template

// Инициализация
func init() {
	t = parseTemplate("simple.html")
}

// Парсинг шаблона
func parseTemplate(filename string) *template.Template {
	t := template.New(filename)
	t.Funcs(funcMap)
	template.Must(t.ParseFiles(path + filename))
	return t
}

// Список функций шаблона
var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

// Форматирование даты
func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

// Обработчик шаблона страницы
func handlePage(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		Title:   "HTML-шаблон",
		Content: "Данная страница была сгенерирована в Go!",
		Date:    time.Now(),
	}
	var b bytes.Buffer
	err := t.Execute(&b, page)
	if err != nil {
		fmt.Fprint(w, "Ошибка выполнения шаблона")
		return
	}
	b.WriteTo(w)
}

func main() {
	fmt.Println(" \n[ HTML-ШАБЛОНЫ ]\n ")

	// Обработчик
	http.HandleFunc("/", handlePage)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
