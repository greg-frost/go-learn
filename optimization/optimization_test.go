package optimization

import (
	"testing"
)

// "Черная дыра"
var blackhole int64

// Базовый размер среза
const size = 1e6

// Создание среза
func makeSlice(n int) []int64 {
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = int64(i + 1)
	}
	return res
}

// Бенчмарк суммы всех значений
func BenchmarkSum(b *testing.B) {
	s := makeSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := Sum(s)
		blackhole = total
	}
}

// Бенчмарк суммы каждого второго значения
func BenchmarkSum2(b *testing.B) {
	s := makeSlice(size * 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := Sum2(s)
		blackhole = total
	}
}

// Бенчмарк суммы каждого восьмого значения
func BenchmarkSum8(b *testing.B) {
	s := makeSlice(size * 8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := Sum8(s)
		blackhole = total
	}
}

// Создание среза пар
func makePairSlice(n int) []Pair {
	res := make([]Pair, n)
	for i := 0; i < n; i++ {
		res[i] = Pair{
			int64(i + 1),
			int64(n - i + 1),
		}
	}
	return res
}

// Бенчмарк суммы среза пар
func BenchmarkSumPair(b *testing.B) {
	pairs := makePairSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := SumPair(pairs)
		blackhole = total
	}
}
