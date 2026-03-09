package main

import (
	"fmt"
)

// Среднее и минимальное значение
func AvgAndMin(arr []float64) (avg, min float64) {
	if len(arr) == 0 {
		return
	}

	min = arr[0]
	for _, v := range arr {
		avg += v
		if v < min {
			min = v
		}
	}
	avg = avg / float64(len(arr))

	return
}

// Поиск индексов, сумма значений которых равна k
func FindKeysBySumVals(arr []int, k int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == k {
				return []int{i, j}
			}
		}
	}

	return nil
}

// Поиск индексов, сумма значений которых равна k (быстрый алгоритм)
func FindKeysBySumValsFast(arr []int, k int) []int {
	seen := make(map[int]int)
	for i, v := range arr {
		if j, ok := seen[k-v]; ok {
			return []int{i, j}
		}
		seen[v] = i
	}

	return nil
}

// Печать среза
func PrintSlice(s []int) {
	fmt.Printf("len=%d  cap=%d  s=%v\n", len(s), cap(s), s)
}

// Удаление элемента по индексу
func RemoveAtIndex(s []int, i int) []int {
	if i >= len(s) {
		return s
	}

	n := len(s) - 1
	s[i], s[n] = s[n], s[i]
	return s[:n]
}

func main() {
	fmt.Println(" \n[ МАССИВЫ ]\n ")

	// Входные данные
	var (
		arr = []float64{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
		}
		matrix = [...][]int{
			{1, 1, 1},
			{0, 1, 1},
			{0, 0, 1},
		}
		s = []int{
			2, 3, 5,
			7, 11, 13,
		}
		slice, merge []float64
		avg, min     float64
	)

	// Матрица
	fmt.Println("Матрица:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%-3d", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println()

	// Массив
	avg, min = AvgAndMin(arr)
	fmt.Println("Массив:", arr)
	fmt.Printf("Среднее: %.2f\n", avg)
	fmt.Println("Минимальное:", min)
	fmt.Println()

	// Срез
	slice = arr[3:12]
	avg, min = AvgAndMin(slice)
	fmt.Println("Срез:", slice)
	fmt.Printf("Среднее: %.2f\n", avg)
	fmt.Println("Минимальное:", min)
	fmt.Println()

	// Размеры среза
	fmt.Println("Размеры срезов:")
	PrintSlice(s)
	s = s[:0]
	PrintSlice(s)
	s = s[:4]
	PrintSlice(s)
	s = s[2:]
	PrintSlice(s)
	s = RemoveAtIndex(s, 1)
	PrintSlice(s)
	fmt.Println()

	// Операции
	merge = make([]float64, 5)
	copy(merge, arr)
	merge = append(merge, 48, 49, 50)
	merge = append(merge, slice...)
	fmt.Println("Странный массив:")
	fmt.Println(merge)
	fmt.Println()

	// Индексы среза
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 10
	fmt.Printf("Индексы среза %v,\nсумма значений которых = %d: %v\n",
		values, k, FindKeysBySumVals(values, k))
}
