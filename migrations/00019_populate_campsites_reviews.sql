-- +goose Up
-- +goose StatementBegin
-- Insert campsites
INSERT INTO base.user_campsites (user_id, coordinates, name, description) VALUES
    ('f1eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', POINT(37.7749,-122.4194), 'Golden Gate Camp', 'Scenic spot'),
    ('f2eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', POINT(34.0522,-118.2437), 'LA Camp', 'City view');

-- Insert reviews
INSERT INTO base.campsite_reviews (user_id, campsite_id, headline, description) VALUES
    ('f1eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 1, 'Great!', 'Loved the views'),
    ('f2eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 2, 'Peaceful', 'Quiet retreat');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base.campsite_reviews;
DELETE FROM base.user_campsites;
-- +goose StatementEnd