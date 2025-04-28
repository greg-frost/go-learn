package util

import (
	"context"

	"go-learn/avito/internal/model"
)

// Тип ключа роли
type roleField string

// Значение ключа роли
var roleKey roleField = "role"

// Сохранение роли в контексте
func PutRoleIntoContext(ctx context.Context, role model.Role) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

// Получение роли из контекста
func GetRoleFromContext(ctx context.Context) model.Role {
	value, _ := ctx.Value(roleKey).(model.Role)
	return value
}
