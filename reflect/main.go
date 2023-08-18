package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Структура "пользовательский тип"
type MyType struct {
	IntField int
	StrField string
	PtrField *float64 // ок, если будет не указатель
	//SliceField []int // недопустимо для сравнения, но ok для рефлексии
}

// Сравнение пользовательских типов
func (mt1 MyType) IsEqual(mt2 MyType) bool {
	return mt1 == mt2
}

// Глубокое сравнение пользовательских типов
func (mt1 MyType) IsDeepEqual(mt2 MyType) bool {
	return reflect.DeepEqual(mt1, mt2)
}

// Структура "нулевой пользовательский тип"
type MyNilType struct{}

// Наивная проверка на nil
func IsNaiveNil(obj interface{}) bool {
	return obj == nil
}

// Реальная проверка на nil
func IsNil(obj interface{}) bool {
	if obj == nil {
		return true
	}

	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr {
		return false
	}

	if objValue.IsNil() {
		return true
	}

	return false
}

// Структура "пользовательская структура"
type MyStruct struct {
	A int
	B string
	C bool
}

// Печать структуры
func PrintStruct(v interface{}) {
	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.Ptr:
		if val.Elem().Kind() != reflect.Struct {
			fmt.Printf("Pointer to %v : %v", val.Elem().Type(), val.Elem())
			return
		}
		val = val.Elem() // если указатель на структуру, берем ее

	case reflect.Struct: // работаем со структурой
	default:
		fmt.Printf("%v : %v", val.Type(), val)
		return
	}

	fmt.Printf("Структура %v (полей - %d):\n", val.Type(), val.NumField())
	for fieldIndex := 0; fieldIndex < val.NumField(); fieldIndex++ {
		field := val.Field(fieldIndex)
		fmt.Printf("\t%v %v: %v\n", val.Type().Field(fieldIndex).Name, field.Type(), field)
	}
}

// Изменение поля структуры
func ChangeFieldByName(v interface{}, fname string, newval int) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}

	field := val.FieldByName(fname)
	if field.IsValid() {
		if field.CanSet() { // ок для экспортируемых полей, переданных по указателю
			switch field.Kind() {
			case reflect.Int:
				field.SetInt(int64(newval))
			case reflect.String:
				field.SetString(strconv.Itoa(newval))
			}
		}
	}
}

func main() {
	fmt.Println(" \n[ РЕФЛЕКСИЯ ]\n ")

	floatValue1, floatValue2 := 10.0, 10.0
	a := MyType{IntField: 1, StrField: "str", PtrField: &floatValue1}
	b := MyType{IntField: 1, StrField: "str", PtrField: &floatValue2}

	fmt.Printf("Обычное равенство a и b: %v\n", a.IsEqual(b))
	fmt.Printf("Глубокое равенство a и b: %v\n", a.IsDeepEqual(b))

	fmt.Println()

	/* Type и Kind */

	fmt.Println("Type и Kind:")
	fmt.Println()

	var Bool *bool
	fmt.Printf(
		"< *bool >\nType: %v\nKind: %v\n\n",
		reflect.ValueOf(Bool).Type(), reflect.ValueOf(Bool).Kind(),
	)

	var Float float32
	fmt.Printf(
		"< float32 >\nType: %v\nKind: %v\n\n",
		reflect.ValueOf(Float).Type(), reflect.ValueOf(Float).Kind(),
	)

	var Map map[string]int
	fmt.Printf(
		"< map[string]int >\nKind: %v\nType: %v\n\n",
		reflect.ValueOf(Map).Type(), reflect.ValueOf(Map).Kind(),
	)

	Struct := struct{ Value int }{}
	fmt.Printf(
		"< struct{Value int} >\nKind: %v\nType: %v\n\n",
		reflect.ValueOf(Struct).Type(), reflect.ValueOf(Struct).Kind(),
	)

	/* Перебор полей структуры */

	fmt.Println("Печать структур:")
	fmt.Println()

	s := &MyStruct{
		A: 3,
		B: "some",
		C: false,
	}

	ChangeFieldByName(s, "A", 5)

	PrintStruct(s)
	PrintStruct(struct {
		E int
		C string
	}{10, "text"})

	fmt.Println()

	/* Разное */

	fmt.Println("Разное:")
	fmt.Println()

	var t *MyNilType
	fmt.Printf("%v - это naive nil? %v\n", reflect.TypeOf(t), IsNaiveNil(t))
	fmt.Printf("%v - это nil? %v\n", reflect.TypeOf(t), IsNil(t))

	fmt.Println()

	varInt := 100
	varIntValue := reflect.ValueOf(varInt)
	fmt.Println("100 - это zero?", varIntValue.IsZero())
	fmt.Println("Значение (100) =", varIntValue.Int())

	fmt.Println()

	var varPtr *int
	varPtrValue := reflect.ValueOf(varPtr)
	fmt.Println("*int - это nil?", varPtrValue.IsNil())
	fmt.Println("*int - это zero?", varPtrValue.IsZero())

	fmt.Println()

	var varBool *bool
	fmt.Println("*bool - это nil?", reflect.ValueOf(varBool).IsNil())

	trueVal := true
	reflect.ValueOf(&varBool).Elem().Set(reflect.ValueOf(&trueVal))
	fmt.Println("(присвоено значение true)")

	fmt.Println("bool (true) - это nil?", reflect.ValueOf(varBool).IsNil())
	fmt.Println("bool (true) =", reflect.ValueOf(varBool).Elem().Bool()) // вывод через рефлексию
	//fmt.Println("bool (true) =", *varBool) // или обычным образом
}
