-- +goose Up
-- +goose StatementBegin
-- Create campsite_reviews Table
CREATE TABLE IF NOT EXISTS base.campsite_reviews (
    "id"             SERIAL PRIMARY KEY NOT NULL,
    "created_utc"    TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_utc"    TIMESTAMP NULL,
	
	"user_id"		 UUID NOT NULL,
	"campsite_id"	 INT NOT NULL,
	"headline"		 VARCHAR(128),
	"description"    VARCHAR(512),

	FOREIGN KEY (user_id) REFERENCES base.users(id),
	FOREIGN KEY (campsite_id) REFERENCES base.user_campsites(id) 
);
CREATE UNIQUE INDEX idx_headline ON base.campsite_reviews(headline);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop campsite_reviews Table
DROP TABLE IF EXISTS base.campsite_reviews
-- +goose StatementEnd
