package handler

import "github.com/gorilla/mux"

// Интерфейс "обработчик"
type Handler interface {
	Register(router *mux.Router)
}
