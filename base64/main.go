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
	stdEnc := base64.StdEncoding.EncodeToString(data)
	stdDec, _ := base64.StdEncoding.DecodeString(stdEnc)
	fmt.Println("Encode:", stdEnc)
	fmt.Println("Decode:", string(stdDec))
	fmt.Println()

	// URL-кодирование
	urlEnc := base64.URLEncoding.EncodeToString(data)
	urlDec, _ := base64.URLEncoding.DecodeString(urlEnc)
	fmt.Println("URL Encode:", urlEnc)
	fmt.Println("URL Decode:", string(urlDec))
	fmt.Println()

	// Сравнение
	if string(data) == string(stdDec) && string(data) == string(urlDec) {
		fmt.Println("Раскодировано верно")
	} else {
		fmt.Println("Раскодировано неверно")
	}
}
