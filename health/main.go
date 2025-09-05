package main

import (
	"fmt"
	"log"
	"net/http"
)

// Проверка доступности
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	fmt.Println(" \n[ ПРОВЕРКА РАБОТОСПОСОБНОСТИ ]\n ")

	// Обработчики
	http.HandleFunc("/health", healthHandler)

	// Запуск сервера
	fmt.Println("Ожидаю обновлений...")
	fmt.Println("(на http://localhost:8080)")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
