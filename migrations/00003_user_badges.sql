-- +goose Up
-- +goose StatementBegin
-- Create user_badges Table
CREATE TABLE IF NOT EXISTS base.user_badges (
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
-- Drop user_badges Table
DROP TABLE IF EXISTS base.user_badges;

-- +goose StatementEnd
