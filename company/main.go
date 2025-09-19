package main

import (
	"fmt"
)

// Интерфейс "работник"
type Worker interface {
	Work(tasks []string) string
}

// Структура "компания"
type Company struct {
	personell []Worker
}

// Наем сотрудника
func (c *Company) Hire(newbie Worker) {
	c.personell = append(c.personell, newbie)
}

// Работа сотрудника
func (c Company) Process(id int, tasks []string) string {
	return c.personell[id].Work(tasks)
}

// Структура "человек"
type Person struct {
	name     string
	homework string
	children []*Person
}

// Список детей человека
func (p Person) Children() []*Person {
	return p.children
}

// Работа человека
func (p Person) Work(tasks []string) string {
	res := p.name + " работает:"
	for _, task := range tasks {
		res += "\n Я выполняю " + task
	}
	res += fmt.Sprintf("\n(домашняя работа: %s)", p.homework)
	return res
}

// Стрингер человека
func (p Person) String() string {
	return p.name
}

// Структура "робот"
type Robot struct {
	model       string
	serialId    int
	workCounter int
}

// Работа робота
func (r *Robot) Work(tasks []string) string {
	res := fmt.Sprintf("%s работает:", r)
	for _, task := range tasks {
		res += "\n Я выполняю " + task + ", хозяин"
	}
	r.workCounter += len(tasks)
	res += fmt.Sprintf("\n(выполнено работ: %d)", r.workCounter)
	return res
}

// Стрингер робота
func (r Robot) String() string {
	return fmt.Sprintf("Робот %s серийный_номер %d", r.model, r.serialId)
}

func main() {
	fmt.Println(" \n[ КОМПАНИЯ ]\n ")

	// Сотрудники
	person1 := Person{name: "Grog", homework: "пить"}
	person2 := Person{name: "Grig", homework: "жать"}
	person3 := Person{name: "Greg", homework: "жить",
		children: []*Person{&person1, &person2}}
	robot := &Robot{model: "GF Mk.1", serialId: 12226122}

	// Наем
	var company Company
	company.Hire(person1)
	company.Hire(person2)
	company.Hire(person3)
	company.Hire(robot)

	// Вывод
	fmt.Println("Сотрудники:")
	fmt.Println(person3, person3.Children())
	fmt.Println()
	fmt.Println(company.Process(2, []string{"еда", "отдых", "программирование"}))
	fmt.Println()
	fmt.Println(company.Process(3, []string{"починка", "поломка", "головоломка"}))
}
