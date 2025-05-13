package frontend

import (
	"fmt"
)

// Конструктор фронтэнда
func NewFrontEnd(frontendType string) (FrontEnd, error) {
	switch frontendType {
	case "rest":
		return NewRestFrontEnd(), nil
	case "grpc":
		return NewGrpcFrontEnd(), nil
	default:
		return nil, fmt.Errorf("нет фронтэнда %s", frontendType)
	}
}
