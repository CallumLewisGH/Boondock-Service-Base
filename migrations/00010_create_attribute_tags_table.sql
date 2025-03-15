-- +goose Up
-- +goose StatementBegin
-- Create attribute_tags Table
CREATE TABLE IF NOT EXISTS base.attribute_tags (
    "id"             UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"attribute_id"   UUID NOT NULL,
	"value"          VARCHAR(128),

    FOREIGN KEY (attribute_id) REFERENCES base.attributes(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop attribute_tags Table
DROP TABLE IF EXISTS base.attribute_tags
-- +goose StatementEnd
