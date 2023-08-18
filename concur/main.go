package main

import (
	"context"
	"fmt"
	"time"
)

// Структура "входные данные"
type Input struct {
	A int
	B int
}

// Типы каналов A, B
type AOut int
type BOut int

// Типы канала C
type CIn struct {
	A int
	B int
}
type COut struct {
	A int
	B int
}

// Получение результата A
func getResultA(ctx context.Context, val int) (AOut, error) {
	time.Sleep(time.Millisecond * 45)
	return AOut(val * 1), nil
}

// Получение результата B
func getResultB(ctx context.Context, val int) (BOut, error) {
	time.Sleep(time.Millisecond * 35)
	return BOut(val * 2), nil
}

// Получение результата C
func getResultC(ctx context.Context, val CIn) (COut, error) {
	time.Sleep(time.Millisecond * 5)
	return COut{
		A: val.A * 3,
		B: val.B * 5,
	}, nil
}

// Структура "процессор"
type processor struct {
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC  chan CIn
	errs chan error
}

// Основная функция управления
func GatherAndProcess(ctx context.Context) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*50)
	defer cancel()

	data, ok := InputFromContext(ctx)
	if !ok {
		return COut{}, fmt.Errorf("Не найдено значение в контексте")
	}

	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		inC:  make(chan CIn, 1),
		outC: make(chan COut, 1),
		errs: make(chan error, 2),
	}
	p.launch(ctx, data)

	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return COut{}, err
	}

	p.inC <- inputC
	out, err := p.waitForC(ctx)
	return out, err
}

// Запуск всех обработчиков
func (p *processor) launch(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- aOut
	}()

	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()

	go func() {
		select {
		case <-ctx.Done():
			return
		case inputC := <-p.inC:
			cOut, err := getResultC(ctx, inputC)
			if err != nil {
				p.errs <- err
				return
			}
			p.outC <- cOut
		}
	}()
}

// Ожидание ответов из A и B
func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
	var inputC CIn
	count := 0

	for count < 2 {
		select {
		case a := <-p.outA:
			inputC.A = int(a)
			count++
		case b := <-p.outB:
			inputC.B = int(b)
			count++
		case err := <-p.errs:
			return CIn{}, err
		case <-ctx.Done():
			return CIn{}, ctx.Err()
		}
	}

	return inputC, nil
}

// Ожидание ответа из C
func (p *processor) waitForC(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}

const ctxKey string = "input"

// Запись значения в контекст
func ContextWithInput(ctx context.Context, input Input) context.Context {
	return context.WithValue(ctx, ctxKey, input)
}

// Чтение значения из контекста
func InputFromContext(ctx context.Context) (Input, bool) {
	input, ok := ctx.Value(ctxKey).(Input)
	return input, ok
}

func main() {
	fmt.Println(" \n[ КОНКУРЕНТНОСТЬ ]\n ")

	ctx := context.Background()
	input := Input{1, 2}
	ctx = ContextWithInput(ctx, input)

	res, err := GatherAndProcess(ctx)
	if err != nil {
		fmt.Println("ОШИБКА:", err)
		return
	}

	fmt.Println("Результат:", res)
}
