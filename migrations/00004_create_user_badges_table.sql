-- +goose Up
-- +goose StatementBegin
-- Create user_bages Table
CREATE TABLE IF NOT EXISTS base.user_badges (
    "id"             SERIAL PRIMARY KEY NOT NULL,
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"user_id"		  UUID NOT NULL,
    "badge_id"        UUID NOT NULL,

	FOREIGN KEY (user_id) REFERENCES base.users(id),
	FOREIGN KEY (badge_id) REFERENCES base.badges(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop user_bages Table
DROP TABLE IF EXISTS base.user_badges;
-- +goose StatementEnd
