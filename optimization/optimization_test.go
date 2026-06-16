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

// Создание связного списка
func makeLinkedList(n int) *Node {
	root := new(Node)
	curr := root
	for i := 0; i < n; i++ {
		curr.Next = &Node{
			Value: int64(i + 1),
		}
		curr = curr.Next
	}
	return root.Next
}

// Бенчмарк суммы значений связного списка
func BenchmarkSumLinkedList(b *testing.B) {
	node := makeLinkedList(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := SumLinkedList(node)
		blackhole = total
	}
}

// Создание среза пар
func makeSliceOfPair(n int) []Pair {
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
	pairs := makeSliceOfPair(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := SumPair(pairs)
		blackhole = total
	}
}

// Создание пар срезов
func makePairsOfSlices(n int) Pairs {
	res := Pairs{
		make([]int64, n),
		make([]int64, n),
	}
	for i := 0; i < n; i++ {
		res.a[i] = int64(i + 1)
		res.b[i] = int64(n - i + 1)
	}
	return res
}

// Бенчмарк суммы пар срезов
func BenchmarkSumPairs(b *testing.B) {
	pairs := makePairsOfSlices(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := SumPairs(pairs)
		blackhole = total
	}
}

// Создание двумерного среза (512 столбцов)
func makeSlice512(n int) [][512]int64 {
	res := make([][512]int64, n)
	for i := 0; i < n; i++ {
		for j := 0; i < 512; i++ {
			res[i][j] = int64(i + j + 1)
		}
	}
	return res
}

// Бенчмарк суммы первых значений рядов (512 столбцов)
func BenchmarkSumRows512(b *testing.B) {
	s := makeSlice512(size / 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := SumRows512(s)
		blackhole = total
	}
}

// Создание двумерного среза (513 столбцов)
func makeSlice513(n int) [][513]int64 {
	res := make([][513]int64, n)
	for i := 0; i < n; i++ {
		for j := 0; i < 513; i++ {
			res[i][j] = int64(i + j + 1)
		}
	}
	return res
}

// Бенчмарк суммы первых значений рядов (513 столбцов)
func BenchmarkSumRows513(b *testing.B) {
	s := makeSlice513(size / 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := SumRows513(s)
		blackhole = total
	}
}

// Создание среза входных данных
func makeSliceOfInput(n int) []Input {
	res := make([]Input, n)
	for i := 0; i < n; i++ {
		res[i] = Input{
			int64(i + 1),
			int64(n - i + 1),
		}
	}
	return res
}

// Бенчмарк подсчета сумм
func BenchmarkCount(b *testing.B) {
	inputs := makeSliceOfInput(size / 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := Count(inputs)
		blackhole = result.sumA
		blackhole = result.sumB
	}
}

// Бенчмарк подсчета сумм (оптимизированного)
func BenchmarkCountOptimized(b *testing.B) {
	inputs := makeSliceOfInput(size / 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := CountOptimized(inputs)
		blackhole = result.sumA
		blackhole = result.sumB
	}
}

// Бенчмарк инкремента
func BenchmarkIncrement(b *testing.B) {
	var s [2]int64
	for i := 0; i < b.N; i++ {
		s = Increment(s, size/10)
		blackhole = s[0]
		blackhole = s[1]
	}
}

// Бенчмарк инкремента (оптимизированного)
func BenchmarkIncrementOptimized(b *testing.B) {
	var s [2]int64
	for i := 0; i < b.N; i++ {
		s = IncrementOptimized(s, size/10)
		blackhole = s[0]
		blackhole = s[1]
	}
}

// Создание среза полей
func makeSliceOfField(n int) []Field {
	res := make([]Field, n)
	for i := 0; i < n; i++ {
		res[i] = Field{
			i: int64(i + 1),
		}
	}
	return res
}

// Бенчмарк суммы полей
func BenchmarkSumFields(b *testing.B) {
	fields := makeSliceOfField(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := SumFields(fields)
		blackhole = total
	}
}

// Создание среза полей (выровненного)
func makeSliceOfFieldAligned(n int) []FieldAligned {
	res := make([]FieldAligned, n)
	for i := 0; i < n; i++ {
		res[i] = FieldAligned{
			i: int64(i + 1),
		}
	}
	return res
}

// Бенчмарк суммы полей (оптимизированной)
func BenchmarkSumFieldsOptimized(b *testing.B) {
	fields := makeSliceOfFieldAligned(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total := SumFieldsOptimized(fields)
		blackhole = total
	}
}
