package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println(" \n[ ЧИСЛА ]\n ")

	// Парсинг
	fmt.Println("Парсинг")
	fmt.Println("-------")
	fmt.Println()
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println("Int:", i)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println("Uint:", u)
	b, _ := strconv.ParseBool("true")
	fmt.Println("Bool:", b)
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println("Float:", f)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println("Hex:", d)
	k, _ := strconv.Atoi("135")
	fmt.Println("Atoi:", k)
	_, err := strconv.Atoi("zero")
	if err != nil {
		fmt.Println("Ошибка парсинга числа")
	}
	fmt.Println()

	// Рандом
	fmt.Println("Рандом")
	fmt.Println("------")
	fmt.Println()
	fmt.Printf("Ints: %d, %d, %d\n",
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
	)
	fmt.Println("Float:", rand.Float64())
	fmt.Printf("Floats [5, 10]: %f, %f\n",
		(rand.Float64()*5)+5,
		(rand.Float64()*5)+5,
	)
	var seed int64 = 1500
	s1 := rand.NewSource(seed)
	r1 := rand.New(s1)
	fmt.Printf("Ints (одно зерно): %d, %d, %d\n",
		r1.Intn(100),
		r1.Intn(100),
		r1.Intn(100),
	)
	s2 := rand.NewSource(seed)
	r2 := rand.New(s2)
	fmt.Printf("Ints (одно зерно): %d, %d, %d\n",
		r2.Intn(100),
		r2.Intn(100),
		r2.Intn(100),
	)
}
