package main

import (
	"fmt"
	"math"
)

// Интерфейс "фигура"
type Shape interface {
	area() float64
	perimeter() float64
}

// Структура "мультифигура"
type MultiShape struct {
	shapes []Shape
}

// Площадь мультифигуры
func (m MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

// Общая площадь фигур
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// Структура "круг"
type Circle struct {
	x, y, r float64
}

// Площадь круга
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

// Периметр круга
func (c Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

// Структура "прямоугольник"
type Rectangle struct {
	x1, y1, x2, y2 float64
}

// Площадь прямоугольника
func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// Периметр прямоугольника
func (r Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (l + w) * 2
}

// Расстояние между двумя точками
func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Структура "мужчина"
type Men struct {
	age   int
	isMen bool
}

// Структура "человек"
type Person struct {
	name string
	*Men
}

// Убить человека
func (p Person) kill() {
	var who, what string
	if p.isMen {
		who = "Мужик"
		what = "был убит"
	} else {
		who = "Баба"
		what = "была убита"
	}
	fmt.Printf("%s %s в возрасте %d лет\n", who, what, p.age)
}

// Изменение возраста
func ChangeAge(p Person, age int) {
	p.age = age
}

func main() {
	fmt.Println(" \n[ ООП ]\n ")

	// Круг
	c := new(Circle)
	c.x = 0
	c.y = 5
	c.r = 7
	fmt.Println("Круг:", *c)
	fmt.Println("Площадь:", math.Round(c.area()))
	fmt.Println("Периметр:", math.Round(c.perimeter()))
	fmt.Println()

	// Прямоугольник
	r := Rectangle{x1: 1, y1: 3, x2: 10, y2: 12}
	fmt.Println("Прямоугольник:", r)
	fmt.Println("Площадь:", r.area())
	fmt.Println("Периметр:", r.perimeter())
	fmt.Println()

	// Мультифигура
	s := new(MultiShape)
	fmt.Println("Фигуры:", *s)
	fmt.Println("Общая площадь (метод):", math.Round((*s).area()))
	fmt.Println("Общая площадь (функция):", math.Round(totalArea(c, &r)))
	fmt.Println()

	// Человек
	var p = Person{"Greg", &Men{34, true}}
	fmt.Printf("%#+v\n", p)
	fmt.Printf("Имя: %s, возраст: %d, пол (мужик): %t\n", p.name, p.age, p.Men.isMen)
	ChangeAge(p, 35)
	(&p).kill()
}
