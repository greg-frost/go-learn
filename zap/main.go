package main

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	fmt.Println(" \n[ ЛОГГЕР ZAP ]\n ")

	// Стандартный логгер
	// (типизированный, строгий, быстрый)
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Стандартный:")
	logger.Info("не удалось получить URL",
		zap.String("url", "example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	fmt.Println()

	// Расширенный логгер
	// (удобный, "сахарный", медленный)
	sugar := logger.Sugar()

	fmt.Println("Расширенный:")
	sugar.Infow("не удалось получить URL",
		"url", "example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	fmt.Println()

	// Динамическая фильтрация
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = "" // Без меток времени
	cfg.Sampling = &zap.SamplingConfig{
		Initial:    3, // Регистрация до 3 событий в секунду
		Thereafter: 3, // Дальше - только 1 событие из 3
		Hook: func(e zapcore.Entry, d zapcore.SamplingDecision) {
			if d == zapcore.LogDropped {
				fmt.Println("(событие отброшено)")
			}
		},
	}
	logger, _ = cfg.Build()    // Инициализация из конфиругации
	zap.ReplaceGlobals(logger) // Замена глобального логгера

	fmt.Println("Фильтрация:")
	for i := 1; i <= 10; i++ {
		zap.S().Infow(
			"Событие",
			"event", i,
		)
	}
}
