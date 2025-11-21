package store

// Конфигурация хранилища
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// Конструктор конфигурации
func NewConfig() *Config {
	return &Config{}
}
