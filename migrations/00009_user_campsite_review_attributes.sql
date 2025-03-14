-- +goose Up
-- +goose StatementBegin
-- Create user_campsite_review_attributes Table
CREATE TABLE IF NOT EXISTS base.user_campsite_review_attributes (
    "id"             UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"name"		     VARCHAR(128) NOT NULL,
	"description"	 VARCHAR(128) NOT NULL,
    "parameter_type" VARCHAR(128) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop user_campsite_review_attributes Table
DROP TABLE IF EXISTS base.user_campsite_review_attributes
-- +goose StatementEnd
