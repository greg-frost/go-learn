package main

import (
	"fmt"
	"io"
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
	b, err := readAll(r, 3)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

// Чтение всего ридера (с повторами при ошибке)
func readAll(r io.Reader, retries int) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				return b, nil
			}
			retries--
			if retries < 0 {
				return b, err
			}
		}
	}
}

func main() {
	fmt.Println(" \n[ ТЕСТИРОВАНИЕ IO ]\n ")

	str := "NO pain IS a NO gain"

	// Создание ридера
	lcr := &LowerCaseReader{
		reader: strings.NewReader(str),
	}

	// Использование ридера
	fmt.Println("Все символы:", str)
	fmt.Print("Только нижний регистр: ")
	Print(lcr)
}
