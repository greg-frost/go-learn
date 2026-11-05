package main

import (
	"crypto/sha256"
	"fmt"
)

// Количество групп
const groupsCount = 3

// Размер фильтра
const filterSize = 10

// Структура "фильтр Блума"
type BloomFilter struct {
	bits []bool
}

// Конструктор
func NewBloomFilter() *BloomFilter {
	return &BloomFilter{
		bits: make([]bool, filterSize),
	}
}

// Вставка значения
func (bf *BloomFilter) Insert(key string) {
	bits := groups(key, groupsCount, filterSize)
	for _, i := range bits {
		bf.bits[i] = true
	}
}

// Проверка наличия значения
func (bf *BloomFilter) Contains(key string) bool {
	bits := groups(key, groupsCount, filterSize)
	for _, i := range bits {
		if bf.bits[i] == false {
			return false
		}
	}
	return true
}

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

	bloom := NewBloomFilter()

	// Вставка
	insert := []string{"Apple", "Banana", "Peach", "Mango", "Kiwi"}
	fmt.Println("Вставка значений")
	for _, s := range insert {
		bloom.Insert(s)
		fmt.Println(s)
	}
	fmt.Println()

	// Проверка
	fmt.Println("Проверка значений")
	check := []string{"Applebee", "Bananan", "Peach", "Mongo", "Kill"}
	for _, s := range check {
		res := bloom.Contains(s)
		fmt.Printf("%s: %t\n", s, res)
	}
}
