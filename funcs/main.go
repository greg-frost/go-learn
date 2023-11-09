package main

import (
	"fmt"
	"math"
)

// Сумма
func sum(vals ...int) (res int) {
	for _, v := range vals {
		res += v
	}
	return
}

// Максимум
func max(vals ...int) (max int) {
	max = vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

// Половина
func half(val int) (res int, isEven bool) {
	res = val / 2
	if res%2 == 0 {
		isEven = true
	}
	return
}

// Факториал
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

// Ряд Фибоначчи
func fibRow(n int) []int {
	res := make([]int, n+1)

	res[0] = 0
	res[1] = 1

	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res
}

// Значение Фибоначчи
func fibValue(n int) int {
	if n <= 1 {
		return n
	}

	return fibValue(n-1) + fibValue(n-2)
}

var memo = map[int]int{}

// Значение Фибоначчи (с мемоизацией)
func fibValueMemo(n int) int {
	if n <= 1 {
		return n
	}

	if v, ok := memo[n]; ok {
		return v
	}

	memo[n] = fibValueMemo(n-1) + fibValueMemo(n-2)

	return memo[n]
}

// Замыкание Фибоначчи
func fibClosure() func() int {
	prev := 0
	next := 1

	return func() int {
		curr := prev
		prev, next = next, prev+next
		return curr
	}
}

// Генератор нечетных чисел
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Генератор четных чисел
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Тип "функция"
type prodder func(a, b int) int

// Выполнение типа-функции
func prodFunc(p prodder) int {
	return p(3, 25)
}

// Структура "вычитатель"
type subber struct {
	val int
}

// Вычитание из
func (s subber) subFrom(val int) int {
	return val - s.val
}

// Абстрактная функция
func compute(fn func(float64, float64) float64) float64 {
	return fn(2, 3)
}

// Замыкание
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Тип "фигура"
type figures int

const (
	square figures = iota
	circle
	triangle
	hexagon
)

// Выбор функции для вычисления площади фигуры
func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(x float64) float64 { return x * x }, true
	case circle:
		return func(x float64) float64 { return 3.142 * x * x }, true
	case triangle:
		return func(x float64) float64 { return 0.433 * x * x }, true
	default:
		return nil, false
	}
}

// Отложенная функция
func deferrer(d int, dp *int) {
	fmt.Println("Отложенная функция:", d, *dp)
}

// Очевидная функция
func intuitive() string {
	value := "Казалось бы,"
	defer func() { value = "На самом деле..." }()
	return value
}

// Неочевидная функция
func unintuitive() (value string) {
	defer func() { value = "На самом деле..." }()
	return "Казалось бы,"
}

func main() {
	fmt.Println(" \n[ ФУНКЦИИ ]\n ")

	var a, b, c int

	fmt.Print("Введите до трех чисел (через пробел): ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println()

	/* Сумма через функцию */

	sum := sum(a, b, c)
	fmt.Println("Общая сумма равна", sum)

	/* Сумма серез срез */

	slice := []int{a, b}
	res := 0
	sumSlice := func(slice []int) int {
		for _, v := range slice {
			res += v
		}
		return res
	}
	fmt.Println("Сумма среза равна", sumSlice(slice))
	fmt.Println("Сумма двух срезов равна", sumSlice(slice))

	fmt.Println()

	/* Максимум, половина и факториал */

	fmt.Println("Наибольшее число", max(a, b, c))

	res, isEven := half(sum)
	fmt.Println("Половина (четная)", res, "(", isEven, ")")

	fmt.Println("Факториал равен", fact(res))

	fmt.Println()

	/* Фибоначчи */

	fmt.Println("Ряд Фибоначчи", fibRow(res))
	fmt.Println("Число Фибоначчи", fibValue(res))
	fmt.Println("Число Фибоначчи (мемо.)", fibValueMemo(res))
	fmt.Println("Замыкание Фибоначчи:")
	f := fibClosure()

	for i := 0; i <= res; i++ {
		fmt.Print(f(), " ")
	}

	fmt.Println(" \n ")

	/* Нечетные и четные числа */

	nextOdd := makeOddGenerator()
	fmt.Println("Нечетные числа:", nextOdd(), nextOdd(), nextOdd())

	nextEven := makeEvenGenerator()
	fmt.Println("Четные числа:", nextEven(), nextEven(), nextEven())

	fmt.Println()

	/* Тип-функция и формы методы */

	fmt.Println("Тип-функция:", prodFunc(func(a int, b int) int {
		return a * b
	}))

	mySubber := subber{val: 50}
	fmt.Println("Обычный метод:", mySubber.subFrom(120))

	fs1 := mySubber.subFrom
	fmt.Println("Значение метода:", fs1(115))

	fs2 := subber.subFrom
	fmt.Println("Выражение метода:", fs2(mySubber, 110))

	fmt.Println()

	/* Абстрактная функция */

	fmt.Println("Абстракные функции:")
	hypot := func(x, y float64) float64 {
		return math.Pow(2*x, 2*y)
	}
	fmt.Println("hypot(2, 4) =", hypot(2, 4))

	fmt.Println("compute(hypot) =", compute(hypot))
	fmt.Println("compute(math.Pow) =", compute(math.Pow))

	fmt.Println()

	/* Замыкания */

	fmt.Println("Замыкания:")
	pos, neg := adder(), adder()
	for i := 0; i < max(a, b, c); i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println()

	/* Отложенная функция */

	myFigure := triangle
	myX := 3.5

	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("ОШИБКА: Нет функции для расчета фигуры")
		return
	}

	myArea := ar(myX)
	fmt.Printf("Фигура: %d, площадь: %2.2f\n\n", myFigure, myArea)

	/* Отложенные функции */

	df := 10
	defer deferrer(df, &df)
	df = 100

	fmt.Println(intuitive())
	fmt.Println(unintuitive())
}
