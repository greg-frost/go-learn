package main

import (
	"fmt"
)

// Обнуление
func zero(xPtr *int) {
	*xPtr = 0
}

// Копирование
func copy(ptr1, ptr2 *int) {
	*ptr1 = *ptr2
}

// Привязка
func bind(ptr1, ptr2 *int) {
	ptr1 = ptr2
}

// Установка значения
func set(ptr *int, val int) {
	*ptr = val
}

// Квадратный корень
func square(x *float64) {
	*x = *x * *x
}

// Замена переменных местами
func swap(ptr1, ptr2 *int) {
	ptr3 := new(int)
	*ptr3 = *ptr1
	*ptr1 = *ptr2
	*ptr2 = *ptr3
}

// Замена переменных местами (более элегантная)
func swappy(ptr1, ptr2 *int) {
	*ptr1, *ptr2 = *ptr2, *ptr1
}

func main() {
	fmt.Println(" \n[ УКАЗАТЕЛИ ]\n ")

	x := 1
	y := 2
	ptr := new(int)

	// Обнуление, копирование, привязка
	fmt.Println("Было:", x, y)
	zero(&x)
	copy(&x, &y)
	bind(&x, &y) // Не сработает

	// Установка значения
	y = 5
	set(ptr, 7)
	fmt.Println("Стало:", x, y, *ptr)
	fmt.Println()

	// Квадратный корень
	s := 1.5
	square(&s)
	fmt.Println("Квадрат:", s)

	// Замена переменных
	swappy(&x, &y)
	swap(&x, &y)
	swappy(&x, &y)
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
}
