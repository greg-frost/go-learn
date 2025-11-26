package sqlstore

import (
	"database/sql"

	"go-learn/rest4/internal/app/model"
	"go-learn/rest4/internal/app/store"
)

// Структура "хранилище пользователей"
type UserRepository struct {
	store *Store
}

// Создание пользователя
func (r *UserRepository) Create(user *model.User) error {
	// Валидация
	if err := user.Validate(); err != nil {
		return err
	}

	// Подготовка
	if err := user.BeforeCreate(); err != nil {
		return err
	}

	// Сохранение
	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES($1, $2) RETURNING id",
		user.Email,
		user.EncryptedPassword,
	).Scan(
		&user.ID,
	)
}

// Поиск пользователя по Email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return user, nil
}
