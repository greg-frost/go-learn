package main

import (
	"crypto/sha256"
	"fmt"
)

// Количество групп
const groupsCount = 3

// Размер фильтра
const filterSize = 10

// Хэш строки
func hash(str string) []byte {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(str))
	return sha256Hash.Sum(nil)
}

// Группы адресов
func groups(str string, n, m int) []int {
	res := make([]int, n)
	h := hash(str)
	s := len(h) / n
	for i := 0; i < n; i++ {
		var sum int
		for j := i * s; j < len(h); j++ {
			sum += int(h[j])
		}
		res[i] = sum % m
	}
	return res
}

func main() {
	fmt.Println(" \n[ ФИЛЬТР БЛУМА ]\n ")

	insert := []string{"Apple", "Banana", "Peach", "Mango", "Kiwi"}
	check := []string{"Applebee", "Bananan", "Peach", "Mongo", "Kill"}

	for _, s := range insert {
		fmt.Println(s, groups(s, groupsCount, filterSize))
	}
	fmt.Println()
	for _, s := range check {
		fmt.Println(s, groups(s, groupsCount, filterSize))
	}
}
