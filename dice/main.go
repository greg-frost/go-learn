package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" \n[ КУБИК ]\n ")

	var input string
	var n, m int
	var err error

	fmt.Println("Введите два числа в формате {N}d{M},")
	fmt.Println("Enter для повторного броска,")
	fmt.Println("Ctrl+C для выхода:")

	for {
		fmt.Scanln(&input)

		input = strings.TrimSpace(strings.ToLower(input))
		if input != "" {
			d := strings.Split(input, "d")
			if len(d) != 2 {
				return
			}
			n, err = strconv.Atoi(d[0])
			if err != nil {
				return
			}
			m, err = strconv.Atoi(d[1])
			if err != nil {
				return
			}
		}

		res := n + rand.Intn(n*m-n+1)
		fmt.Printf("Бросаем кубик %dd%d... Выпало: %d\n", n, m, res)
	}
}
