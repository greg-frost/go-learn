package main

import (
	"fmt"
)

// Структура "коллекция"
type Collection struct {
	elements   []int
	comparator Comparator
}

// Конструктор коллекции
func NewCollection(comparator Comparator) *Collection {
	return &Collection{
		elements:   make([]int, 0),
		comparator: comparator,
	}
}

// Добавление элементов в коллекцию
func (c *Collection) Add(values ...int) {
	for _, v := range values {
		c.elements = append(c.elements, v)
	}
}

// Сортировка коллекции
func (c *Collection) Sort() {
	for j := 1; j < len(c.elements); j++ {
		for i := j - 1; i >= 0; i-- {
			// Шаблонный метод
			if c.comparator.Compare(c.elements[i], c.elements[j]) > 0 {
				c.Swap(i, j)
				j = i
			}
		}
	}
}

// Смена элементов коллекции местами
func (c *Collection) Swap(i, j int) {
	c.elements[i], c.elements[j] = c.elements[j], c.elements[i]
}

// Элементы коллекции
func (c *Collection) Values() []int {
	return c.elements
}

// Интерфейс "сравниватель"
type Comparator interface {
	Compare(a, b int) int
}

// Сравниватель по возрастанию
type AscComparator struct{}

// Конструктор сравнивателя по возрастанию
func NewAscComparator() Comparator {
	return new(AscComparator)
}

// Сравнение по возрастанию
func (*AscComparator) Compare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(" \n[ ШАБЛОННЫЙ МЕТОД ]\n ")

	// Сортировка по возрастанию
	comparator := NewAscComparator()
	collection := NewCollection(comparator)
	collection.Add(5, 7, 1, 2, 6, 3, 4, 9, 8, 10)

	fmt.Println("Сортировка по возрастанию")
	fmt.Println("До сортировки:   ", collection.Values())
	collection.Sort()
	fmt.Println("После сортировки:", collection.Values())
}
