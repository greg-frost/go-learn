-- +goose Up
ALTER TABLE users
    ADD PRIMARY KEY (id);

-- +goose Down
ALTER TABLE users
    DROP CONSTRAINT IF EXISTS users_pkey;