package main

import (
	"bytes"
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
var path = os.Getenv("GOPATH") + "/src/learn/upload/"
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

// Обработчик потоковой загрузки
func handleStreamUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t.Execute(w, nil)
		return
	}

	mr, err := r.MultipartReader()
	if err != nil {
		log.Fatal(err)
	}

	values := make(map[string][]string)
	maxValueBytes := int64(10 << 20)
	var parts int

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		name := part.FormName()
		if name == "" {
			continue
		}
		filename := part.FileName()
		var b bytes.Buffer

		// Текстовое поле
		if filename == "" {
			n, err := io.CopyN(&b, part, maxValueBytes)
			if err != nil && err != io.EOF {
				fmt.Fprint(w, "Ошибка чтения сообщения")
				return
			}
			maxValueBytes -= n
			if maxValueBytes == 0 {
				fmt.Fprint(w, "Сообщение слишком большое")
				return
			}
			values[name] = append(values[name], b.String())
			continue
		}

		// Файловое поле
		filename = path + filename
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		for {
			buffer := make([]byte, 100000)
			m, err := part.Read(buffer)
			if err == io.EOF {
				break
			}
			out.Write(buffer[:m])
			parts++
		}

		go deleteFile(filename, 3*time.Second)
	}

	fmt.Fprintln(w, "Загружено фрагментов:", parts)

	formValue := values["name"][0]
	if formValue != "" {
		fmt.Fprintln(w, "Текстовое поле:", formValue)
	}
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
	http.HandleFunc("/stream", handleStreamUpload)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
