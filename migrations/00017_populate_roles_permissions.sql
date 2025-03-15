-- +goose Up
-- +goose StatementBegin
-- Insert roles
INSERT INTO base.roles (id, name, description) VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'admin', 'Administrator role'),
    ('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'user', 'Standard user role');

-- Insert permissions
INSERT INTO base.permissions (id, name, description) VALUES
    ('c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'create_user', 'Create users'),
    ('d3eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'delete_user', 'Delete users'),
    ('e4eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'edit_user', 'Edit users'),
    ('f5eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 'view_user', 'View users'),
    ('f6eebc99-9c0b-4ef8-bb6d-6bb9bd380a17', 'manage_roles', 'Manage roles');

-- Assign permissions
INSERT INTO base.role_permissions (role_id, permission_id) VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a13'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd3eebc99-9c0b-4ef8-bb6d-6bb9bd380a14'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'e4eebc99-9c0b-4ef8-bb6d-6bb9bd380a15'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'f5eebc99-9c0b-4ef8-bb6d-6bb9bd380a16'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'f6eebc99-9c0b-4ef8-bb6d-6bb9bd380a17'),
    ('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'f5eebc99-9c0b-4ef8-bb6d-6bb9bd380a16');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base.role_permissions;
DELETE FROM base.permissions;
DELETE FROM base.roles;
-- +goose StatementEnd