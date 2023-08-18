package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println(" \n[ МЬЮТЕКСЫ ]\n ")

	m := new(sync.Mutex)

	/* Последовательные рутины */

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			//defer m.Unlock() // можно и так

			fmt.Print(i, " start ")
			time.Sleep(time.Second)
			fmt.Print(i, " end ")

			m.Unlock()
		}(i)
	}

	/* Ожидание ввода */

	var input string
	fmt.Scanln(&input)
}
