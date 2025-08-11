package main

import (
	"fmt"
)

// Интерфейс "состояние"
type State interface {
	Order()
	Pay()
	Deliver()
	Recieve()
	Cancel()
	Return()
}

func main() {
	fmt.Println(" \n[ СОСТОЯНИЕ ]\n ")
}
