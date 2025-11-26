package sqlstore_test

import (
	"testing"

	"go-learn/rest4/internal/app/model"
	"go-learn/rest4/internal/app/store"
	"go-learn/rest4/internal/app/store/sqlstore"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.com"

	// Пользователь не найден
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// Пользователь найден
	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
