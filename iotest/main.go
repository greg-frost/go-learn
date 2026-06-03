package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// Структура "ридер нижнего регистра"
type LowerCaseReader struct {
	reader io.Reader
}

// Чтение (символов нижнего регистра)
func (lcr *LowerCaseReader) Read(buf []byte) (n int, err error) {
	n, err = lcr.reader.Read(buf)
	var count int
	for i := 0; i < n; i++ {
		if buf[i] >= 'a' && buf[i] <= 'z' {
			buf[count] = buf[i]
			count++
		}
	}
	return count, err
}

// Печать содержимого ридера
func Print(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func main() {
	fmt.Println(" \n[ ТЕСТИРОВАНИЕ IO ]\n ")

	str := "NO pain IS a NO gain"

	// Создание ридера
	lcr := &LowerCaseReader{
		reader: strings.NewReader(str),
	}

	// Чтение содержимого
	lower, err := io.ReadAll(lcr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Все символы:", str)
	fmt.Println("Только нижний регистр:", string(lower))
}
