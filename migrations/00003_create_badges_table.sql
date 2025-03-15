-- +goose Up
-- +goose StatementBegin
-- Create badges Table
CREATE TABLE IF NOT EXISTS base.badges (
    "id"             UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
    "created_utc"    TIMESTAMP        NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP        NULL,

    "name"           VARCHAR(128)     NOT NULL,
    "description"    VARCHAR(128)     NOT NULL,
    "icon_art"       TEXT             NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop badges Table
DROP TABLE IF EXISTS base.badges;

-- +goose StatementEnd
