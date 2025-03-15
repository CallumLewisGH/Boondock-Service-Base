-- +goose Up
-- +goose StatementBegin
-- Create user_campsites Table
CREATE TABLE IF NOT EXISTS base.user_campsites (
    "id"             SERIAL PRIMARY KEY NOT NULL,
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"user_id"		 UUID NOT NULL,
	"coordinates"	 POINT NOT NULL,
	"name"			 VARCHAR(128) NOT NULL,
	"description"    VARCHAR(512) NULL,
	"photo" 	     TEXT NULL,

	FOREIGN KEY (user_id) REFERENCES base.users(id)
);
CREATE UNIQUE INDEX idx_name ON base.user_campsites(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop user_campsites Table
DROP TABLE IF EXISTS base.user_campsites;
-- +goose StatementEnd
