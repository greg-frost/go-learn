package main

import (
	"fmt"
)

// Удаление дубликатов
func removeDuplicates(s []string) []string {
	return uniqueStrings(s)
}

// Определение уникальных чисел
func uniqueInts(values []int) []int {
	if len(values) == 0 {
		return values
	}
	res := make([]int, 0, len(values))
	seen := make(map[int]struct{})
	for _, v := range values {
		if _, ok := seen[v]; !ok {
			res = append(res, v)
			seen[v] = struct{}{}
		}
	}
	return res
}

// Определение уникальных строк
func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return values
	}
	res := make([]string, 0, len(values))
	seen := make(map[string]bool)
	for _, v := range values {
		if !seen[v] {
			res = append(res, v)
			seen[v] = true
		}
	}
	return res
}

func main() {
	fmt.Println(" \n[ КАРТЫ ]\n ")

	// Поиск и перебор
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
	}
	fmt.Println("Элементы:")
	for k, v := range elements {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println()
	if el, ok := elements["Li"]; ok {
		fmt.Printf("Элемент: %s (%s)\n", el["name"], el["state"])
	}
	fmt.Println()

	// Уникальные значения
	animals := []string{"кошка", "собака", "птица", "собака", "попугай", "кошка"}
	fmt.Println("Животные:", animals)
	fmt.Println("Без повторений:", removeDuplicates(animals))
	fmt.Println()

	// Непредсказуемость карт
	// (порядок вывода и изменение при итерации)
	fmt.Println("Непредсказуемость карт:")
	numbers := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	for k, v := range numbers {
		if v {
			numbers[10+k] = true
		}
	}
	for k, v := range numbers {
		fmt.Printf("%v: %v\n", k, v)
	}
}
