package main

import (
	"fmt"

	"go-learn/import/math"
	"go-learn/import/slice"
)

func main() {
	fmt.Println(" \n[ ИМПОРТ ]\n ")

	/* Математические операции */

	var a, b int = 5, 7

	fmt.Println("Сумма", a, "и", b, "равна", math.Sum(a, b))
	fmt.Println("Произведение", a, "и", b, "равно", math.Prod(a, b))
	fmt.Println("Произведение через сумму", a, "и", b, "равно", math.ProdBySum(a, b))

	fmt.Println()

	/* Операции со срезами */

	s := slice.Slice{1, 2, 3}
	fmt.Println("Срез:", s)

	fmt.Println("Сумма среза:", slice.SumSlice(s))

	slice.MapSlice(s, func(i slice.Element) slice.Element {
		return i * 3
	})
	fmt.Println("Срез, умноженный на три:", s)

	fmt.Println("Свертка слайса сложением:",
		slice.FoldSlice(s,
			func(x slice.Element, y slice.Element) slice.Element {
				return x + y
			},
			0))

	fmt.Println("Свертка слайса умножением:",
		slice.FoldSlice(s,
			func(x slice.Element, y slice.Element) slice.Element {
				return x * y
			},
			1))
}
