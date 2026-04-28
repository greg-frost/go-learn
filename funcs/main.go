package main

import (
	"fmt"
	"math"
	"time"
)

// Сумма
func Sum(vals ...int) (res int) {
	for _, v := range vals {
		res += v
	}
	return
}

// Максимум
func Max(vals ...int) (max int) {
	max = math.MinInt
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

// Половина
func Half(val int) (res int, isEven bool) {
	res = val / 2
	if res%2 == 0 {
		isEven = true
	}
	return
}

// Факториал
func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

// Ряд Фибоначчи
func FibRow(n int) []int {
	res := make([]int, n+1)

	res[0], res[1] = 0, 1
	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res
}

// Значение Фибоначчи
func FibValue(n int) int {
	if n <= 1 {
		return n
	}
	return FibValue(n-1) + FibValue(n-2)
}

// Кэш для значений Фибоначчи
var memo = make(map[int]int)

// Значение Фибоначчи (с мемоизацией)
func FibValueMemo(n int) int {
	if n <= 1 {
		return n
	}
	if v, ok := memo[n]; ok {
		return v
	}

	res := FibValueMemo(n-1) + FibValueMemo(n-2)
	memo[n] = res
	return res
}

// Замыкание Фибоначчи
func FibClosure() func() int {
	prev, next := 0, 1

	return func() int {
		curr := prev
		prev, next = next, prev+next
		return curr
	}
}

// Генератор нечетных чисел
func MakeOddGenerator() func() uint {
	var i uint = 1
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Генератор четных чисел
func MakeEvenGenerator() func() uint {
	var i uint = 0
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Тип "функция"
type Function func(a, b int) int

// Выполнение
func TypeFunc(f Function) int {
	return f(3, 25)
}

// Структура "вычитатель"
type Subber struct {
	val int
}

// Вычитание
func (s Subber) SubFrom(val int) int {
	return val - s.val
}

// Абстрактная функция
func Compute(fn func(float64, float64) float64) float64 {
	return fn(2, 3)
}

// Замыкание
func Adder() func(int) int {
	var sum int
	return func(x int) int {
		sum += x
		return sum
	}
}

// Тип "фигура"
type Figures int

const (
	square Figures = iota
	circle
	triangle
	hexagon
)

// Выбор функции для вычисления площади фигуры
func area(f Figures) (func(float64) float64, bool) {
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

// Замыкание для подсчета количества вызовов
func CountCall(f func(string)) func(string) {
	var count int
	return func(s string) {
		count++
		fmt.Println("Кол-во вызовов:", count)
		f(s)
	}
}

// Замыкание для подсчета времени выполнения
func MetricCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now()
		f(s)
		fmt.Println("Время выполнения:", time.Since(start))
	}
}

// Простая функция печати
func MyPrinter(s string) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println(s)
}

// Отложенная функция
func Deferrer(d int, dp *int) {
	fmt.Println("Отложенная функция:", d, *dp)
}

// Очевидная функция
func Intuitive() string {
	value := "Казалось бы,"
	defer func() { value = "На самом деле..." }()
	return value
}

// Неочевидная функция
func Unintuitive() (value string) {
	defer func() { value = "На самом деле..." }()
	return "Казалось бы,"
}

func main() {
	fmt.Println(" \n[ ФУНКЦИИ ]\n ")

	var a, b, c int
	fmt.Print("Введите до трех чисел (через пробел): ")
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println()

	// Сумма через функцию
	sum := Sum(a, b, c)
	fmt.Println("Общая сумма:", sum)

	// Сумма через срез
	var res int
	slice := []int{a, b}
	sumSlice := func(slice []int) int {
		for _, v := range slice {
			res += v
		}
		return res
	}
	fmt.Println("Сумма среза:", sumSlice(slice))
	fmt.Println("Сумма двух срезов:", sumSlice(slice))
	fmt.Println()

	// Максимум, половина и факториал
	fmt.Println("Наибольшее число:", Max(a, b, c))
	res, isEven := Half(sum)
	fmt.Printf("Половина (четная): %d (%t)\n", res, isEven)
	fmt.Println("Факториал:", Fact(res))
	fmt.Println()

	// Фибоначчи
	fmt.Println("Ряд Фибоначчи:", FibRow(res))
	fmt.Println("Число Фибоначчи:", FibValue(res))
	fmt.Println("Число Фибоначчи (мемо.):", FibValueMemo(res))
	fmt.Println("Замыкание Фибоначчи:")
	f := FibClosure()
	for i := 0; i <= res; i++ {
		fmt.Print(f(), " ")
	}
	fmt.Println()
	fmt.Println()

	// Нечетные и четные числа
	nextOdd := MakeOddGenerator()
	fmt.Println("Нечетные числа:", nextOdd(), nextOdd(), nextOdd())
	nextEven := MakeEvenGenerator()
	fmt.Println("Четные числа:", nextEven(), nextEven(), nextEven())
	fmt.Println()

	// Тип-функция и формы методы
	fmt.Println("Тип-функция:", TypeFunc(func(a int, b int) int {
		return a * b
	}))
	mySubber := Subber{val: 50}
	fmt.Println("Обычный метод:", mySubber.SubFrom(120))
	fs1 := mySubber.SubFrom
	fmt.Println("Значение метода:", fs1(115))
	fs2 := Subber.SubFrom
	fmt.Println("Выражение метода:", fs2(mySubber, 110))
	fmt.Println()

	// Абстрактная функция
	fmt.Println("Абстракные функции:")
	hypot := func(x, y float64) float64 {
		return math.Pow(2*x, 2*y)
	}
	fmt.Println("hypot(2, 4) =", hypot(2, 4))
	fmt.Println("compute(hypot) =", Compute(hypot))
	fmt.Println("compute(math.Pow) =", Compute(math.Pow))
	fmt.Println()

	// Замыкания
	fmt.Println("Замыкания:")
	pos, neg := Adder(), Adder()
	for i := 0; i < Max(a, b, c); i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fmt.Println()

	// Полиморфная функция
	myFigure := triangle
	myX := 3.5
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("ОШИБКА: Нет функции для расчета фигуры")
		return
	}
	myArea := ar(myX)
	fmt.Printf("Фигура: %d, площадь: %2.2f\n\n", myFigure, myArea)

	// Отслеживание количества вызовов
	fmt.Println("Количество вызовов функции:")
	countedPrint := CountCall(MyPrinter)
	countedPrint("Hello")
	countedPrint("World")
	fmt.Println()

	// Отслеживание времени выполнения
	fmt.Println("Время выполнения функции:")
	metricPrint := MetricCall(countedPrint)
	metricPrint("Привет")
	metricPrint("Мир")
	fmt.Println()

	// Отложенные функции
	df := 10
	defer Deferrer(df, &df)
	df = 100
	fmt.Println(Intuitive())
	fmt.Println(Unintuitive())
}
