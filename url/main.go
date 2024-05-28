package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
)

func main() {
	fmt.Println(" \n[ URL ]\n ")

	// Пример
	s := "postgres://user:pass@host.com:5432/path?key=value#anchor"

	// Парсинг URL
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	/* Компоненты URL */

	fmt.Println("Схема (протокол):", u.Scheme)

	fmt.Println("Пользователь:", u.User)
	fmt.Println("Логин:", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("Пароль:", p)

	fmt.Println("Адрес:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Хост:", host)
	fmt.Println("Порт:", port)

	fmt.Println("Путь:", u.Path)

	fmt.Println("Строка запроса:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("Карта параметров:", m)
	fmt.Println("Параметр key:", m["key"][0])

	fmt.Println("Якорь:", u.Fragment)
}
