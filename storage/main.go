package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"go-learn/base"
)

// Интерфейс "файл"
type File interface {
	Load(path string) (io.ReadCloser, error)
	Save(path string, body io.ReadSeeker) error
}

// Общие ошибки
var (
	ErrFileNotFound   = errors.New("Файл не найден")
	ErrCannotLoadFile = errors.New("Не удалось загрузить файл")
	ErrCannotSaveFile = errors.New("Не удалось сохранить файл")
)

// Путь
var path = base.Dir("storage")

// Структура "локальный файл"
type LocalFile struct {
	Base string
}

// Конструктор хранилища
func newFileStore() (File, error) {
	return &LocalFile{Base: path}, nil
}

// Загрузка
func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)

	var e error
	o, err := os.Open(p)
	if err != nil && os.IsNotExist(err) {
		log.Printf("Не удалось обнаружить %s", p)
		e = ErrFileNotFound
	} else if err != nil {
		log.Printf("Ошибка при загрузке файла %s: %s", p, err)
		e = ErrCannotLoadFile
	}

	return o, e
}

// Сохранение
func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)

	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		log.Printf("Не удалось создать каталог %s", d)
		return ErrCannotSaveFile
	}

	f, err := os.Create(p)
	if err != nil {
		log.Printf("Не удалось создать файл %s", p)
		return ErrCannotSaveFile
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	if err != nil {
		log.Print("Не удалось скопировать текст в новый файл")
		return ErrCannotSaveFile
	}

	return nil
}

func main() {
	fmt.Println(" \n[ ХРАНИЛИЩЕ ]\n ")

	// Текст
	content := "Простой текст непростого человека..."
	body := bytes.NewReader([]byte(content))
	store, err := newFileStore()
	if err != nil {
		log.Fatal(err)
	}

	// Сохранение
	fmt.Println("Сохранение...")
	filename := filepath.Join("data", "file.txt")
	err = store.Save(filename, body)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(filepath.Join(path, filepath.Dir(filename)))

	// Загрузка
	fmt.Println("Загрузка...")
	c, err := store.Load(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	o, err := ioutil.ReadAll(c)
	if err != nil {
		log.Fatal(err)
	}

	// Вывод и проверка
	s := string(o)
	fmt.Println("Тексты идентичны:", content == s)
}
