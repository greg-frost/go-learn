package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Проверка жизнеспособности (доступности)
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Поверхностная проверка работоспособности
func shallowHealthHandler(w http.ResponseWriter, r *http.Request) {
	// Создание временного файла
	tmp, err := ioutil.TempFile(os.TempDir(), "shallow-")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer os.Remove(tmp.Name())

	// Проверка записи в файл
	text := []byte("Проверка работоспособности")
	if _, err := tmp.Write(text); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	// Проверка закрытия файла
	if err := tmp.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	fmt.Println(" \n[ ПРОВЕРКА РАБОТОСПОСОБНОСТИ ]\n ")

	// Обработчики
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/shallow", shallowHealthHandler)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
