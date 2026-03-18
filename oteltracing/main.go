package main

import (
	"context"
	"fmt"

	// "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/exporters/stdout"
	// "go.opentelemetry.io/otel/exporters/trace/jaeger"
	// "go.opentelemetry.io/otel/label"
	// export "go.opentelemetry.io/otel/sdk/export/trace"
	// sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

// Имя сервиса
const serviceName = "fibonacci"

// Вычисление числа Фибоначчи
func Fibonacci(ctx context.Context, n int) chan int {
	ch := make(chan int)

	go func() {
		// Получение провайдера трассировки
		tr := otel.GetTracerProvider().Tracer(serviceName)

		// Начало трассировки (спана)
		cctx, sp := tr.Start(ctx,
			fmt.Sprintf("Fibonacci(%d)", n),
			trace.WithAttributes(label.Int("n", n)))
		defer sp.End()

		// Вычисление
		res := 1
		if n > 1 {
			a := Fibonacci(cctx, n-1)
			b := Fibonacci(cctx, n-2)
			res = <-a + <-b
		}

		// Добавление атрибута
		sp.SetAttributes(label.Int("result", res))

		ch <- res
	}()

	return ch
}

func main() {
	fmt.Println(" \n[ OPEN TELEMETRY (ТРАССИРОВКА) ]\n ")
}
