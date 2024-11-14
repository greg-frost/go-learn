-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD PRIMARY KEY (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP CONSTRAINT IF EXISTS users_pkey;
-- +goose StatementEnd