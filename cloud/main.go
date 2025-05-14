package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go-learn/base"
	"go-learn/cloud/core"
	"go-learn/cloud/frontend"
	"go-learn/cloud/transact"
)

// Путь
var path = base.Dir("cloud")

func init() {
	os.Setenv("TLOG_TYPE", "file")
	os.Setenv("FRONTEND_TYPE", "rest")

	os.Setenv("TLOG_FILE", filepath.Join(
		path, "data", "transaction.log"))

	os.Setenv("TLOG_DB_NAME", "learn")
	os.Setenv("TLOG_DB_HOST", "localhost")
	os.Setenv("TLOG_DB_USER", "postgres")
	os.Setenv("TLOG_DB_PASS", "admin")
}

func main() {
	fmt.Println(" \n[ GO CLOUD ]\n ")

	// Регистратор транзакций
	fmt.Println("Регистратор транзакций:", os.Getenv("TLOG_TYPE"))
	tl, err := transact.NewTransactionLogger(os.Getenv("TLOG_TYPE"))
	if err != nil {
		log.Fatal(err)
	}

	// Хранилище пар ключ-значение
	store := core.NewKeyValueStore(tl)
	restored, _ := store.Restore()

	fmt.Printf("Восстановлено транзакций: %d\n\n", restored)

	// Фронтэнд
	fmt.Println("Фронтэнд:", os.Getenv("FRONTEND_TYPE"))
	fe, err := frontend.NewFrontEnd(os.Getenv("FRONTEND_TYPE"))
	if err != nil {
		log.Fatal(err)
	}

	// Запуск сервера
	log.Fatal(fe.Start(store))
}
