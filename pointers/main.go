package main

import (
	"fmt"
)

// Обнуление
func Zero(xPtr *int) {
	*xPtr = 0
}

// Копирование
func Copy(ptr1, ptr2 *int) {
	*ptr1 = *ptr2
}

// Привязка
func Bind(ptr1, ptr2 *int) {
	ptr1 = ptr2
}

// Установка значения
func Set(ptr *int, val int) {
	*ptr = val
}

// Квадратный корень
func Square(x *float64) {
	*x = *x * *x
}

// Замена переменных местами
func Swap(ptr1, ptr2 *int) {
	ptr3 := new(int)
	*ptr3 = *ptr1
	*ptr1 = *ptr2
	*ptr2 = *ptr3
}

// Замена переменных местами (более элегантная)
func Swappy(ptr1, ptr2 *int) {
	*ptr1, *ptr2 = *ptr2, *ptr1
}

// Структура "ничто"
type Nil struct {
	value string
}

// Допустимый метод nil-структуры
func (n *Nil) OK() string {
	return "все ОК"
}

// Недопустимый метод nil-структуры
func (n *Nil) Panic() string {
	return n.value
}

func main() {
	fmt.Println(" \n[ УКАЗАТЕЛИ ]\n ")

	x := 1
	y := 2
	ptr := new(int)

	// Обнуление, копирование, привязка
	fmt.Println("Было:", x, y)
	Zero(&x)
	Copy(&x, &y)
	Bind(&x, &y) // Не сработает

	// Установка значения
	y = 5
	Set(ptr, 7)
	fmt.Println("Стало:", x, y, *ptr)
	fmt.Println()

	// Квадратный корень
	s := 1.5
	Square(&s)
	fmt.Println("Квадрат:", s)

	// Замена переменных
	Swappy(&x, &y)
	Swap(&x, &y)
	Swappy(&x, &y)
	fmt.Println("Замена:", x, y)
	fmt.Println()

	// Указатель на указатель
	a := 100
	var b *int = &a
	var c **int = &b
	fmt.Println("Адрес a:", &a)
	fmt.Println("Значение a:", a)
	fmt.Println("Адрес b:", &b)
	fmt.Println("Значение b:", b)
	fmt.Println("Разыменование b:", *b)
	fmt.Println("Адрес c:", &c)
	fmt.Println("Значение c:", c)
	fmt.Println("Разыменование c:", *c)
	fmt.Println("Двойное разыменование c:", **c)
	fmt.Println()

	// Nil-методы
	var n *Nil
	fmt.Println("Допустимый nil-метод:", n.OK())
	fmt.Println("Недопустимый nil-метод:",
		"(будет паника)",
		// n.Panic(),
	)
}
