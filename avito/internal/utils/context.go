package util

import (
	"context"

	"go-learn/avito/internal/model"
)

type roleCtx string

var roleKey roleCtx = "role"

func PutRoleIntoContext(ctx context.Context, role model.Role) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

func GetRoleFromContext(ctx context.Context) model.Role {
	value, _ := ctx.Value(roleKey).(model.Role)
	return value
}
