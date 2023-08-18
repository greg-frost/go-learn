package main

import (
	"fmt"
	"strings"
)

// Печать интерфейса
func describeIntf(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Определение типа
func testType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Дважды %v равно %v\n", v, v*2)
	case string:
		fmt.Printf("%q - это строка (%v байт)\n", v, len(v))
	case fmt.Stringer:
		fmt.Printf("%q - это реализация Stringer\n", v.String())
	default:
		fmt.Printf("Непонятно, что за тип %T!\n", v)
	}
}

// Перегруженное умножение
func mul(a interface{}, b int) interface{} {
	switch v := a.(type) {
	case int:
		return v * b
	case string:
		return strings.Repeat(v, b)
	case fmt.Stringer:
		return strings.Repeat(v.String(), b)
	default:
		return nil
	}
}

func main() {
	fmt.Println(" \n[ ТИПЫ ]\n ")

	/* Пустой интерфейс */

	fmt.Println("Пустой интерфейс:")

	var i interface{}
	describeIntf(i)

	i = 66
	describeIntf(i)

	i = "Привет"
	describeIntf(i)

	fmt.Println()

	/* Приведение типов */

	var j interface{} = "Hello"

	fmt.Println("Приведение строки:")

	s := j.(string)
	fmt.Println("Строка:", s)

	s, ok := j.(string)
	fmt.Println("Строка:", s, "(ok)", ok)

	//f = j.(int) // будет паника
	//fmt.Println("Целое:", f)

	f, ok := j.(int)
	fmt.Println("Целое:", f, "(ok)", ok)

	fmt.Println()

	/* Проверка типа */

	fmt.Println("Проверка типов:")

	testType(21)
	testType("String")
	testType(true)

	fmt.Println()

	/* Проверка типа */

	fmt.Println("Интерфейсы и nil:")

	var str *string
	fmt.Println("var str *string == nil:", str == nil)

	var intf interface{}
	fmt.Println("var intf interface{} == nil:", intf == nil)

	intf = str
	fmt.Println("intf = str, intf == nil:", intf == nil)

	fmt.Println()

	/* Перегруженное умножение */

	vs := "Hello"
	vi := 5
	va := 2

	fmt.Println("Перегруженное умножение:")
	fmt.Println("Число:", mul(vi, va))
	fmt.Println("Строка:", mul(vs, va))
}
