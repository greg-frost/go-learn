package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Тип "массив"
type Array []int

// Тип "функция сортировки"
type SortFunc func(Array) (Array, int, int)

// Замер времени выполнения
func metricTime(start time.Time) time.Duration {
	return time.Now().Sub(start)
}

// Генерация массива случайных чисел
func GenerateArray(size, min, max int) (a Array, duration time.Duration) {
	// Замер времени
	defer func(t time.Time) {
		duration = metricTime(t)
	}(time.Now())

	// Генерация массива
	a = make(Array, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		a[i] = rand.Int()%(max+1-min) + min
	}

	return a, duration
}

// Границы элементов массива
func arrSizes(a Array) (min, max int) {
	min, max = a[0], a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}

// Проверка отсортированности массива
func isSorted(a Array) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

// Абстрактная сортировка
func Sort(fSort SortFunc, arr Array) (a Array, iterations, depth int, duration time.Duration) {
	// Копирование массива
	a = make(Array, len(arr))
	copy(a, arr)

	// Замер времени
	defer func(t time.Time) {
		duration = metricTime(t)
	}(time.Now())

	// Сортировка
	a, iterations, depth = fSort(a)

	return a, iterations, depth, duration
}

// Сортировка пузырьком №1 (продолжающаяся, пока есть перестановки)
func bubbleRunSort(a Array) (_ Array, iterations, depth int) {
	isRunning := true

	for isRunning {
		isRunning = false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				isRunning = true
			}
			iterations++
		}
	}

	return a, iterations, depth
}

// Сортировка пузырьком №2 (с вытеснением большего значения вверх)
func bubblePopSort(a Array) (_ Array, iterations, depth int) {
	for j := len(a) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
			iterations++
		}
	}

	return a, iterations, depth
}

// Сортировка выбором
func selectSort(a Array) (_ Array, iterations, depth int) {
	for i := 0; i < len(a)-1; i++ {
		k := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[k] {
				k = j
			}
			iterations++
		}
		a[i], a[k] = a[k], a[i]
	}

	return a, iterations, depth
}

// Сортировка вставками (с копированием)
func insertCopySort(a Array) (_ Array, iterations, depth int) {
	var t int
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[i] < a[j-1] {
			j--
			iterations++
		}

		t = a[i]
		copy(a[j+1:], a[j:i])
		a[j] = t
		iterations += i - j + 1
	}
	return a, iterations, depth
}

// Сортировка вставками (с перестановками)
func insertSwapSort(a Array) (_ Array, iterations, depth int) {
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[j] < a[j-1] {
			a[j], a[j-1] = a[j-1], a[j]
			j--
			iterations++
		}
		iterations++
	}

	return a, iterations, depth
}

// Сортировка расческой
func combSort(a Array) (_ Array, iterations, depth int) {
	const factor = 1.247
	stepFactor := float64(len(a)) / factor

	for stepFactor > 1 {
		step := int(math.Round(stepFactor))
		for i, j := 0, step; j < len(a); i, j = i+1, j+1 {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
			iterations++
		}
		stepFactor /= factor
	}

	return a, iterations, depth
}

// Быстрая сортировка
func quickSort(a Array) (_ Array, iterations, depth int) {
	if len(a) <= 1 {
		return a, iterations, depth
	}

	pivot := a[0]

	left := make(Array, 0, len(a)/2)
	middle := make(Array, 0, len(a)/100)
	right := make(Array, 0, len(a)/2)

	for _, v := range a {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			middle = append(middle, v)
		case v > pivot:
			right = append(right, v)
		}
		iterations++
	}
	depth++

	leftA, leftI, leftD := quickSort(left)
	rightA, rightI, rightD := quickSort(right)

	a = make(Array, 0, len(a))

	a = append(a, leftA...)
	a = append(a, middle...)
	a = append(a, rightA...)

	iterations += leftI + rightI
	depth += leftD + rightD

	return a, iterations, depth
}

// Слияние двух массивов
func merge(left, right Array) Array {
	merged := make(Array, 0, len(left)+len(right))

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	merged = append(merged, left...)
	merged = append(merged, right...)

	return merged
}

// Сортировка слиянием
func mergeSort(a Array) (_ Array, iterations, depth int) {
	if len(a) <= 1 {
		return a, iterations, depth
	}

	middle := len(a) / 2

	b := make(Array, len(a)-middle)
	copy(b, a[middle:])
	a = a[0:middle]

	iterations++
	depth++

	leftA, leftI, leftD := mergeSort(a)
	rightA, rightI, rightD := mergeSort(b)

	iterations += leftI + rightI
	depth += leftD + rightD

	return merge(leftA, rightA), iterations, depth
}

