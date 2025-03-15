-- +goose Up
-- +goose StatementBegin
-- Create attribute_parameters Table
CREATE TABLE IF NOT EXISTS base.attribute_parameters (
    "id"               SERIAL PRIMARY KEY NOT NULL,
    "created_utc"      TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"      TIMESTAMP NULL,
	
	"attribute_id"	   UUID NOT NULL,
    "review_id"        INT NOT NULL,
    "score_value"      INT NULL, 
    "boolean_value"    BOOLEAN NULL,

	FOREIGN KEY (attribute_id) REFERENCES base.attributes(id) ,
    FOREIGN KEY (review_id) REFERENCES base.campsite_reviews(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop attribute_parameters Table
DROP TABLE IF EXISTS base.attribute_parameters
-- +goose StatementEnd
