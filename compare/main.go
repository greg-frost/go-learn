package main

import (
	"fmt"
	"math"
)

// Сравнение чисел float
func compareFloat(a, b, precise float64) bool {
	return math.Abs(a-b) < precise
}

// Структура "сравнимая"
type Comparable struct {
	Int    int
	String string
	Array  [3]int
}

// Структура "несравнимая"
type NonComparable struct {
	Int    *int
	String *string
}

func main() {
	fmt.Println(" \n[ СРАВНЕНИЕ ]\n ")

	// Сравнение чисел int
	fmt.Println("Int")
	i1 := 10002000
	i2 := 10002000
	fmt.Println("i1 =", i1)
	fmt.Println("i2 =", i2)
	fmt.Println("Равенство:", i1 == i2)
	fmt.Println()

	// Сравнение чисел float
	fmt.Println("Float")
	f1 := 1.0001020202022
	f2 := 1.0001023450119
	fmt.Println("f1 =", f1)
	fmt.Println("f2 =", f2)
	fmt.Println("Равенство:", f1 == f2)
	fmt.Println("Сравнение (0.000001):", compareFloat(f1, f2, 1e-6))
	fmt.Println("Сравнение (0.000000001):", compareFloat(f1, f2, 1e-9))
	fmt.Println()

	// Сравнение сравнимых структур
	fmt.Println("Comparable")
	cs1 := Comparable{
		Int:    1,
		String: "one",
		Array:  [3]int{1, 2, 3},
	}
	cs2 := Comparable{
		Int:    1,
		String: "one",
		Array:  [3]int{1, 2, 3},
	}
	fmt.Println("s1 =", cs1)
	fmt.Println("s2 =", cs2)
	fmt.Println("Равенство:", cs1 == cs2)

	fmt.Println()

	// Сравнение несравнимых структур
	fmt.Println("NonComparable")
	oneInt := 1
	oneString := "one"
	ns1 := NonComparable{
		Int:    &oneInt,
		String: &oneString,
	}
	ns2 := NonComparable{
		Int:    &oneInt,
		String: &oneString,
	}
	fmt.Println("s1 =", ns1)
	fmt.Println("s2 =", ns2)
	fmt.Println("Равенство:", ns1 == ns2)
}
