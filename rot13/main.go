package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	const (
		From, To = 65, 91
		from, to = 97, 123
		shift    = 13
	)

	n, _ := r.r.Read(b)
	readed := 0

	for i := 0; i < n; i++ {
		pos := int(b[i])

		if pos >= From && pos < To {
			pos = From + (((pos - From) + shift) % (To - From))
		} else if pos >= from && pos < to {
			pos = from + (((pos - from) + shift) % (to - from))
		}

		b[i] = byte(pos)
		readed++
	}

	return readed, io.EOF
}

func main() {
	fmt.Println(" \n[ ROT13 ]\n ")

	sipher := "Lbh penpxrq gur pbqr!"

	fmt.Println("Шифр:")
	fmt.Println(sipher)

	fmt.Println()
	fmt.Println("Расшифровка:")

	s := strings.NewReader(sipher)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	fmt.Println()
}
