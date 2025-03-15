-- +goose Up
-- +goose StatementBegin
-- Insert attributes
INSERT INTO base.attributes (id, name, description, parameter_type, has_tags) VALUES
    ('a1eebc99-9c0b-4ef8-bb6d-6bb9bd380a24', 'Scenic', 'Beautiful views', 'boolean', false),
    ('a2eebc99-9c0b-4ef8-bb6d-6bb9bd380a25', 'Access', 'Accessibility', 'string', true);

-- Insert tags
INSERT INTO base.attribute_tags (attribute_id, value) VALUES
    ('a2eebc99-9c0b-4ef8-bb6d-6bb9bd380a25', 'Wheelchair'),
    ('a2eebc99-9c0b-4ef8-bb6d-6bb9bd380a25', 'Stroller');

-- Insert parameters
INSERT INTO base.attribute_parameters (attribute_id, review_id, boolean_value) VALUES
    ('a1eebc99-9c0b-4ef8-bb6d-6bb9bd380a24', 1, true);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base.attribute_parameters;
DELETE FROM base.attribute_tags;
DELETE FROM base.attributes;
-- +goose StatementEnd