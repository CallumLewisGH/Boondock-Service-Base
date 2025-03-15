-- +goose Up
-- +goose StatementBegin
-- Create users Table
CREATE TABLE IF NOT EXISTS base.users (
	"id"             UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
	"created_utc"    TIMESTAMP        NOT NULL DEFAULT NOW(),
	"deleted_utc"    TIMESTAMP        NULL,

	"user_name"       VARCHAR(128)    NOT NULL,
	"email"           VARCHAR(128)    NOT NULL,
	"hash_salt"       VARCHAR(128)    NOT NULL,
	"password_hash"   VARCHAR(128)    NOT NULL,
	"profile_picture" TEXT            NULL
);
CREATE UNIQUE INDEX idx_users_email ON base.users(email);
CREATE UNIQUE INDEX idx_users_user_name ON base.users(user_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop users Table
DROP TABLE IF EXISTS base.users;
-- +goose StatementEnd
