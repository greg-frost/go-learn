package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" \n[ БРОСОК КУБИКОВ ]\n ")

	fmt.Println("Введите два числа в формате {N}d{M},")
	fmt.Println("Enter для повторного броска,")
	fmt.Println("Ctrl+C для выхода:")

	var (
		input string
		n, m  int
		err   error
	)

	for {
		fmt.Scanln(&input)

		input = strings.ToLower(input)
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
		fmt.Printf("Бросаем кубики %dd%d... Выпало: %d\n", n, m, res)
	}
}
