package main

import (
	"fmt"
	"log"
	"os"
)

// Тип "уровень лога"
type LogLevel int

// Уровни лога
const (
	LogLevelInfo LogLevel = iota + 1
	LogLevelWarning
	LogLevelError
)

// Структура "расширенный логгер"
type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

// Установка уровня лога
func (l *LogExtended) SetLogLevel(level LogLevel) {
	if !level.IsValid() {
		return
	}
	l.logLevel = level
}

// Проверка валидности уровня лога
func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

// Печать информации
func (l *LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO", msg)
}

// Печать предупреждения
func (l *LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN", msg)
}

// Печать ошибки
func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERR", msg)
}

// Печать лога
func (l *LogExtended) println(level LogLevel, prefix, msg string) {
	if level < l.logLevel {
		return
	}
	l.Logger.Println(prefix, msg)
}

// Конструктор расширенного логгера
func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

func main() {
	fmt.Println(" \n[ УРОВНЕВОЕ ЛОГИРОВАНИЕ ]\n ")

	// Создание логгера
	logger := NewLogExtended()

	// Установка предельного уровня
	logger.SetLogLevel(LogLevelWarning)

	// Логирование
	fmt.Println("Логи уровня Warning и выше:")
	fmt.Println()
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
