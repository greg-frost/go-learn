package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Путь и шаблон
var path = os.Getenv("GOPATH") + "/src/golearn/upload/"
var t = template.Must(template.ParseFiles(path + "form.html"))

// Обработчик загрузки
func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t.Execute(w, nil)
		return
	}

	in, h, err := r.FormFile("file")
	if err != nil {
		fmt.Fprint(w, "Файл не был загружен...")
		return
	}
	defer in.Close()

	filename := path + h.Filename
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	io.Copy(out, in)

	contentType := typeByHeader(h)
	// contentType := typeByExt(h.Filename)
	// contentType := typeByContent(in)

	w.Header().Set("Content-Type", contentType)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)

	go deleteFile(filename, 3*time.Second)
}

// Обработчик множественной загрузки
func handleMultipleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t.Execute(w, nil)
		return
	}

	err := r.ParseMultipartForm(16 << 20)
	if err != nil {
		log.Fatal(err)
	}

	form := r.MultipartForm
	files := form.File["file"]
	var uploaded int

	for _, f := range files {
		in, err := f.Open()
		if err != nil {
			continue
		}
		defer in.Close()

		filename := path + f.Filename
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		io.Copy(out, in)

		go deleteFile(filename, 3*time.Second)
		uploaded++
	}

	fmt.Fprintln(w, "Загружено файлов:", uploaded)
}

// Отложенное удаление файла
func deleteFile(filename string, delay time.Duration) {
	time.Sleep(delay)
	os.Remove(filename)
}

// Определение типа файла по заголовку
func typeByHeader(header *multipart.FileHeader) string {
	return header.Header["Content-Type"][0]
}

// Определение типа файла по расширению
func typeByExt(filename string) string {
	ext := filepath.Ext(filename)
	return mime.TypeByExtension(ext)
}

// Определение типа файла по содержимому
func typeByContent(file multipart.File) string {
	buffer := make([]byte, 512)
	file.Read(buffer)
	return http.DetectContentType(buffer)
}

func main() {
	fmt.Println(" \n[ ЗАГРУЗКА ФАЙЛОВ ]\n ")

	// Обработчики
	http.HandleFunc("/", handleUpload)
	http.HandleFunc("/multiple", handleMultipleUpload)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
