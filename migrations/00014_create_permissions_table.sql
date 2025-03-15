-- +goose Up
-- +goose StatementBegin
-- Create permissions Table
CREATE TABLE IF NOT EXISTS base.permissions (
    "id"             UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"name"		     VARCHAR(128) NOT NULL,
    "description"    VARCHAR(128) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop permissions Table
DROP TABLE IF EXISTS base.permissions;
-- +goose StatementEnd
