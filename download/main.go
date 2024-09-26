package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"go-learn/base"
)

// Путь и таймаут
var path = base.Dir("download")
var timeout = 1 * time.Second

// Скачивание
func download(location string, file *os.File, retries int64) error {
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}
	current := info.Size()
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "bytes="+start+"-")
		// req.Header.Set("Range", fmt.Sprintf("bytes=%d-", current))
	}

	client := &http.Client{Timeout: timeout}
	res, err := client.Do(req)
	if err != nil {
		if hasTimeout(err) && retries > 0 {
			fmt.Println("Еще одна попытка...")
			return download(location, file, retries-1)
		}
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("Не удалось выполнить запрос, статус: %s", res.Status)
	}
	// if res.Header.Get("Accept-Ranges") != "bytes" {
	// 	fmt.Println("Продолжение скачивания не поддерживается!")
	// 	retries = 0
	// }

	_, err = io.Copy(file, res.Body)
	if err != nil {
		if hasTimeout(err) && retries > 0 {
			fmt.Println("Еще одна попытка...")
			return download(location, file, retries-1)
		}
		return err
	}

	return nil
}

// Был ли таймаут
func hasTimeout(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}
	errMsg := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errMsg) {
		return true
	}
	return false
}

func main() {
	fmt.Println(" \n[ СКАЧИВАНИЕ ФАЙЛОВ ]\n ")

	fmt.Println("Идет скачивание...")

	// Создание локального файла
	filename := filepath.Join(path, "file.zip")
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(filename)
	defer file.Close()

	// Скачивание удаленного файла
	location := "https://www.learningcontainer.com/" +
		"wp-content/uploads/2020/05/sample-large-zip-file.zip"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Отчет
	fmt.Println("Скачано байт:", info.Size())
}
