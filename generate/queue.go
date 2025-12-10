package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"go-learn/base"
)

// Путь и шаблон
var path = base.Dir("generate")
var tpl = `package {{.Package}}

// Структура "очередь"
type {{.MyType}}Queue struct {
	q []{{.MyType}}
}

// Конструктор очереди
func New{{.MyType}}Queue() *{{.MyType}}Queue {
	return &{{.MyType}}Queue{
		q: []{{.MyType}}{},
	}
}

// Вставка значения в конец очереди
func (o *{{.MyType}}Queue) Insert(v {{.MyType}}) {
	o.q = append(o.q, v)
}

// Получение значения из начала очереди
func (o *{{.MyType}}Queue) Remove() {{.MyType}} {
	if len(o.q) == 0 {
		panic("Пусто!")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}`

func main() {
	fmt.Println(" \n[ КОДОГЕНЕРАЦИЯ ]\n ")

	// Парсинг шаблона
	t := template.Must(template.New("queue").Parse(tpl))
	var count int

	// Обработка параметров
	for i := 1; i < len(os.Args); i++ {
		// Создание файла
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(filepath.Join(path, dest))
		if err != nil {
			log.Printf("Не удалось создать %s: %s (пропуск)", dest, err)
			continue
		}

		// Определение имени пакета
		packageName := os.Getenv("GOPACKAGE")
		if packageName == "" {
			packageName = "main"
		}

		// Параметры и выполнение шаблона
		vals := map[string]string{
			"MyType":  os.Args[i],
			"Package": packageName,
		}
		t.Execute(file, vals)
		count++

		file.Close()
	}

	fmt.Println("Создано файлов с кодом:", count)
}
