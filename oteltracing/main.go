package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

const (
	serviceName = "fibonacci" // Имя сервиса
	// jaegerEndpoint = "http://localhost:14268/api/traces"
)

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

			// С отменой
			select {
			case x := <-a:
				select {
				case y := <-b:
					res = x + y
				case <-ctx.Done():
					return
				}
			case y := <-b:
				select {
				case x := <-a:
					res = x + y
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}

		// Добавление атрибута
		sp.SetAttributes(label.Int("result", res))

		ch <- res
	}()

	return ch
}

// Обработчик Фибоначии
func handleFib(w http.ResponseWriter, r *http.Request) {
	var n int
	var err error

	// Парсинг параметра
	queryN := r.URL.Query()["n"]
	if len(queryN) != 1 {
		err = errors.New("неверное число аргументов")
	} else {
		n, err = strconv.Atoi(queryN[0])
	}
	if err != nil {
		http.Error(w, "не удалось распознать параметр n", 400)
		return
	}

	// Вычисление Фибоначчи
	ctx := r.Context()
	res := <-Fibonacci(ctx, n)

	// Трассировка
	// (получение спана из контекста)
	if sp := trace.SpanFromContext(ctx); sp != nil {
		// Добавление атрибутов
		sp.SetAttributes(
			label.Int("parameter", n),
			label.Int("result", res),
		)
	}

	fmt.Fprintln(w, res)
}

func main() {
	fmt.Println(" \n[ OPEN TELEMETRY (ТРАССИРОВКА) ]\n ")

	// Стандартный экспортер
	stdExporter, err := stdout.NewExporter(
		stdout.WithPrettyPrint(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Экспортер Jaeger
	// jaegerExporter, err := jaeger.NewRawExporter(
	// 	jaeger.WithCollectorEndpoint(jaegerEndpoint),
	// 	jaeger.WithProcess(jaeger.Process{
	// 		ServiceName: serviceName,
	// 	}),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Провайдер трассировки
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(stdExporter),
		// sdktrace.WithSyncer(jaegerExporter),
	)

	// Регистрация провайдера
	otel.SetTracerProvider(tp)

	// Обработчик (с инструментацией)
	http.Handle("/",
		otelhttp.NewHandler(http.HandlerFunc(handleFib), "root"),
	)

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
