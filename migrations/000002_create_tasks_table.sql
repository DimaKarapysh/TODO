-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS winners
(
    id    SERIAL PRIMARY KEY NOT NULL,
    name  VARCHAR(255)       NOT NULL,
    points INT                NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS winners;
-- +goose StatementEnd
