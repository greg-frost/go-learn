package main

import (
	"fmt"
)

// Структура "сотрудник"
type Employee struct {
	Name string
	ID   string
}

// Описание сотрудника
func (e Employee) Description() string {
	return fmt.Sprintf("%s (ID: %s)", e.Name, e.ID)
}

// Структура "менеджер"
type Manager struct {
	Employee
	Reports []Employee
}

// Поиск новых менеджеров
func (m Manager) FindNewEmployees() []Employee {
	return []Employee{
		{
			Name: "Ada Wong",
			ID:   "098263",
		},
		{
			Name: "John Smith",
			ID:   "655312",
		},
		{
			Name: "Charles the Prince",
			ID:   "122213",
		},
	}
}

// Структура "внутренняя"
type Inner struct {
	A int
}

// Печать внутреннего числа
func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("внутреннее: %d", val)
}

// Удвоение внутреннего числа
func (i Inner) Double() string {
	result := i.A * 2
	return i.IntPrinter(result)
}

// Структура "внешняя"
type Outer struct {
	Inner
	S string
}

// Печать внешнего числа
func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("внешнее: %d", val)
}

func main() {
	fmt.Println(" \n[ КОМПОЗИЦИЯ ]\n ")

	/* Композиция */

	m := Manager{
		Employee: Employee{
			Name: "Greg Frost",
			ID:   "000001",
		},
		Reports: []Employee{},
	}

	fmt.Println("Менеджер:")
	fmt.Println()

	fmt.Println("Имя:", m.Name)
	fmt.Println("Идентификатор:", m.ID)
	fmt.Println(m.Description())

	m.Reports = m.FindNewEmployees()
	m.Reports = append(m.Reports, m.FindNewEmployees()...)
	m.Reports = append(m.Reports, m.FindNewEmployees()...)
	fmt.Println("Подчиненные:")
	fmt.Println(m.Reports)

	fmt.Println()

	/* Виртуальные методы */

	fmt.Println("Виртуальные методы:")

	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Привет",
	}
	fmt.Print(o.S, ", ", o.Double())
	fmt.Println()
}
