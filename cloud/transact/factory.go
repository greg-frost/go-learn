package transact

import (
	"fmt"
	"os"

	"go-learn/cloud/core"
)

// Конструктор регистратора
func NewTransactionLogger(loggerType string) (core.TransactionLogger, error) {
	switch loggerType {
	case "file":
		return NewFileTransactionLogger(
			os.Getenv("TLOG_FILE"),
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
		return nil, fmt.Errorf("нет регистратора транзакций %s", loggerType)
	}
}
