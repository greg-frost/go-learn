package main

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
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
}
