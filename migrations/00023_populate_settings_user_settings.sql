-- +goose Up
-- +goose StatementBegin
-- Insert settings
INSERT INTO base.settings (id, parameter_type, name, description, default_parameter_value) VALUES
    ('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'boolean', 'Dark Mode', 'Dark mode', 'true'),
    ('c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a23', 'string', 'Language', 'UI language', 'en');

-- Assign settings
INSERT INTO base.user_settings_parameters (parameter_type, user_id, setting_id, value) VALUES
    ('boolean', 'f1eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 'c1eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'true'),
    ('string', 'f2eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 'c2eebc99-9c0b-4ef8-bb6d-6bb9bd380a23', 'fr');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base.user_settings_parameters;
DELETE FROM base.settings;
-- +goose StatementEnd