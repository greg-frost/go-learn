package adapters

import "github.com/julienschmidt/httprouter"

// Интерфейс "обработчик"
type Handler interface {
	Register(router *httprouter.Router)
}
