-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID,
    name VARCHAR(255),
    description TEXT
);

-- +goose Down
DROP TABLE IF EXISTS users;