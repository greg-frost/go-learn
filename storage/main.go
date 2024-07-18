package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Путь
var path = os.Getenv("GOPATH") + "/src/golearn/storage"

// Интерфейс "файл"
type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

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
	return os.Open(p)
}

// Сохранение
func (l LocalFile) Save(path string, body io.ReadSeeker) error {
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)

	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	return err
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
	filename := "data/file.txt"
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
