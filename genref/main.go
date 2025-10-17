package main

import (
	"fmt"
	"reflect"
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
	fmt.Println()

	// Статический тип
	staticSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Статический тип:")
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

	// start := time.Now()
	// fmt.Println("Время выполнения:", time.Since(start))
}
