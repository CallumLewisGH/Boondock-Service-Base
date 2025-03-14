-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS base;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS base;
-- +goose StatementEnd
