package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Структура "тестовая пара"
type testPair struct {
	values [2]int
	result int
}

// Тестовые пары для сложения
var sumTests = []testPair{
	{[2]int{1, 1}, 2},
	{[2]int{5, 12}, 17},
	{[2]int{3, 92}, 95},
	{[2]int{-3, -10}, -13},
	{[2]int{-5, 10}, 5},
}

// Тестовые пары для умножения
var prodTests = []testPair{
	{[2]int{2, 4}, 8},
	{[2]int{2, 12}, 24},
	{[2]int{3, 100}, 300},
	{[2]int{-5, -6}, 30},
	{[2]int{-10, 10}, -100},
}

/* Тесты */

// Тест суммы
func TestSum(t *testing.T) {
	for _, pair := range sumTests {
		g := Sum(pair.values[0], pair.values[1])
		e := pair.result

		if g != e {
			t.Error("Сумма: получено", g, ", ожидается", e)
		}
	}
}

// Тест произведения
func TestProd(t *testing.T) {
	for _, pair := range prodTests {
		g := Prod(pair.values[0], pair.values[1])
		e := pair.result

		if g != e {
			t.Error("Произведение: получено", g, ", ожидается", e)
		}
	}
}

// Тест произведения через сумму
func TestProdBySum(t *testing.T) {
	for _, pair := range prodTests {
		g := ProdBySum(pair.values[0], pair.values[1])
		e := pair.result

		if g != e {
			t.Error("Произведение через сумму: получено", g, ", ожидается", e)
		}
	}
}

/* Testify и Gotests */

// Тест деления
func TestDivide(t *testing.T) {
	t.Run("ZeroNumerator", func(t *testing.T) {
		result, err := Divide(0, 1)
		require.NoError(t, err)
		assert.Equal(t, 0, result)
	})

	t.Run("BothNonZero", func(t *testing.T) {
		result, err := Divide(4, 2)
		require.NoError(t, err)
		assert.Equal(t, 2, result)
	})

	t.Run("ZeroDenominator", func(t *testing.T) {
		_, err := Divide(1, 0)
		require.Error(t, err)
	})
}

// Тест деления (table-driven)
func TestDivideTD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "ZeroNumerator",
			args: args{
				a: 0,
				b: 1,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "BothNonZero",
			args: args{
				a: 4,
				b: 2,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "ZeroDenominator",
			args: args{
				a: 1,
				b: 0,
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Тест примерного значения
func TestEstimate(t *testing.T) {
	t.Run("Small", func(t *testing.T) {
		assert.Equal(t, "small", Estimate(5))
	})

	t.Run("Medium", func(t *testing.T) {
		assert.Equal(t, "medium", Estimate(50))
	})

	t.Run("Big", func(t *testing.T) {
		assert.Equal(t, "big", Estimate(100))
	})
}

// Тест примерного значения (table-driven)
func TestEstimateTD(t *testing.T) {
	testCases := []struct {
		Name          string
		InputValue    int
		ExpectedValue string
	}{
		{
			Name:          "Small",
			InputValue:    5,
			ExpectedValue: "small",
		},
		{
			Name:          "Medium",
			InputValue:    50,
			ExpectedValue: "medium",
		},
		{
			Name:          "Big",
			InputValue:    100,
			ExpectedValue: "big",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.EqualValues(t, tc.ExpectedValue, Estimate(tc.InputValue))
		})
	}
}
