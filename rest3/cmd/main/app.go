package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"time"

	"go-learn/base"
	"go-learn/rest3/internal/config"
	"go-learn/rest3/internal/user"
	"go-learn/rest3/internal/user/db"
	"go-learn/rest3/pkg/client/mongodb"
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

	// БД (MongoDB)
	mongodbClient, err := mongodb.NewClient(
		context.Background(),
		cfg.MongoDB.Host,
		cfg.MongoDB.Port,
		cfg.MongoDB.Username,
		cfg.MongoDB.Password,
		cfg.MongoDB.Database,
		cfg.MongoDB.AuthDB,
	)
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(mongodbClient, cfg.MongoDB.Collection, log)

	u := user.User{
		Username:     "greg_frost",
		Email:        "noreply@example.com",
		PasswordHash: "1234567890abcdef",
	}

	uID, err := storage.Create(context.Background(), u)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Создание пользователя:", uID)

	one, err := storage.FindOne(context.Background(), uID)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Поиск пользователя:", one)

	one.Email = "reply@example.com"
	err = storage.Update(context.Background(), one)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Обновление пользователя:", one)

	err = storage.Delete(context.Background(), uID)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Удаление пользователя:", uID)

	one, err = storage.FindOne(context.Background(), uID)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Поиск пользователя:", one)

	count := 3
	for i := 0; i < count; i++ {
		storage.Create(context.Background(), u)
	}
	log.Info("Создание пользователей:", count)

	all, err := storage.FindAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Поиск всех пользователей:", all)

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
		serverAddr := fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
		listener, err = net.Listen("tcp", serverAddr)
		caption = fmt.Sprintf("(на http://%s)", serverAddr)
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
