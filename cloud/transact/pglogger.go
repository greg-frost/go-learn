package transact

import (
	"database/sql"
	"fmt"

	"go-learn/cloud/core"
)

// Структура "регистратор транзакций в БД"
type PostgresTransactionLogger struct {
	events chan<- core.Event
	errors <-chan error
	db     *sql.DB
}

// Структура "параметры подключения к БД"
type PostgresDBParams struct {
	dbName   string
	host     string
	user     string
	password string
}

// Размер буфера событий
const PostgresEventsCapacity = 16

// Проверка наличия необходимых таблиц
func (l *PostgresTransactionLogger) verifyTablesExists() (bool, error) {
	query := `SELECT EXISTS (SELECT true FROM information_schema.tables
			  WHERE table_schema = 'public'	AND table_name = 'transactions')`

	row := l.db.QueryRow(query)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, fmt.Errorf("ошибка SQL-запроса: %w", err)
	}

	return exists, nil
}

// Создание необходимых таблиц
func (l *PostgresTransactionLogger) createTables() error {
	query := `CREATE TABLE IF NOT EXISTS transactions
			  (sequence SERIAL PRIMARY KEY, event_type INT, 
			  key TEXT, value TEXT)`

	_, err := l.db.Exec(query)

	return err
}

// Запись транзакции добавления
func (l *PostgresTransactionLogger) WritePut(key, value string) {
	l.events <- core.Event{EventType: core.EventPut, Key: key, Value: value}
}

// Запись транзакции удаления
func (l *PostgresTransactionLogger) WriteDelete(key string) {
	l.events <- core.Event{EventType: core.EventDelete, Key: key}
}

// Получение канала ошибок
func (l *PostgresTransactionLogger) Err() <-chan error {
	return l.errors
}

// Запуск регистратора
func (l *PostgresTransactionLogger) Run() {
	events := make(chan core.Event, PostgresEventsCapacity)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		query := `INSERT INTO transactions
				  (event_type, key, value)
				  VALUES ($1, $2, $3)`

		for e := range events {
			_, err := l.db.Exec(query,
				e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- err
				return
			}
		}
	}()
}

// Чтение событий
func (l *PostgresTransactionLogger) Read() (<-chan core.Event, <-chan error) {
	events := make(chan core.Event)
	errors := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errors)

		query := `SELECT sequence, event_type, key, value
			  	  FROM transactions ORDER BY sequence`

		rows, err := l.db.Query(query)
		if err != nil {
			errors <- fmt.Errorf("ошибка SQL-запроса: %w", err)
			return
		}
		defer rows.Close()

		var e core.Event

		for rows.Next() {
			err = rows.Scan(&e.Sequence, &e.EventType, &e.Key, &e.Value)
			if err != nil {
				errors <- fmt.Errorf("ошибка чтения строки SQL-ответа: %w", err)
				return
			}

			events <- e
		}

		if err := rows.Err(); err != nil {
			errors <- fmt.Errorf("ошибка SQL-ответа: %w", err)
			return
		}
	}()

	return events, errors
}

// Конструктор регистратора
func NewPostgresTransactionLogger(config PostgresDBParams) (core.TransactionLogger, error) {
	conn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		config.host, config.dbName, config.user, config.password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к БД: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("не удалось установить соединение с БД: %w", err)
	}

	logger := &PostgresTransactionLogger{db: db}

	exists, err := logger.verifyTablesExists()
	if err != nil {
		return nil, fmt.Errorf("не удалось проверить наличие таблиц: %w", err)
	}
	if !exists {
		if err := logger.createTables(); err != nil {
			return nil, fmt.Errorf("не удалось создать таблицы: %w", err)
		}
	}

	return logger, nil
}
