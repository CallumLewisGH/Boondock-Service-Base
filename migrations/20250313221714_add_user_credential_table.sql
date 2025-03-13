-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_credentials (
	"id"             UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
	"created_utc"    TIMESTAMP        NOT NULL DEFAULT NOW(),
	"deleted_utc"    TIMESTAMP        NULL,

	"user_name"       VARCHAR(128)    NOT NULL,
	"email"           VARCHAR(128)    NOT NULL,
	"hash_salt"       VARCHAR(128)    NOT NULL,
	"password_hash"   VARCHAR(128)    NOT NULL,
	"profile_picture" TEXT            NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_credentials;
-- +goose StatementEnd
