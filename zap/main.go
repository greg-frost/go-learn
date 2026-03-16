package main

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
)

func main() {
	fmt.Println(" \n[ ЛОГГЕР ZAP ]\n ")

	// Инициализация
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	// Использование
	logger.Info("не удалось получить URL",
		zap.String("url", "example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
