-- +goose Up
-- +goose StatementBegin
-- Insert users
INSERT INTO base.users (id, user_name, email, password_hash) VALUES
    ('f1eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 'admin_user', 'admin@example.com', 'salt1'),
    ('f2eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 'regular_user', 'user@example.com', 'salt2');

-- Assign roles
INSERT INTO base.user_roles (role_id, user_id) VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'f1eebc99-9c0b-4ef8-bb6d-6bb9bd380a18'),
    ('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'f2eebc99-9c0b-4ef8-bb6d-6bb9bd380a19');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base.user_roles;
DELETE FROM base.users;
-- +goose StatementEnd