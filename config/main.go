package main

import (
	"fmt"
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

func main() {
	fmt.Println(" \n[ КОНФИГУРАЦИЯ ]\n ")

	cfg, err := loadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Host:", cfg.Host)
	fmt.Println("Port:", cfg.Port)
	fmt.Println("Tags:")
	for tag, value := range cfg.Tags {
		fmt.Printf("   %s: %s\n", tag, value)
	}
}
