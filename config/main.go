package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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
// var config Config

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

// Вычисление кэша файла
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

func main() {
	fmt.Println(" \n[ КОНФИГУРАЦИЯ ]\n ")

	// Загрузка
	cfg, err := loadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Вывод
	fmt.Println("Host:", cfg.Host)
	fmt.Println("Port:", cfg.Port)
	fmt.Println("Tags:")
	for tag, value := range cfg.Tags {
		fmt.Printf("   %s: %s\n", tag, value)
	}
}
