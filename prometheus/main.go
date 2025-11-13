package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Метрика - счетчик
var counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "golearn",
		Name:      "counter",
		Help:      "Счетчик - только увеличивается",
	})

// Метрика - шкала
var gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "golearn",
		Name:      "gauge",
		Help:      "Шкала - изменяется произвольно",
	})

// Метрика - гистограмма
var histogram = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: "golearn",
		Name:      "histogram",
		Help:      "Гистограмма - группирует значения",
	})

// Метрика - сводка
var summary = prometheus.NewSummary(
	prometheus.SummaryOpts{
		Namespace: "golearn",
		Name:      "summary",
		Help:      "Сводка - продвинутая гистограмма",
	})

func main() {
	fmt.Println(" \n[ PROMETHEUS ]\n ")

	// Регистрация метрик
	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	// Передача метрик
	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)
			time.Sleep(2 * time.Second)
		}
	}()

	// Обработчик метрик
	http.Handle("/metrics", promhttp.Handler())

	// Запуск сервера
	fmt.Println("Ожидаю соединений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
