package math

import "errors"

// Сумма
func Sum(a int, b int) int {
	return a + b
}

// Произведение
func Prod(a int, b int) int {
	return a * b
}

// Произведение через сумму
func ProdBySum(val int, times int) (res int) {
	if times >= 0 {
		for i := 0; i < times; i++ {
			res += val
		}
	} else {
		for i := 0; i > times; i-- {
			res -= val
		}
	}
	return
}

// Деление
func Divide(num int, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("деление на ноль")
	}
	return num / den, nil
}

// Примерное значение
func Estimate(value int) string {
	switch {
	case value < 10:
		return "small"
	case value < 100:
		return "medium"
	default:
		return "big"
	}
}
