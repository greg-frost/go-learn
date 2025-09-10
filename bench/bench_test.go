package bench

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"
	"text/template"

	"go-learn/base"
)

// Путь к файлу
var path = base.Dir("bench")
var filename = filepath.Join(path, "bench.go")

// "Черная дыра"
var blackhole int

// Тестирование длины файла
func TestFileLen(t *testing.T) {
	g, err := FileLen(filename, 1)
	if err != nil {
		t.Fatal(err)
	}
	e := 388
	if g != e {
		t.Errorf("Длина файла: получено %d, ожидается %d", g, e)
	}
}

// Бенчмарк длины файла
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

// Бенчмарк шаблонов
func BenchmarkTemplates(b *testing.B) {
	b.Logf("b.N = %d\n", b.N)
	tpl := "Hello {{.Name}}"
	data := map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(tpl)
		t.Execute(&buf, data)
		buf.Reset()
	}
}

// Бенчмарк шаблонов (скомпилированный)
func BenchmarkCompiledTemplates(b *testing.B) {
	b.Logf("b.N = %d\n", b.N)
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t.Execute(&buf, data)
		buf.Reset()
	}
}

// Бенчмарк шаблонов (параллельный)
func BenchmarkParallelTemplates(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := map[string]string{
		"Name": "World",
	}
	// Гонка данных!
	// var buf bytes.Buffer
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
