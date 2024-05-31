package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
)

func main() {
	fmt.Println(" \n[ РЕГУЛЯРНЫЕ ВЫРАЖЕНИЯ ]\n ")

	// Шаблон и строка
	pattern := "п([а-яё]+)к"
	text := "персик пёсик пик"
	fmt.Println("Шаблон:", pattern)
	fmt.Println("Строка:", text)
	fmt.Println()

	// Соответствие
	match, err := regexp.MatchString(pattern, text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(`Соответствие:`, match)

	// Компиляция
	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	// Обязательная компиляция
	r = regexp.MustCompile(pattern)
	fmt.Println("Соответствие в байтах:", r.Match([]byte(text)))
	fmt.Println()

	// Поиск
	fmt.Println("Поиск строки:", r.FindString(text))
	fmt.Println("Индексы строки:", r.FindStringIndex(text))
	fmt.Println("Поиск подстрок:", r.FindStringSubmatch(text))
	fmt.Println("Индексы подстрок:", r.FindStringSubmatchIndex(text))
	fmt.Println("Поиск всех строк:", r.FindAllString(text, -1))
	fmt.Println("Поиск двух строк:", r.FindAllString(text, 2))
	fmt.Println("Индексы всех подстрок:")
	fmt.Println(r.FindAllStringSubmatchIndex(text, -1))

	fmt.Println()

	// Замена
	fmt.Println("Замена:", r.ReplaceAllString("этот персик созрел", "<фрукт>"))
	fmt.Println("В верхний регистр:", string(
		r.ReplaceAllFunc([]byte("этот персик"), bytes.ToUpper),
	))
}
