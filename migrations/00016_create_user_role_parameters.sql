-- +goose Up
-- +goose StatementBegin
-- Create user_roles Table
CREATE TABLE IF NOT EXISTS base.user_roles (
    "id"             SERIAL PRIMARY KEY NOT NULL,
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"role_id"        UUID NOT NULL,
    "user_id"        UUID NOT NULL,

    FOREIGN KEY (role_id) REFERENCES base.roles(id),
    FOREIGN KEY (user_id) REFERENCES base.users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop user_roles Table
DROP TABLE IF EXISTS base.user_roles;
-- +goose StatementEnd
