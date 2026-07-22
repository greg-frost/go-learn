package apiserver

import "net/http"

// Структура "райтер ответа"
type responseWriter struct {
	http.ResponseWriter
	code int
}

// Запись заголовка
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.code = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
