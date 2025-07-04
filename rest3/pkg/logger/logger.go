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
func (h *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range h.Writer {
		w.Write([]byte(line))
	}
	return nil
}

// Подучение уровней логирования
func (h *WriterHook) Levels() []logrus.Level {
	return h.LogLevels
}

// Экземпляр (синглтон)
var e *logrus.Entry

// Структура "логгер"
type Logger struct {
	*logrus.Entry
}

// Конструктор
func New() *Logger {
	return &Logger{e}
}

// Конструктор с полем
func (l *Logger) NewWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

// Инициализация
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
	// err := os.MkdirAll(filepath.Join(path, "logs"), 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// Файл
	// file, err := os.OpenFile(filepath.Join(path, "logs/all.log"),
	// 	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	// if err != nil {
	// 	panic(err)
	// }

	// Отключение дефолтного вывода
	l.SetOutput(io.Discard)

	// Мультирайтер ...
	// multi := io.MultiWriter(os.Stdout, file)
	// l.SetOutput(multi)

	// ... или хуки
	l.AddHook(&WriterHook{
		Writer: []io.Writer{
			os.Stdout,
			// file,
		},
		LogLevels: logrus.AllLevels,
	})

	// Уровень логирования
	l.SetLevel(logrus.TraceLevel)

	// Сохранение логгера
	e = logrus.NewEntry(l)
}
