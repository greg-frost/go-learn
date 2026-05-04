package main

import (
	"context"
	"fmt"
	"strconv"

	"golang.org/x/sync/errgroup"
)

// Структура "данные"
type Input struct {
	ID string
}

// Структура "результат"
type Result struct {
	ID int
}

// Обработка
func Process(ctx context.Context, inputs []Input) ([]Result, error) {
	// Создание errgroup
	g, ctx := errgroup.WithContext(ctx)
	results := make([]Result, len(inputs))

	// Запуск задач
	for i, input := range inputs {
		i := i
		input := input
		g.Go(func() error {
			result, err := Convert(input)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	// Ожидание результатов
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

// Конвертирование
func Convert(input Input) (Result, error) {
	id, err := strconv.Atoi(input.ID)
	if err != nil {
		return Result{}, err
	}
	return Result{ID: id}, nil
}

func main() {
	fmt.Println(" \n[ ERRGROUP ]\n ")

	// Данные (только числа)
	inputs := []Input{
		{"1"},
		{"2"},
		{"two"},
		{"3"},
		{"4"},
		{"5"},
	}

	// Обработка всей группы до первой ошибки
	results, err := Process(context.Background(), inputs)
	if err != nil {
		fmt.Println("Что-то пошло не так:", err)
		return
	}
	fmt.Println("Результаты:")
	for _, res := range results {
		fmt.Print(res.ID, " ")
	}
	fmt.Println()
}
