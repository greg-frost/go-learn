package transact

import (
	"fmt"
	"os"
	"path/filepath"

	"go-learn/base"
	"go-learn/cloud/core"
)

// Путь
var path = base.Dir("cloud")

// Конструктор регистратора
func NewTransactionLogger(loggerType string) (core.TransactionLogger, error) {
	switch loggerType {
	case "file":
		return NewFileTransactionLogger(
			filepath.Join(path, os.Getenv("TLOG_FILENAME")),
		)

	case "postgres":
		return NewPostgresTransactionLogger(
			PostgresDBParams{
				dbName:   os.Getenv("TLOG_DB_NAME"),
				host:     os.Getenv("TLOG_DB_HOST"),
				user:     os.Getenv("TLOG_DB_USER"),
				password: os.Getenv("TLOG_DB_PASS"),
			},
		)
	default:
		return nil, fmt.Errorf("регистратора транзакций %s не существует", loggerType)
	}
}
