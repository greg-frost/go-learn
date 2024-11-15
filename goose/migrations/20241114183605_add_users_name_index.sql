-- +goose Up
CREATE INDEX idx_users_name ON users(name);

-- +goose Down
DROP INDEX IF EXISTS idx_users_name;