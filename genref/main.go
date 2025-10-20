package main

import (
	"fmt"
	"reflect"
	"time"
)

// Обнуление среза (статический тип)
func nullifySliceStatic(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = 0
	}
}

// Обнуление среза (дженерики)
func nullifySliceGenerics[T any](slice []T) {
	var nullValue T
	for i := 0; i < len(slice); i++ {
		slice[i] = nullValue
	}
}

// Обнуление среза (рефлексия)
func nullifySliceReflect(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return
	}
	for i := 0; i < val.Len(); i++ {
		val.Index(i).SetZero()
	}
}

func main() {
	fmt.Println(" \n[ ДЖЕНЕРИКИ И РЕФЛЕКСИЯ ]\n ")

	// Обнуление
	fmt.Println("Обнуление среза")
	fmt.Println("---------------")

	// Статический тип
	staticSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Статический:")
	fmt.Println(staticSlice)
	nullifySliceStatic(staticSlice)
	fmt.Println(staticSlice)
	fmt.Println()

	// Дженерики
	genericsSlice := []float64{1.23, 4.567, 8.9}
	// genericsSlice := []bool{true, true, false, true, false}
	fmt.Println("Дженерики:")
	fmt.Println(genericsSlice)
	nullifySliceGenerics(genericsSlice)
	fmt.Println(genericsSlice)
	fmt.Println()

	// Рефлексия
	reflectSlice := []string{"one", "two", "three"}
	// reflectSlice := []rune{'a', 'b', 'c'}
	fmt.Println("Рефлексия:")
	fmt.Println(reflectSlice)
	nullifySliceReflect(reflectSlice)
	fmt.Println(reflectSlice)
	fmt.Println()

	// Сравнение
	fmt.Println("Сравнение скорости")
	fmt.Println("------------------")

	// Параметры
	const times = 1e5
	var slice = make([]int, 1000)

	// Статический тип
	start := time.Now()
	for i := 0; i < times; i++ {
		nullifySliceStatic(slice)
	}
	fmt.Println("Статический:", time.Since(start))

	// Дженерики
	start = time.Now()
	for i := 0; i < times; i++ {
		nullifySliceGenerics(slice)
	}
	fmt.Println("Дженерики:", time.Since(start))

	// Рефлексия
	start = time.Now()
	for i := 0; i < times; i++ {
		nullifySliceReflect(slice)
	}
	fmt.Println("Рефлексия:", time.Since(start))
}
