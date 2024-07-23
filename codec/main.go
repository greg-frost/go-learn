package main

import (
	"fmt"
	"log"

	"golearn/codec/user"

	"github.com/ugorji/go/codec"
)

func main() {
	fmt.Println(" \n[ CODEC ]\n ")

	// JSON-обработчик
	jh := new(codec.JsonHandle)

	// Кодирование
	u := &user.User{
		Name:  "Greg",
		Email: "greg-frost@yandex.ru",
	}
	var out []byte
	err := codec.NewEncoderBytes(&out, jh).Encode(&u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Кодирование:")
	fmt.Println(string(out))
}
