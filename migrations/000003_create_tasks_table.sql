-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks
(
    id       SERIAL PRIMARY KEY NOT NULL,
    name     VARCHAR(255)       NOT NULL,
    priority INT                NOT NULL,
    done     BOOL               NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
