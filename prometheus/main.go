package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Демо-метрика - счетчик
var counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "demo",
		Name:      "counter",
		Help:      "Счетчик - только увеличивается",
	})

// Демо-метрика - шкала
var gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "demo",
		Name:      "gauge",
		Help:      "Шкала - изменяется произвольно",
	})

// Демо-метрика - гистограмма
var histogram = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: "demo",
		Name:      "histogram",
		Help:      "Гистограмма - группирует значения",
	})

// Демо-метрика - сводка
var summary = prometheus.NewSummary(
	prometheus.SummaryOpts{
		Namespace: "demo",
		Name:      "summary",
		Help:      "Сводка - продвинутая гистограмма",
	})

// Go-метрики
const numRoutines = "/sched/goroutines:goroutines"   // Количество горутин
const totalMemory = "/memory/classes/total:bytes"    // Выделенная память
const freeMemory = "/memory/classes/heap/free:bytes" // Освобожденная память

// Go-метрика - количество горутин
var numRoutinesGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "golang",
		Name:      "num_routines",
		Help:      "Количество горутин",
	})

// Go-метрика - выделенная память
var totalMemoryGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "golang",
		Name:      "total_memory",
		Help:      "Выделенная память",
	})

// Go-метрика - освобожденная память
var freeMemoryGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "golang",
		Name:      "free_memory",
		Help:      "Освобожденная память",
	})

func main() {
	fmt.Println(" \n[ PROMETHEUS ]\n ")

	// Регистрация демо-метрик
	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	// Регистрация go-метрик
	prometheus.MustRegister(numRoutinesGauge)
	prometheus.MustRegister(totalMemoryGauge)
	prometheus.MustRegister(freeMemoryGauge)

	// Срез для метрик
	ms := []metrics.Sample{
		{Name: numRoutines},
		{Name: totalMemory},
		{Name: freeMemory},
	}

	// Демо-метрики
	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)

			time.Sleep(2 * time.Second)
		}
	}()

	// Go-метрики
	go func() {
		for {
			// Создание горутин и выделение памяти
			for i := 0; i < 3; i++ {
				go func() {
					_ = make([]int, 1_000_000)
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				}()
			}

			// Сборка мусора
			runtime.GC()

			// Чтение метрик
			metrics.Read(ms)
			numRoutinesValue := ms[0].Value.Uint64()
			totalMemoryValue := ms[1].Value.Uint64()
			freeMemoryValue := ms[2].Value.Uint64()

			// Передача метрик
			numRoutinesGauge.Set(float64(numRoutinesValue))
			totalMemoryGauge.Set(float64(totalMemoryValue))
			freeMemoryGauge.Set(float64(freeMemoryValue))

			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}()

	// Обработчик метрик
	http.Handle("/metrics", promhttp.Handler())

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
