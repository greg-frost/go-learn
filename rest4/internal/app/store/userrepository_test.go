package store_test

import (
	"testing"

	"go-learn/rest4/internal/app/model"
	"go-learn/rest4/internal/app/store"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.com"

	// Пользователь не найден
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	// Пользователь найден
	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
