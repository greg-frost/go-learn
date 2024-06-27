package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// Путь и шаблоны
var path = os.Getenv("GOPATH") + "/src/golearn/templates2/tmpl/"
var t = map[string]*template.Template{}
var qc template.HTML

// Структура "страница"
type Page struct {
	Title   string
	Content template.HTML
	Date    time.Time
}

// Структура "пользователь"
type User struct {
	Username string
	Name     string
}

// Структура "цитата"
type Quote struct {
	Quote  string
	Person string
}

// Список функций шаблона
var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

// Инициализация
func init() {
	t["simple"] = parseTemplate("simple.html")
	t["page"] = parseTemplate("base.html", "page.html", "quote.html")
	t["user"] = parseTemplate("base.html", "user.html")
	prepareQuote()
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

// Подготовка цитаты
func prepareQuote() {
	quote := &Quote{
		Quote:  "Данная страница была собрана из разных шаблонов, прямо как Франкенштейн...",
		Person: "Морозов Григорий",
	}
	var b bytes.Buffer
	t["page"].ExecuteTemplate(&b, "quote.html", quote)
	qc = template.HTML(b.String())
}

// Форматирование даты
func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
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
		Content: qc,
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
