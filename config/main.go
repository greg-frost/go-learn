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
func loadConfiguration(filename string) (Config, error) {
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
func printConfiguration(config Config) {
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

func main() {
	fmt.Println(" \n[ КОНФИГУРАЦИЯ ]\n ")

	// Загрузка
	cfg, err := loadConfiguration("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Печать
	printConfiguration(cfg)
}