// Сортировка подсчетом
func countSort(a Array) (_ Array, iterations, depth int) {
	if len(a) <= 1 {
		return a, iterations, depth
	}

	_, max := arrSizes(a)
	iterations += len(a)

	count := make([]int, max+1)
	for _, v := range a {
		count[v]++
		iterations++
	}

	a = make(Array, 0)
	for v, c := range count {
		if c > 0 {
			for i := 0; i < c; i++ {
				a = append(a, v)
				iterations++
			}
			depth++
		}
	}

	return a, iterations, depth
}

// Блочная сортировка (многопоточная)
func blockSort(a Array) (_ Array, iterations, depth int) {
	if len(a) <= 1 {
		return a, iterations, depth
	}

	const blocksCount = 10
	var wg sync.WaitGroup

	min, max := arrSizes(a)
	iterations += len(a)

	blockSize := (max - min) / blocksCount
	if blockSize == 0 {
		blockSize = 1
	}
	blocks := make([]Array, max/blockSize+1)

	var id int
	for _, v := range a {
		id = v / blockSize
		blocks[id] = append(blocks[id], v)
		iterations++
	}

	for _, block := range blocks {
		if len(block) > 1 {
			wg.Add(1)
			go func(b Array) {
				defer wg.Done()
				_, bIterations, bDepth := combSort(b)
				iterations += bIterations
				depth += bDepth + 1
			}(block)
		}
	}

	wg.Wait()

	a = make(Array, 0)
	for _, block := range blocks {
		if len(block) > 0 {
			a = append(a, block...)
		}
	}

	return a, iterations, depth
}

// Печать массива
func printArray(arr Array, printSize int) {
	if len(arr) > printSize {
		fmt.Printf("%v %v\n", arr[:printSize/2], arr[len(arr)-printSize/2:])
	} else {
		fmt.Println(arr)
	}
}

// Печать отчета по массиву
func PrintArrayReport(size, min, max int, arr Array, printSize int) {
	fmt.Printf("Массив из %d случайных элементов от %d до %d:\n", size, min, max)
	printArray(arr, printSize)
}

// Печать отчета по сортировке
func PrintSortReport(name string, iterations, depth int, duration time.Duration, arr Array, printSize int) {
	fmt.Printf("%s\nИтераций: %d\nГлубина: %d\nВремя: %v\n", name, iterations, depth, duration)
	printArray(arr, printSize)
}

func main() {
	fmt.Println(" \n[ СОРТИРОВКИ ]\n ")

	size := 10
	min := 0
	max := 1000

	const printSize = 10      // размер фрагмента массива для печати
	const slowCap = 10_000    // лимит на размер для медленных сортировок
	const midCap = 10_000_000 // лимит на размер для средних сортировок

	var duration time.Duration

	/* Генерация массива */

	fmt.Print("Введите размер массива (и минимум, максимум): ")
	fmt.Scanf("%d %d %d", &size, &min, &max)

	arr, duration := GenerateArray(size, min, max)

	fmt.Println()
	PrintArrayReport(size, min, max, arr, printSize)
	fmt.Println("Время генерации:", duration)

	/* Сортировки */

	arrSort := make(Array, size)
	var iterations, depth int

	sorts := []struct {
		caption string
		fSort   SortFunc
		isSlow  bool
		isMid   bool
	}{
		{"Сортировка пузырьком, продолжающаяся", bubbleRunSort, true, false},
		{"Сортировка пузырьком, с вытеснением", bubblePopSort, true, false},
		{"Сортировка выбором", selectSort, true, false},
		{"Сортировка вставками, с копированием", insertCopySort, true, false},
		{"Сортировка вставками, с перестановками", insertSwapSort, true, false},
		{"Сортировка расческой", combSort, false, true},
		{"Быстрая сортировка", quickSort, false, true},
		{"Сортировка слиянием", mergeSort, false, true},
		{"Сортировка подсчетом", countSort, false, false},
		{"Блочная сортировка", blockSort, false, true},
	}

	for _, sort := range sorts {
		if sort.isSlow && size > slowCap {
			continue
		}
		if sort.isMid && size > midCap {
			continue
		}

		fmt.Println()
		arrSort, iterations, depth, duration = Sort(sort.fSort, arr)
		PrintSortReport(sort.caption, iterations, depth, duration, arrSort, printSize)
		if !isSorted(arrSort) {
			fmt.Println("(массив не отсортирован)")
		}
	}
}
