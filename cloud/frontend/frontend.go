package frontend

import "go-learn/cloud/core"

// Интерфейс "фронтэнд"
type FrontEnd interface {
	Start(kvs *core.KeyValueStore) error
}
