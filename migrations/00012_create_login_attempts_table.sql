-- +goose Up
-- +goose StatementBegin
-- Create login_attempts Table
CREATE TABLE IF NOT EXISTS base.login_attempts (
	"id"                SERIAL PRIMARY KEY NOT NULL,
	"created_utc"       TIMESTAMP        NOT NULL DEFAULT NOW(),
	"deleted_utc"       TIMESTAMP        NULL,

	"user_name"         VARCHAR(128)    NULL,
	"email"             VARCHAR(128)    NULL,
	"login_successfull" BOOLEAN NOT     NULL
);
CREATE INDEX idx_created_utc ON base.login_attempts USING BRIN (created_utc);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop login_attempts Table
DROP TABLE IF EXISTS base.login_attempts;
-- +goose StatementEnd
