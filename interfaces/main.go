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

// Конструктор
func NewClient(firstName, lastName string, age int, isMarried bool) Client {
	return Client{
		L: LogicProvider{
			firstName: firstName,
			lastName:  lastName,
			age:       age,
			isMerried: isMarried,
		},
	}
}

func main() {
	fmt.Println(" \n[ ИНТЕРФЕЙС ]\n ")

	fmt.Println("Приветствие:")
	c := NewClient("Eikichi", "Onidzuka", 22, false)
	c.Program()
}
