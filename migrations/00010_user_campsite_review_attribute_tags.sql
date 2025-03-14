-- +goose Up
-- +goose StatementBegin
-- Create user_campsite_review_attribute_tags Table
CREATE TABLE IF NOT EXISTS base.user_campsite_review_attribute_tags (
    "id"             UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"attribute_id"   UUID NOT NULL,
	"value"          VARCHAR(128),

    FOREIGN KEY (attribute_id) REFERENCES base.user_campsite_review_attributes(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop user_campsite_review_attribute_tags Table
DROP TABLE IF EXISTS base.user_campsite_review_attribute_tags
-- +goose StatementEnd
