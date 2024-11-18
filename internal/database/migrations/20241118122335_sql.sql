-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id       serial PRIMARY KEY NOT NULL,
    username VARCHAR(255)       NOT NULL UNIQUE,
    password VARCHAR(255)       NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
