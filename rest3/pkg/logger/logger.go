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

var path = base.Dir("rest3")

// Инициализация логгера
func Init() {
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
	allFile, err := os.OpenFile(filepath.Join(path, "logs/all.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	// Отключение вывода
	l.SetOutput(io.Discard)

	fmt.Println(allFile)
}
