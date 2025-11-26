package teststore_test

import (
	"testing"

	"go-learn/rest4/internal/app/model"
	"go-learn/rest4/internal/app/store/teststore"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
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
