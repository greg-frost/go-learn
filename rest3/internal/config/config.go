package config

import (
	"path/filepath"
	"sync"

	"go-learn/base"
	"go-learn/rest3/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

// Структура "конфигурация"
type Config struct {
	IsDebug *bool `yaml:"is_debug" env:"IS_DEBUG" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env:"TYPE" env-default:"port"`
		BindIP string `yaml:"bind_ip" env:"BIND_IP" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env:"PORT" env-default:"8080"`
	} `yaml:"listen" env:"LISTEN"`
	MongoDB struct {
		Host       string `yaml:"host" env:"HOST"`
		Port       string `yaml:"port" env:"PORT"`
		Database   string `yaml:"database" env:"DATABASE"`
		AuthDB     string `yaml:"auth_db" env:"AUTH_DB"`
		Username   string `yaml:"username" env:"USERNAME"`
		Password   string `yaml:"password" env:"PASSWORD"`
		Collection string `yaml:"collection" env:"COLLECTION"`
	} `yaml:"mongodb" env:"MONGODB"`
}

// Экземпляр (синглтон)
var instance *Config

// Однократное выполнение
var once sync.Once

// Путь
var path = base.Dir("rest3")

// Конструктор
func New() *Config {
	once.Do(func() {
		log := logger.New()
		log.Info("Чтение конфигурации приложения")

		instance = new(Config)
		if err := cleanenv.ReadConfig(filepath.Join(path, "config.yaml"), instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info(help)
			log.Fatal(err)
		}
	})

	return instance
}
