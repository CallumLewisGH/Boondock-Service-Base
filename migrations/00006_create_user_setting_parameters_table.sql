-- +goose Up
-- +goose StatementBegin
-- Create user_settings_parameters Table
CREATE TABLE IF NOT EXISTS base.user_settings_parameters (
    "id"             SERIAL PRIMARY KEY NOT NULL,
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
    "parameter_type"  VARCHAR(128) NOT NULL,
	"user_id"		 UUID NOT NULL,
    "setting_id"     UUID NOT NULL,
    "value"          VARCHAR(128) NULL,

	FOREIGN KEY (user_id) REFERENCES base.users(id),
	FOREIGN KEY (setting_id) REFERENCES base.settings(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop user_settings_parameters Table
DROP TABLE IF EXISTS base.user_settings_parameters;
-- +goose StatementEnd
