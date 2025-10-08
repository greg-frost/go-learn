package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Структура "генератор"
type generator struct {
	rnd rand.Source
}

// Конструктор генератора
func NewRand(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

// Генерация случайных чисел
func (g *generator) Read(bytes []byte) (n int, err error) {
	// for i := range bytes {
	// 	randInt := g.rnd.Int63() // Случайное число от 0 до 2^63
	// 	randByte := byte(randInt)
	// 	bytes[i] = randByte
	// }
	for i := 0; i+8 <= len(bytes); i += 8 {
		binary.LittleEndian.PutUint64(
			bytes[i:i+8],
			uint64(g.rnd.Int63()),
		)
	}
	return len(bytes), nil
}

// Интерфейс "хэшер"
type Hasher interface {
	io.Writer
	Hash() byte
}

// Структура "хэш"
type hash struct {
	result byte
}

// Конструктор хэша
func NewHash(start byte) Hasher {
	return &hash{
		result: start,
	}
}

// Запись хэша
func (h *hash) Write(bytes []byte) (n int, err error) {
	for _, b := range bytes {
		h.result = (h.result^b)<<1 + b%2
	}
	return len(bytes), nil
}

// Получение хэша
func (h hash) Hash() byte {
	return h.result
}

// Структура "ограниченный ридер"
type LimitedReader struct {
	reader io.Reader
	left   int
}

// Конструктор ридера
func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{reader: r, left: n}
}

// Чтение ридера
func (r *LimitedReader) Read(p []byte) (int, error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	if r.left < len(p) {
		p = p[0:r.left]
	}
	n, err := r.reader.Read(p)
	r.left -= n
	return n, err
}

func main() {
	fmt.Println(" \n[ IO-ПАКЕТ ]\n ")

	// Рандом
	fmt.Println("Рандом:")
	generator := NewRand(time.Now().UnixNano())
	buf := make([]byte, 16)
	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf)
		fmt.Printf("%v (%d)\n", buf, n)
	}
	fmt.Println()

	// Хэш
	hasher := NewHash(0)
	hasher.Write(buf)
	fmt.Printf("Хэш: %v \n", hasher.Hash())
	fmt.Println()

	// Ограниченное чтение
	r := strings.NewReader("Some io.Reader stream to read\n")
	lr := LimitReader(r, 4)
	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("...")
}
