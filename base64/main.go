package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println(" \n[ BASE64 ]\n ")

	// Сообщение
	data := []byte("message1234567890!?$*&()'-=@~")
	fmt.Println("Сообщение:", string(data))
	fmt.Println()

	// Стандартное кодирование
	sEnc := base64.StdEncoding.EncodeToString(data)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Encode:", sEnc)
	fmt.Println("Decode:", string(sDec))
	fmt.Println()

	// URL кодирование
	uEnc := base64.URLEncoding.EncodeToString(data)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println("URL Encode:", uEnc)
	fmt.Println("URL Decode:", string(uDec))
	fmt.Println()

	if string(data) == string(sDec) && string(data) == string(uDec) {
		fmt.Println("Раскодировано верно")
	} else {
		fmt.Println("Раскодировано неверно")
	}
}
