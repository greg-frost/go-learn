package main

import (
	"fmt"
	"log"
	"time"

	"golearn/codec/user"

	"github.com/ugorji/go/codec"
)

func main() {
	fmt.Println(" \n[ CODEC ]\n ")

	// JSON-обработчик
	jh := new(codec.JsonHandle)

	// Данные
	u := &user.User{
		Name:  "Greg",
		Email: "greg-frost@yandex.ru",
	}
	var out []byte
	var u2 user.User

	times := 1000000
	fmt.Println("Число повторов:", times)
	fmt.Println()

	// Кодирование
	start := time.Now()
	for i := 0; i < times; i++ {
		err := codec.NewEncoderBytes(&out, jh).Encode(&u)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Кодирование:")
	fmt.Println(string(out))
	fmt.Println("Заняло", time.Now().Sub(start))
	fmt.Println()

	// Декодирование
	start = time.Now()
	for i := 0; i < times; i++ {
		err := codec.NewDecoderBytes(out, jh).Decode(&u2)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Декодирование:")
	fmt.Printf("%+v\n", u2)
	fmt.Println("Заняло", time.Now().Sub(start))
}
