package main

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

// Печать интерфейса
func DescribeIntf(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Определение типа
func TestType(i interface{}) {
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
func Mul(a interface{}, b int) interface{} {
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
func Sum(v ...interface{}) float64 {
	var res float64
	for _, val := range v {
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
			ref := reflect.ValueOf(val)
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
func SumRef(v ...interface{}) float64 {
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

// Структура "имя"
type Name struct {
	First string
	Last  string
}

// Стрингер
func (n *Name) String() string {
	return n.First + " " + n.Last
}

// Проверка интерфейса Stringer
func IsStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}

// Проверка любого интерфейса
func IsImplements(current, target interface{}) bool {
	iface := reflect.TypeOf(target).Elem()
	v := reflect.ValueOf(iface)
	t := v.Type()
	return t.Implements(iface)
}

func main() {
	fmt.Println(" \n[ ТИПЫ ]\n ")

	// Пустой интерфейс
	fmt.Println("Пустой интерфейс:")
	var i interface{}
	DescribeIntf(i)
	i = 66
	DescribeIntf(i)
	i = "Привет"
	DescribeIntf(i)
	fmt.Println()

	// Приведение типов
	var j interface{} = "Hello"
	fmt.Println("Приведение строки:")
	s := j.(string)
	fmt.Println("Строка:", s)
	s, ok := j.(string)
	fmt.Printf("Строка: %s (ok: %t)\n", s, ok)
	// f = j.(int) // Будет паника
	f, ok := j.(int)
	fmt.Printf("Целое: %d (ok: %t)\n", f, ok)
	fmt.Println()

	// Проверка типов
	fmt.Println("Проверка типов:")
	TestType(21)
	TestType("String")
	TestType(true)
	fmt.Println()

	// Интерфейсы и nil
	fmt.Println("Интерфейсы и nil:")
	var str *string
	fmt.Println("var str *string == nil:", str == nil)
	var intf interface{}
	fmt.Println("var intf interface{} == nil:", intf == nil)
	intf = str
	fmt.Println("intf = str, intf == nil:", intf == nil)
	fmt.Println()

	// Перегруженное умножение
	vs := "Hello"
	vi := 5
	va := 2
	fmt.Println("Перегруженное умножение:")
	fmt.Println("Число:", Mul(vi, va))
	fmt.Println("Строка:", Mul(vs, va))
	fmt.Println()

	// Перегруженная сумма
	type MyInt int64
	var (
		a uint8  = 2
		b int    = 37
		c string = "3.2"
		d MyInt  = 1
	)
	fmt.Println("Перегруженная сумма:")
	fmt.Println("Проверка типа:", Sum(a, b, c, d))
	fmt.Println("Рефлексия:", SumRef(a, b, c, d))
	fmt.Println()

	// Проверка интерфейса
	fmt.Println("Реализация интерфейса:")
	fmt.Println()
	fmt.Println("Проверка типа:")
	name := &Name{First: "Greg", Last: "Frost"}
	if IsStringer(name) {
		fmt.Printf("%T реализует интерфейс *fmt.Stringer\n", name)
	} else {
		fmt.Printf("%T не реализует интерфейс *fmt.Stringer\n", name)
	}
	num := 123
	if IsStringer(num) {
		fmt.Printf("%T реализует интерфейс *fmt.Stringer\n", num)
	} else {
		fmt.Printf("%T не реализует интерфейс *fmt.Stringer\n", num)
	}
	fmt.Println()
	fmt.Println("Рефлексия:")
	stringer := (*fmt.Stringer)(nil)
	if IsImplements(name, stringer) {
		fmt.Printf("%T реализует интерфейс %T\n", name, stringer)
	} else {
		fmt.Printf("%T не реализует интерфейс %T\n", name, stringer)
	}
	writer := (*io.Writer)(nil)
	if IsImplements(num, writer) {
		fmt.Printf("%T реализует интерфейс %T\n", name, writer)
	} else {
		fmt.Printf("%T не реализует интерфейс %T\n", name, writer)
	}
}
