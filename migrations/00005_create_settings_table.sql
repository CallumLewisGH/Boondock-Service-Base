-- +goose Up
-- +goose StatementBegin
-- Create settings Table
CREATE TABLE IF NOT EXISTS base.settings (
    "id"             UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"parameter_type"  VARCHAR(128) NOT NULL,
    "name"            VARCHAR(128) NOT NULL,
    "description"     VARCHAR(128) NOT NULL,
    "default_parameter_value" VARCHAR(128) NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop settings Table
DROP TABLE IF EXISTS base.settings;
-- +goose StatementEnd
