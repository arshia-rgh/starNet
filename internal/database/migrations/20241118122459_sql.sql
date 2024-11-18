-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN role VARCHAR(255) DEFAULT 'normal' NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN role;
-- +goose StatementEnd
