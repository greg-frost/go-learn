package frontend

import (
	"fmt"
)

// Конструктор фронтэнда
func NewFrontEnd(frontendType string) (FrontEnd, error) {
	switch frontendType {
	case "rest":
		return newRestFrontEnd(), nil
	case "grpc":
		return newGrpcFrontEnd(), nil
	default:
		return nil, fmt.Errorf("нет фронтэнда %s", frontendType)
	}
}
