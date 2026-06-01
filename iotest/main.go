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
	bytes  []byte
}

// Чтение (символов нижнего регистра)
func (lcr *LowerCaseReader) Read(bytes []byte) (n int, err error) {
	if lcr.bytes == nil {
		lcr.bytes = make([]byte, 0)
	}
	if len(bytes) == 0 {
		return 0, nil
	}
	n, err = lcr.reader.Read(lcr.bytes)
	var count int
	for i := 0; i < n; i++ {
		if lcr.bytes[i] >= 'a' && lcr.bytes[i] <= 'z' {
			bytes[count] = lcr.bytes[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	fmt.Println(" \n[ ТЕСТИРОВАНИЕ IO ]\n ")

	str := "aBcDeFgHiJ"

	// Создание ридера
	lcr := &LowerCaseReader{
		strings.NewReader(str),
		make([]byte, 10),
	}

	// Чтение ридера
	body, err := io.ReadAll(lcr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Все символы:", str)
	fmt.Println("Только нижний регистр:", string(body))
}
