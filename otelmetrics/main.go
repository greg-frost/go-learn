package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/exporters/metric/prometheus"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/metric"
	// "go.opentelemetry.io/otel"
	// sdkmetric "go.opentelemetry.io/otel/sdk/export/metric"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Имя сервиса
const serviceName = "fibonacci"

// Провайдер метрик
var meterProvider metric.MeterProvider

// Метрика "число запросов"
var requests metric.Int64Counter

// Метки метрик
var labels = []label.KeyValue{
	label.Key("application").String(serviceName),
	label.Key("container_id").String(os.Getenv("HOSTNAME")),
}

// Создание счетчика запросов
func buildRequestsCounter() error {
	// Получение экземпляра из провайдера метрик
	// mp := otel.GetMeterProvider()
	mp := meterProvider
	meter := mp.Meter(serviceName)

	// Создание счетчика метрики
	var err error
	requests, err = meter.NewInt64Counter("fibonacci_requests_total",
		metric.WithDescription("Общее число запросов к сервису Fibonacci"))

	return err
}

// Обновление метрик (вручную)
func watchRuntimeMetrics(ctx context.Context) {
	// Получение экземпляра из провайдера метрик
	// mp := otel.GetMeterProvider()
	mp := meterProvider
	meter := mp.Meter(serviceName)

	// Метрики памяти и горутин
	memory, _ := meter.NewInt64UpDownCounter("memory_usage_bytes",
		metric.WithDescription("Количество использованной памяти"))
	goroutines, _ := meter.NewInt64UpDownCounter("num_goroutines",
		metric.WithDescription("Количество запущенных горутин"))

	// Опрос метрик
	var m runtime.MemStats
	for {
		// Получение метрик
		runtime.ReadMemStats(&m)
		mMemory := memory.Measurement(int64(m.Sys))
		mGoroutines := goroutines.Measurement(int64(runtime.NumGoroutine()))

		// Запись метрик одной пачкой
		meter.RecordBatch(ctx, labels, mMemory, mGoroutines)
		time.Sleep(5 * time.Second)
	}
}

// Обновление метрик (автоматическое)
func buildRuntimeObservers() {
	// Получение экземпляра из провайдера метрик
	// mp := otel.GetMeterProvider()
	mp := meterProvider
	meter := mp.Meter(serviceName)
	var m runtime.MemStats

	// Метрика "использование памяти"
	meter.NewInt64UpDownSumObserver("memory_usage_bytes",
		func(_ context.Context, result metric.Int64ObserverResult) {
			runtime.ReadMemStats(&m)
			result.Observe(int64(m.Sys), labels...)
		},
		metric.WithDescription("Количество использованной памяти"),
	)
	// Метрика "количество горутин"
	meter.NewInt64UpDownSumObserver("num_goroutines",
		func(_ context.Context, result metric.Int64ObserverResult) {
			result.Observe(int64(runtime.NumGoroutine()), labels...)
		},
		metric.WithDescription("Количество запущенных горутин"),
	)
}

// Вычисление числа Фибоначчи
func Fibonacci(ctx context.Context, n int) chan int {
	ch := make(chan int)

	// Увеличение метрики запросов
	requests.Add(ctx, 1, labels...)

	go func() {
		// Вычисление
		res := 1
		if n > 1 {
			a := Fibonacci(ctx, n-1)
			b := Fibonacci(ctx, n-2)

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

	fmt.Fprintln(w, res)
}

func main() {
	fmt.Println(" \n[ OPEN TELEMETRY (МЕТРИКИ) ]\n ")

	// Экспортер Prometheus
	// (можно использовать InstallNewPipeline)
	prometheusExporter, err := prometheus.NewExportPipeline(
		prometheus.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Провайдер метрик
	mp := prometheusExporter.MeterProvider()

	// Установка глобального провайдера
	// otel.SetMeterProvider(mp)
	meterProvider = mp

	// Регистрация счетчика
	buildRequestsCounter()

	// Обновление метрик
	// go watchRuntimeMetrics(ctx)
	buildRuntimeObservers()

	// Обработчики
	http.HandleFunc("/", handleFib)
	http.Handle("/metrics", prometheusExporter)

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
