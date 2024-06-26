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

// Путь и шаблоны
var path = os.Getenv("GOPATH") + "/src/golearn/templates2/tmpl/"
var t = map[string]*template.Template{}

// Инициализация
func init() {
	t["simple"] = parseTemplate("simple.html")
	t["page"] = parseTemplate("base.html", "page.html")
	t["user"] = parseTemplate("base.html", "user.html")
}

// Парсинг шаблона
func parseTemplate(filenames ...string) *template.Template {
	if len(filenames) == 0 {
		return nil
	}
	t := template.New(filenames[0])
	t.Funcs(funcMap)
	for i := 0; i < len(filenames); i++ {
		filenames[i] = path + filenames[i]
	}
	template.Must(t.ParseFiles(filenames...))
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

// Структура "страница"
type Page struct {
	Title   string
	Content string
	Date    time.Time
}

// Структура "пользователь"
type User struct {
	Username string
	Name     string
}

// Обработчик простой страницы
func handleSimple(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		Title:   "HTML-шаблон",
		Content: "Данная страница была сгенерирована в Go!",
		Date:    time.Now(),
	}
	var b bytes.Buffer
	err := t["simple"].Execute(&b, page)
	if err != nil {
		fmt.Fprint(w, "Ошибка выполнения шаблона")
		return
	}
	b.WriteTo(w)
}

// Обработчик составной страницы
func handlePage(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		Title:   "Составная страница",
		Content: "Данная страница была собрана из разных шаблонов, прямо как Франкенштейн...",
	}
	t["page"].ExecuteTemplate(w, "base", page)
}

// Обработчик составной страницы пользователя
func handleUser(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Username: "greg_frost",
		Name:     "Морозов Григорий",
	}
	t["user"].ExecuteTemplate(w, "base", user)
}

func main() {
	fmt.Println(" \n[ HTML-ШАБЛОНЫ ]\n ")

	// Обработчики
	http.HandleFunc("/", handleSimple)
	http.HandleFunc("/page", handlePage)
	http.HandleFunc("/user", handleUser)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
