-- +goose Up
-- +goose StatementBegin
-- Insert badges
INSERT INTO base.badges (id, name, description, icon_art) VALUES
    ('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a20', 'Early Adopter', 'Early user', 'icon1'),
    ('b2eebc99-9c0b-4ef8-bb6d-6bb9bd380a21', 'Contributor', 'Active contributor', 'icon2');

-- Assign badges
INSERT INTO base.user_badges (user_id, badge_id) VALUES
    ('f1eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 'b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a20'),
    ('f2eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 'b2eebc99-9c0b-4ef8-bb6d-6bb9bd380a21');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base.user_badges;
DELETE FROM base.badges;
-- +goose StatementEnd