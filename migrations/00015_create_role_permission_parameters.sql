-- +goose Up
-- +goose StatementBegin
-- Create role_permissions Table
CREATE TABLE IF NOT EXISTS base.role_permissions (
    "id"             SERIAL PRIMARY KEY NOT NULL,
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"role_id"        UUID NOT NULL,
    "permission_id"  UUID NOT NULL,

    FOREIGN KEY (role_id) REFERENCES base.roles(id),
    FOREIGN KEY (permission_id) REFERENCES base.permissions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop role_permissions Table
DROP TABLE IF EXISTS base.role_permissions;
-- +goose StatementEnd
