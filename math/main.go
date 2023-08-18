package main

import (
	"fmt"
	"math"
)

// Сумма
func sum(a int, b int) int {
	return a + b
}

// Произведение
func prod(a int, b int) int {
	return a * b
}

// Произведение через сумму
func prodBySum(val int, times int) (res int) {
	if times >= 0 {
		for i := 0; i < times; i++ {
			res += val
		}
	} else {
		for i := 0; i > times; i-- {
			res -= val
		}
	}
	return
}

// Пропорция
func split(sum int, prop float64) (x, y int) {
	x = int(float64(sum) * prop)
	y = sum - x
	return
}

// Тип "ошибка отрицательного квадрата"
type ErrNegativeSqrt float64

// Вывод ошибки отрицательного квадрата
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Нельзя вычислить Sqrt для: %v", float64(e))
}

// Квадратный корень
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var z float64
	var prev float64

	//z = 1
	//z = x
	z = x / 2

	for i := 0; i < 10; i++ {
		prev = z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 0.00001 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(" \n[ МАТЕМАТИКА ]\n ")

	/* Ввод данных */

	var a, b int

	fmt.Print("Введите два числа (через пробел): ")
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println()

	/* Сумма и произведение */

	fmt.Println("Сумма", a, "и", b, "равна", sum(a, b))
	fmt.Println("Произведение", a, "и", b, "равно", prod(a, b))
	fmt.Println("Произведение через сумму", a, "и", b, "равно", prodBySum(a, b))

	fmt.Println()

	var (
		x float64 = float64(prod(a, b))
		y int     = sum(a, b)
		p float64 = 0.75
	)

	/* Пропорция */

	c, d := split(y, p)
	fmt.Println("Разделение числа", y, "в пропорции", p, "равно", c, "и", d)
	fmt.Printf("Само число %d принадлежит к типу %T\n", y, y)

	fmt.Println()

	/* Квадратный корень */

	sr, _ := sqrt(x)
	fmt.Println("mySQRT   (", x, ") =", sr)
	fmt.Println("mathSQRT (", x, ") =", math.Sqrt(x))
	fmt.Println(sqrt(-x))

	fmt.Println()

	/* Побитовый сдвиг */

	fmt.Println("Степени числа 2:")

	pow := make([]int, 9)
	for i := range pow {
		pow[i] = 1 << uint(i)
		fmt.Printf("%d ", pow[i])
	}

	fmt.Println()
}
