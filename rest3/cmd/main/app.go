package main

import (
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"time"

	"go-learn/base"
	"go-learn/rest3/internal/config"
	"go-learn/rest3/internal/user"
	"go-learn/rest3/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

// Путь
var path = base.Dir("rest3")

func main() {
	fmt.Println(" \n[ REST 3 (THE ART OF DEVELOPMENT) ]\n ")

	// Получение логгера
	log := logger.New()

	// Получение конфигурации
	cfg := config.New()

	// Создание роутера
	log.Info("Создание роутера")
	router := httprouter.New()

	// Регистрация обработчиков
	log.Info("Регистрация обработчиков")
	handler := user.NewHandler(log)
	handler.Register(router)

	// Запуск сервера
	startServer(router, cfg)
}

// Запуск сервера
func startServer(router *httprouter.Router, cfg *config.Config) {
	log := logger.New()
	log.Info("Запуск сервера")

	var listener net.Listener
	var caption string
	var err error

	// Прослушивание соединений
	if cfg.Listen.Type == "sock" { // Сокет
		log.Info("Создание сокета")
		socketPath := filepath.Join(path, "cmd", "main", "app.sock")
		log.Debugf("Путь к сокету: %s", socketPath)

		log.Info("Прослушивание сокета")
		listener, err = net.Listen("unix", socketPath)

		caption = fmt.Sprintf("(на unix socket: %s)", socketPath)
	} else { // Порт
		log.Info("Прослушивание порта")
		listener, err = net.Listen("tcp",
			fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))

		caption = fmt.Sprintf("(на http://%s:%s)", cfg.Listen.BindIP, cfg.Listen.Port)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Настройка
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Запуск
	log.Info("Ожидаю обновлений...")
	log.Info(caption)
	log.Fatal(server.Serve(listener))
}
