package filter

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Целевые переменные
var target []string
var targetInt []int

// Начинается с гласной
func startsWithVowel(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

// Одна случайная строка
func buildString() string {
	size := 5
	out := make([]byte, size)
	for i := 0; i < size; i++ {
		out[i] = byte(rand.Intn(26) + 65)
	}
	return string(out)
}

// Генерация набора случайных строк
func setup() []string {
	rand.Seed(1)
	size := 1000
	vals := make([]string, size)
	for i := 0; i < size; i++ {
		vals[i] = buildString()
	}
	return vals
}

// Генерация набора целых
func setupInt() []int {
	size := 1000
	vals := make([]int, size)
	for i := 0; i < size; i++ {
		vals[i] = i
	}
	return vals
}

// Четное ли число
func isEven(i int) bool {
	return i%2 == 0
}

// Тестирование фильтрации (рефлексия)
func TestFilterReflect(t *testing.T) {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	longNames := Filter(names, func(s string) bool {
		return len(s) > 3
	}).([]string)
	fmt.Println(longNames)
	if diff := cmp.Diff(longNames, []string{"Andrew", "Clara", "Hortense"}); diff != "" {
		t.Error(diff)
	}

	ages := []int{20, 50, 13}
	adults := Filter(ages, func(age int) bool {
		return age >= 18
	}).([]int)
	fmt.Println(adults)
	if diff := cmp.Diff(adults, []int{20, 50}); diff != "" {
		t.Error(diff)
	}
}

// Бенчмарк фильтрации строк (рефлексия)
func BenchmarkFilterReflectString(b *testing.B) {
	vals := setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = Filter(vals, startsWithVowel).([]string)
	}
}

// Бенчмарк фильтрации строк (без рефлексии)
func BenchmarkFilterString(b *testing.B) {
	vals := setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = FilterString(vals, startsWithVowel)
	}
}

// Бенчмарк фильтрации чисел (рефлексия)
func BenchmarkFilterReflectInt(b *testing.B) {
	vals := setupInt()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetInt = Filter(vals, isEven).([]int)
	}
}

// Бенчмарк фильтрации чисел (без рефлексии)
func BenchmarkFilterInt(b *testing.B) {
	vals := setupInt()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		targetInt = FilterInt(vals, isEven)
	}
}
