-- +goose Up
-- +goose StatementBegin
ALTER TABLE videos
    ADD CONSTRAINT videos_title_unique UNIQUE (title),
    ADD CONSTRAINT videos_file_path_unique UNIQUE (file_path);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE videos
    DROP CONSTRAINT videos_title_unique,
    DROP CONSTRAINT videos_file_path_unique;
-- +goose StatementEnd