package bench

import (
	"fmt"
	"os"
	"testing"
)

// Путь к файлу
var filename = os.Getenv("GOPATH") + "/src/golearn/hello/main.go"

// "Черная дыра"
var blackhole int

// Тест
func TestFileLen(t *testing.T) {
	g, err := FileLen(filename, 1)
	if err != nil {
		t.Fatal(err)
	}
	e := 98
	if g != e {
		t.Errorf("Длина файла: получено %d, ожидается %d", g, e)
	}
}

// Бенчмарк
func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen(filename, v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
