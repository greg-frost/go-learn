package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Println(" \n[ ТЕКСТОВЫЕ ШАБЛОНЫ ]\n ")

	// Создание и парсинг
	t1 := template.New("t1")
	t1, err := t1.Parse("Значение: {{.}}\n")
	if err != nil {
		log.Fatal(err)
	}

	// Обязательное создание
	t1 = template.Must(t1.Parse("Значение: {{.}}\n"))

	// Исполнение
	t1.Execute(os.Stdout, 100)
	t1.Execute(os.Stdout, "текст")
	t1.Execute(os.Stdout, []string{"Go", "PHP", "JavaScript"})

	// Функция для создания шаблонов
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Передача структуры
	t2 := Create("t2", "Имя: {{.Name}}\n")
	t2.Execute(os.Stdout, struct{ Name string }{"Greg Frost"})
	t2.Execute(os.Stdout, map[string]string{"Name": "Grigoriy Morozov"})

	// Условное выполнение
	t3 := Create("t3", "Условие: {{if . -}} да {{else -}} нет {{end}}\n")
	t3.Execute(os.Stdout, "не пустая строка")
	t3.Execute(os.Stdout, "")

	// Диапазон (range)
	t4 := Create("t4", "Диапазон: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "PHP", "JavaScript"})
}
