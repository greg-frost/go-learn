package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("неверная UTF-8 строка")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	fmt.Println(" \n[ FUZZING ]\n ")

	text := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(text)
	revrev, _ := Reverse(rev)
	fmt.Printf("Оригинал: %q\n", text)
	fmt.Printf("Перевернуто 1: %q\n", rev)
	fmt.Printf("Перевернуто 2: %q\n", revrev)
}
