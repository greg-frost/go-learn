package apiserver

// Конфигурация сервера
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *StoreConfig
}

// Конфигурация хранилища
type StoreConfig struct {
	DatabaseURL string `toml:"database_url"`
}

// Конструктор конфигурации
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    &StoreConfig{},
	}
}
