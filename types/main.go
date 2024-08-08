package main

import (
	"fmt"
	"reflect"
	"strconv"
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

// Перегруженная сумма
func sum(v ...interface{}) float64 {
	var res float64
	for _, val := range v {
		ref := reflect.ValueOf(val)
		switch val.(type) {
		case int:
			res += float64(val.(int))
		case int32:
			res += float64(val.(int32))
		case int64:
			res += float64(val.(int64))
		case uint:
			res += float64(val.(uint))
		case uint8:
			res += float64(val.(uint8))
		case uint32:
			res += float64(val.(uint32))
		case uint64:
			res += float64(val.(uint64))
		case string:
			a, err := strconv.ParseFloat(ref.String(), 64)
			if err != nil {
				fmt.Printf("Ошибка парсинга строки %s, игнорирую.\n", val)
				continue
			}
			res += a
		default:
			fmt.Printf("Неизвестный тип %T, игнорирую.\n", val)
		}
	}
	return res
}

// Перегруженная сумма (рефлексия)
func sumRef(v ...interface{}) float64 {
	var res float64
	for _, val := range v {
		ref := reflect.ValueOf(val)
		switch ref.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			res += float64(ref.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
			res += float64(ref.Uint())
		case reflect.String:
			a, err := strconv.ParseFloat(ref.String(), 64)
			if err != nil {
				fmt.Printf("Ошибка парсинга строки %s, игнорирую.\n", val)
				continue
			}
			res += a
		default:
			fmt.Printf("Неизвестный тип %T, игнорирую.\n", val)
		}
	}
	return res
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
	fmt.Println()

	/* Перегруженная сумма */

	type MyInt int64

	var (
		a uint8  = 2
		b int    = 37
		c string = "3.2"
		d MyInt  = 1
	)

	fmt.Println("Перегруженная сумма:")
	fmt.Println("Проверка типа:", sum(a, b, c, d))
	fmt.Println("Рефлексия:", sumRef(a, b, c, d))
}
