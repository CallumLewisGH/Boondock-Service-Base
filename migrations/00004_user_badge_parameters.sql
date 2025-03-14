-- +goose Up
-- +goose StatementBegin
-- Create user_bage_parameters Table
CREATE TABLE IF NOT EXISTS base.user_badge_parameters (
    "id"             SERIAL PRIMARY KEY NOT NULL,
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"user_id"		  UUID NOT NULL,
    "badge_id"        UUID NOT NULL,
    "description"     VARCHAR(128) NOT NULL,

	FOREIGN KEY (user_id) REFERENCES base.user_credentials(id),
	FOREIGN KEY (badge_id) REFERENCES base.user_badges(id) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop user_bage_parameters Table
DROP TABLE IF EXISTS base.user_badge_parameters;
-- +goose StatementEnd
