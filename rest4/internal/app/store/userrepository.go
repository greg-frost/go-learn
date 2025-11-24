package store

import "go-learn/rest4/internal/app/model"

// Структура "хранилище пользователей"
type UserRepository struct {
	store *Store
}

// Создание пользователя
func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES($1, $2) RETURNING id",
		user.Email,
		user.EncryptedPassword,
	).Scan(
		&user.ID,
	); err != nil {
		return nil, err
	}

	return user, nil
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
		return nil, err
	}

	return user, nil
}
