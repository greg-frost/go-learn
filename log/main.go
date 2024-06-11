package main

import (
	"fmt"
	"io/ioutil"
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
func (l *LogExtended) SetLogLevel(lvl LogLevel) {
	if !lvl.IsValid() {
		return
	}
	l.logLevel = lvl
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
	l.println(LogLevelInfo, "INFO ", msg)
}

// Печать предупреждения
func (l *LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN ", msg)
}

// Печать ошибки
func (l *LogExtended) Errorln(msg string) {
	l.println(LogLevelError, "ERR ", msg)
}

// Печать лога
func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if srcLogLvl < l.logLevel {
		return
	}

	l.Logger.Println(prefix + msg)
}

// Конструктор расширенного логгера
func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

func main() {
	fmt.Println(" \n[ ЛОГИРОВАНИЕ ]\n ")

	// Стандартный
	fmt.Println("Стандартный логгер:")
	log.Printf("Hello\n")
	log.Println("World")
	fmt.Println()

	// Настроенный
	fmt.Println("Настроенный логгер:")
	flags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stderr, "stderr ", flags)
	logger.Printf("Hello\n")
	logger.Println("World")
	fmt.Println()

	// Файловый
	fmt.Println("Файловый логгер:")
	f, err := os.CreateTemp("", "temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())
	file := log.New(f, "file ", flags)
	file.Printf("Hello\n")
	file.Println("World")
	bs, err := ioutil.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	// Кастомный
	fmt.Println("Кастомный логгер:")
	custom := NewLogExtended()
	custom.SetLogLevel(LogLevelWarning)
	custom.Infoln("Не должно напечататься")
	custom.Warnln("Hello")
	custom.Errorln("World")
	custom.Println("Debug")
}
