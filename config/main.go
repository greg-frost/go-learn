package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"go-learn/base"

	"gopkg.in/yaml.v2"
)

// Структура "конфигурация"
type Config struct {
	Host string            `yaml:"host"`
	Port uint16            `yaml:"port"`
	Tags map[string]string `yaml:"tags"`
}

// Конфигурация
var config Config

// Путь
var path = base.Dir("config")

// Загрузка конфигурации
func loadConfig(filename string) (Config, error) {
	d, err := os.ReadFile(filepath.Join(path, filename))
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = yaml.Unmarshal(d, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// Печать конфигурации
func printConfig(config Config) {
	fmt.Println("Host:", config.Host)
	fmt.Println("Port:", config.Port)
	fmt.Println("Tags:")
	for tag, value := range config.Tags {
		fmt.Printf("   %s: %s\n", tag, value)
	}
}

// Вычисление хэша файла
func calculateFileHash(filename string) (string, error) {
	file, err := os.Open(filepath.Join(path, filename))
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	sum := fmt.Sprintf("%x", hash.Sum(nil))

	return sum, nil
}

// Наблюдение за конфигурацией
func watchConfig(filename string) (<-chan string, <-chan error, error) {
	updates := make(chan string)
	errs := make(chan error)
	var hash string

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			newHash, err := calculateFileHash(filename)
			if err != nil {
				errs <- err
				continue
			}

			if newHash != hash {
				hash = newHash
				updates <- filename
			}
		}
	}()

	return updates, errs, nil
}

// Прослушивание изменений
func startListening(updates <-chan string, errs <-chan error) {
	for {
		select {
		case filename := <-updates:
			cfg, err := loadConfig(filename)
			if err != nil {
				fmt.Println("Ошибка загрузки конфигурации:", err)
				continue
			}

			config = cfg

			fmt.Println()
			fmt.Println("Конфигурация изменилась!")
			printConfig(config)

		case err := <-errs:
			fmt.Println("Ошибка наблюдения за конфигурацией:", err)
		}
	}
}

func init() {
	// Регистрация наблюдения
	updates, errs, err := watchConfig("config.yml")
	if err != nil {
		panic(err)
	}

	// Прослушивание изменений
	go startListening(updates, errs)
}

func main() {
	fmt.Println(" \n[ КОНФИГУРАЦИЯ ]\n ")

	fmt.Println("Измените файл конфигурации")
	fmt.Println("или нажмите любую кнопку для выхода...")

	// Ожидание ввода
	var input string
	fmt.Scanln(&input)
}
