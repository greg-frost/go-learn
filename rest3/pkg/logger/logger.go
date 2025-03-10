package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"go-learn/base"

	"github.com/sirupsen/logrus"
)

// Путь
var path = base.Dir("rest3")

// Структура "хук райтера"
type WriterHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Запись логов в райтеры
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return nil
}

// Подучение уровней логирования
func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// Запись логгера
var e *logrus.Entry

// Структура "логгер"
type Logger struct {
	*logrus.Entry
}

// Получение логгера
func GetLogger() Logger {
	return Logger{e}
}

// Получение логгера с полем
func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

// Инициализация логгера
func init() {
	// Создание
	l := logrus.New()

	// Настройка
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := filepath.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function),
				fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: true,
		FullTimestamp: true,
	}

	// Папка
	err := os.MkdirAll(filepath.Join(path, "logs"), 0644)
	if err != nil {
		panic(err)
	}

	// Файл
	file, err := os.OpenFile(filepath.Join(path, "logs/all.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	// Отключение дефолтного вывода
	l.SetOutput(io.Discard)

	// Мультирайтер ...
	// multi := io.MultiWriter(os.Stdout, file)
	// l.SetOutput(multi)

	// ... или хуки
	l.AddHook(&WriterHook{
		Writer:    []io.Writer{file, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	// Уровень логирования
	l.SetLevel(logrus.TraceLevel)

	// Сохранение логгера
	e = logrus.NewEntry(l)
}
