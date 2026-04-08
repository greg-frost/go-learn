package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Тип "массив"
type Array []int

// Тип "функция сортировки"
type SortFunc func(Array) (Array, int, int)

// Замер времени выполнения
func metricTime(start time.Time) time.Duration {
	return time.Since(start)
}

// Генерация массива случайных чисел
func GenerateArray(size, min, max int) (a Array, duration time.Duration) {
	// Замер времени
	defer func(t time.Time) {
		duration = metricTime(t)
	}(time.Now())

	// Генерация массива
	a = make(Array, size)
	for i := 0; i < size; i++ {
		a[i] = min + rand.Intn(max-min+1)
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
func IsSorted(a Array) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

// Абстрактная сортировка
func Sort(sort SortFunc, arr Array) (a Array, iterations, depth int, duration time.Duration) {
	// Копирование массива
	a = make(Array, len(arr))
	copy(a, arr)

	// Замер времени
	defer func(t time.Time) {
		duration = metricTime(t)
	}(time.Now())

	// Сортировка
	a, iterations, depth = sort(a)
	return a, iterations, depth, duration
}

// Сортировка пузырьком (продолжающаяся, пока есть перестановки)
func BubbleRunSort(a Array) (_ Array, iterations, depth int) {
	var sorted bool
	for !sorted {
		sorted = true
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
			iterations++
		}
	}
	return a, iterations, depth
}

// Сортировка пузырьком (с вытеснением большего значения вверх)
func BubblePopSort(a Array) (_ Array, iterations, depth int) {
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
func SelectSort(a Array) (_ Array, iterations, depth int) {
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
func InsertCopySort(a Array) (_ Array, iterations, depth int) {
	var t int
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[i] <= a[j-1] {
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
func InsertSwapSort(a Array) (_ Array, iterations, depth int) {
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
func CombSort(a Array) (_ Array, iterations, depth int) {
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

// Сортировка кучей (пирамидальная)
func HeapSort(a Array) (_ Array, iterations, depth int) {
	n := len(a)
	for i := (n - 1) / 2; i >= 0; i-- {
		iterations += sink(i, a)
	}

	for n > 0 {
		a[0], a[n-1] = a[n-1], a[0]
		iterations += sink(0, a[:n-1])
		n--
	}
	return a, iterations, depth
}

// Погружение в кучу
func sink(i int, a Array) int {
	var iterations int
	n := len(a)
	j := 2*i + 1

	for j < n {
		if j < n-1 && a[j] < a[j+1] {
			j++
		}
		if a[i] >= a[j] {
			break
		}

		a[i], a[j] = a[j], a[i]
		i = j
		j = 2*i + 1
		iterations++
	}

	return iterations
}

// Сортировка слиянием (с копированием)
func MergeCopySort(a Array) (_ Array, iterations, depth int) {
	if len(a) <= 1 {
		return a, iterations, depth
	}

	mid := len(a) / 2
	b := make(Array, len(a)-mid)
	copy(b, a[mid:])
	a = a[0:mid]
	iterations += len(a) + 2*len(b)
	depth++

	la, li, ld := MergeCopySort(a)
	ra, ri, rd := MergeCopySort(b)
	iterations += li + ri
	depth += (ld + rd) / 2

	return mergeCopy(la, ra), iterations, depth
}

// Слияние массивов (с копированием)
func mergeCopy(l, r Array) Array {
	m := make(Array, 0, len(l)+len(r))

	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			m = append(m, l[0])
			l = l[1:]
		} else {
			m = append(m, r[0])
			r = r[1:]
		}
	}
	m = append(m, l...)
	m = append(m, r...)

	return m
}

// Сортировка слиянием (с перестановками)
func MergeSwapSort(a Array) (_ Array, iterations, depth int) {
	iterations, depth = mergeSwapSort(a, 0, len(a)-1)
	return a, iterations, depth
}

// Рекурсия сортировки слиянием (с перестановками)
func mergeSwapSort(a Array, l, h int) (iterations, depth int) {
	var li, ri, ld, rd int

	if l < h {
		m := l + (h-l)/2
		li, ld = mergeSwapSort(a, l, m)
		ri, rd = mergeSwapSort(a, m+1, h)
		mergeSwap(a, l, m, h)
		iterations += h - l
		depth++
	}
	iterations += li + ri
	depth += (ld + rd) / 2

	return iterations, depth
}

// Слияние массивов (с перестановками)
func mergeSwap(a Array, l, m, h int) {
	c := make(Array, h-l+1)
	for k := l; k <= h; k++ {
		c[k-l] = a[k]
	}

	cm := m - l + 1
	ch := h - l + 1
	i, j := 0, cm
	for k := l; k <= h; k++ {
		if i >= cm {
			a[k] = c[j]
			j++
		} else if j >= ch {
			a[k] = c[i]
			i++
		} else if c[i] <= c[j] {
			a[k] = c[i]
			i++
		} else {
			a[k] = c[j]
			j++
		}
	}
}

// Предел сортировки слиянием
const mergeSortThreshold = 2048

// Сортировка слиянием (параллельная)
func MergeParallelSort(a Array) (_ Array, iterations, depth int) {
	iterations, depth = mergeParallelSort(a, 0, len(a)-1)
	return a, iterations, depth
}

// Рекурсия сортировки слиянием (параллельная)
func mergeParallelSort(a Array, l, h int) (iterations, depth int) {
	var li, ri, ld, rd int

	if l < h {
		m := l + (h-l)/2
		if h-l > mergeSortThreshold {
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				li, ld = mergeParallelSort(a, l, m)
			}()
			go func() {
				defer wg.Done()
				ri, rd = mergeParallelSort(a, m+1, h)
			}()
			wg.Wait()
		} else {
			li, ld = mergeSwapSort(a, l, m)
			ri, rd = mergeSwapSort(a, m+1, h)
		}
		mergeSwap(a, l, m, h)
		iterations += h - l
		depth++
	}
	iterations += li + ri
	depth += (ld + rd) / 2

	return iterations, depth
}

// Быстрая сортировка (с копированием)
func QuickCopySort(a Array) (_ Array, iterations, depth int) {
	if len(a) <= 1 {
		return a, iterations, depth
	}

	p := pivot(0, len(a)-1)
	left := make(Array, 0, len(a)/2)
	mid := make(Array, 0, len(a)/100)
	right := make(Array, 0, len(a)/2)
	for _, v := range a {
		switch {
		case v < a[p]:
			left = append(left, v)
		case v == a[p]:
			mid = append(mid, v)
		case v > a[p]:
			right = append(right, v)
		}
		iterations++
	}
	depth++

	la, li, ld := QuickCopySort(left)
	ra, ri, rd := QuickCopySort(right)
	a = make(Array, 0, len(a))
	a = append(a, la...)
	a = append(a, mid...)
	a = append(a, ra...)
	iterations += li + ri
	depth += (ld + rd) / 2

	return a, iterations, depth
}

// Быстрая сортировка (с перестановками)
func QuickSwapSort(a Array) (_ Array, iterations, depth int) {
	iterations, depth = quickSwapSort(a, 0, len(a)-1)
	return a, iterations, depth
}

// Рекурсия быстрой сортировки (с перестановками)
func quickSwapSort(a Array, l, h int) (iterations, depth int) {
	var li, ri, ld, rd int

	if l < h {
		_, pl, ph := quickSwapPartition(a, l, h)
		li, ld = quickSwapSort(a, l, pl)
		ri, rd = quickSwapSort(a, ph, h)
		iterations += (h - ph) + (pl - l)
		depth++
	}
	iterations += li + ri
	depth += (ld + rd) / 2

	return iterations, depth
}

// Разбиение быстрой сортировки
func quickSwapPartition(a Array, l, h int) (int, int, int) {
	p := pivot(l, h)
	a[p], a[h] = a[h], a[p]

	j := l
	for i := l; i < h; i++ {
		if a[i] <= a[h] {
			a[i], a[j] = a[j], a[i]
			j++
		}
	}
	a[h], a[j] = a[j], a[h]

	jl, jh := j-1, j+1
	for jl >= l && a[jl] == a[j] {
		jl--
	}
	for jh <= h && a[jh] == a[j] {
		jh++
	}

	return j, jl, jh
}

// Предел быстрой сортировки
const quickSortThreshold = 2048

// Быстрая сортировка (параллельная)
func QuickParallelSort(a Array) (_ Array, iterations, depth int) {
	iterations, depth = quickParallelSort(a, 0, len(a)-1)
	return a, iterations, depth
}

// Рекурсия быстрой сортировки (параллельная)
func quickParallelSort(a Array, l, h int) (iterations, depth int) {
	var li, ri, ld, rd int

	if l < h {
		_, pl, ph := quickSwapPartition(a, l, h)
		if h-l > quickSortThreshold {
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				li, ld = quickParallelSort(a, l, pl)
			}()
			go func() {
				defer wg.Done()
				ri, rd = quickParallelSort(a, ph, h)
			}()
			wg.Wait()
		} else {
			li, ld = quickSwapSort(a, l, pl)
			ri, rd = quickSwapSort(a, ph, h)
		}
		iterations += (h - ph) + (pl - l)
		depth++
	}
	iterations += li + ri
	depth += (ld + rd) / 2

	return iterations, depth
}

// Выбор опорного элемента
func pivot(l, h int) int {
	return l // Первый
	// return l + (h-l)/2 // Средний
	// return l + rand.Intn(h-l+1) // Случайный
}

// Сортировка подсчетом
func CountSort(a Array) (_ Array, iterations, depth int) {
	if len(a) <= 1 {
		return a, iterations, depth
	}

	min, max := arrSizes(a)
	iterations += len(a)
	count := make([]int, max-min+1)
	for _, v := range a {
		count[v-min]++
		iterations++
	}

	a = make(Array, 0, len(a))
	for v, c := range count {
		if c > 0 {
			for k := 0; k < c; k++ {
				a = append(a, min+v)
				iterations++
			}
			depth++
		}
	}

	return a, iterations, depth
}

// Число блоков
var blocksCount = runtime.NumCPU()

// Предел блочной сортировки
const blockSortThreshold = 2048

// Блочная сортировка (многопоточная)
func BlockSort(a Array) (_ Array, iterations, depth int) {
	if len(a) <= 1 {
		return a, iterations, depth
	}

	min, max := arrSizes(a)
	iterations += len(a)
	if min == max {
		return a, iterations, depth
	}

	blockSize := (max-min)/blocksCount + 1
	blocks := make([]Array, blocksCount)
	for _, v := range a {
		i := (v - min) / blockSize
		blocks[i] = append(blocks[i], v)
		iterations++
	}

	var wg sync.WaitGroup
	for i, block := range blocks {
		if len(block) > 1 {
			wg.Add(1)
			func(i int, b Array) {
				defer wg.Done()
				var bIterations, bDepth int
				if len(a) > blockSortThreshold {
					blocks[i], bIterations, bDepth = BlockSort(b)
				} else {
					_, bIterations, bDepth = CombSort(b)
				}
				iterations += bIterations
				depth += bDepth
			}(i, block)
		}
	}
	wg.Wait()

	a = make(Array, 0, len(a))
	for _, block := range blocks {
		if len(block) > 0 {
			a = append(a, block...)
			iterations += len(block)
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
func PrintSortReport(name string, iterations, depth int, duration time.Duration, arr Array, size int) {
	fmt.Printf("%s\nИтераций: %d\nГлубина: %d\nВремя: %v\n", name, iterations, depth, duration)
	printArray(arr, size)
}

func main() {
	fmt.Println(" \n[ СОРТИРОВКИ ]\n ")

	size := 10  // Число элементов
	min := 0    // Минимальное значение
	max := 1000 // Максимальное значение

	const printSize = 10      // Размер фрагмента массива для печати
	const slowCap = 10_000    // Лимит на размер для медленных сортировок
	const midCap = 10_000_000 // Лимит на размер для средних сортировок

	// Генерация массива
	fmt.Print("Введите размер массива (и минимум, максимум): ")
	fmt.Scanf("%d %d %d", &size, &min, &max)
	arr, duration := GenerateArray(size, min, max)
	fmt.Println()
	PrintArrayReport(size, min, max, arr, printSize)
	fmt.Println("Время генерации:", duration)

	// Сортировки
	arrSort := make(Array, size)
	var iterations, depth int
	sorts := []struct {
		caption string
		fSort   SortFunc
		isSlow  bool
		isMid   bool
	}{
		{"Сортировка пузырьком, продолжающаяся", BubbleRunSort, true, false},
		{"Сортировка пузырьком, с вытеснением", BubblePopSort, true, false},
		{"Сортировка выбором", SelectSort, true, false},
		{"Сортировка вставками, с копированием", InsertCopySort, true, false},
		{"Сортировка вставками, с перестановками", InsertSwapSort, true, false},
		{"Сортировка расческой", CombSort, false, true},
		{"Сортировка кучей", HeapSort, false, true},
		{"Сортировка слиянием, с копированием", MergeCopySort, false, true},
		{"Сортировка слиянием, с перестановками", MergeSwapSort, false, true},
		{"Сортировка слиянием, параллельная", MergeParallelSort, false, true},
		{"Быстрая сортировка, с копированием", QuickCopySort, false, true},
		{"Быстрая сортировка, с перестановками", QuickSwapSort, false, true},
		{"Быстрая сортировка, параллельная", QuickParallelSort, false, true},
		{"Сортировка подсчетом", CountSort, false, false},
		{"Блочная сортировка", BlockSort, false, true},
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
		if !IsSorted(arrSort) {
			fmt.Println("(массив не отсортирован)")
		}
	}
}
