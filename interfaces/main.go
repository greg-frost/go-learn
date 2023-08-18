package main

import (
	"fmt"
)

// Интерфейс "логика"
type Logic interface {
	Process(data string) string
}

// Структура "провайдер логики"
type LogicProvider struct {
	firstName string
	lastName  string
	age       int
	isMerried bool
}

// Обработка логики
func (lp LogicProvider) Process(data string) string {
	isMerried := "холост"
	if lp.isMerried == true {
		isMerried = "женат"
	}

	return fmt.Sprintf("%v %v, %v года, %v. %v", lp.firstName, lp.lastName, lp.age, isMerried, data)
}

// Структура "клиент"
type Client struct {
	L Logic
}

// Обработка клиента
func (c Client) Program() {
	fmt.Println(c.L.Process("Yoroshiku!"))
}

func main() {
	fmt.Println(" \n[ ИНТЕРФЕЙС ]\n ")

	c := Client{
		L: LogicProvider{
			"Eikichi",
			"Onidzuka",
			22,
			false,
		},
	}

	fmt.Println("Приветствие:")
	c.Program()
}
