package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// Структура "файл кэша"
type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

// Путь, кэш и мьютекс
var path = os.Getenv("GOPATH") + "/src/learn/"
var cache = map[string]*cacheFile{}
var mutex sync.RWMutex

// Обработка файлов
func serveFiles(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	v, ok := cache[r.URL.Path]
	mutex.RUnlock()

	if !ok {
		mutex.Lock()
		defer mutex.Unlock()

		filename := path + r.URL.Path
		f, err := os.Open(filename)
		defer f.Close()
		if err != nil {
			http.NotFound(w, r)
			return
		}

		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		reader := bytes.NewReader(b.Bytes())
		info, _ := f.Stat()
		v = &cacheFile{
			content: reader,
			modTime: info.ModTime(),
		}
		cache[r.URL.Path] = v
	} else {
		fmt.Println("Найдено в кэше:", r.URL.Path)
	}

	http.ServeContent(w, r, r.URL.Path, v.modTime, v.content)
}

func main() {
	fmt.Println(" \n[ КЭШ-СЕРВЕР ]\n ")

	// Обработчик
	http.HandleFunc("/", serveFiles)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
