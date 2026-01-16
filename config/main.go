package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"go-learn/base"

	"github.com/fsnotify/fsnotify"
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

// Наблюдение за конфигурацией (по изменению хэша)
func watchConfigByHash(filename string) (<-chan string, <-chan error, error) {
	updates := make(chan string)
	errs := make(chan error)
	hash, _ := calculateFileHash(filename)

	go func() {
		ticker := time.NewTicker(3 * time.Second)
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

// Наблюдение за конфигурацией (по сигналам ОС)
func watchConfigByNotify(filename string) (<-chan string, <-chan error, error) {
	updates := make(chan string)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, nil, err
	}

	err = watcher.Add(filepath.Join(path, filename))
	if err != nil {
		return nil, nil, err
	}

	go func() {
		// Начальное изменение файла
		// updates <- filename

		for event := range watcher.Events {
			// Почему-то приходит по два события
			if event.Op&fsnotify.Write == fsnotify.Write {
				updates <- filepath.Base(event.Name)
			}
		}
	}()

	return updates, watcher.Errors, nil
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
	// Регистрация наблюдателя
	updates, errs, err := watchConfigByHash("config.yml")
	if err != nil {
		panic(err)
	}

	// Прослушивание изменений
	go startListening(updates, errs)
}

func main() {
	fmt.Println(" \n[ КОНФИГУРАЦИЯ ]\n ")

	// Загрузка
	var err error
	config, err = loadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Печать
	printConfig(config)

	fmt.Println()
	fmt.Println("Измените файл конфигурации")
	fmt.Println("или нажмите любую кнопку для выхода...")

	// Ожидание ввода
	var input string
	fmt.Scanln(&input)
}
