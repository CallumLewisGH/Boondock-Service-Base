-- +goose Up
-- +goose StatementBegin
-- Insert login attempts
INSERT INTO base.login_attempts (user_name, email, login_successfull) VALUES
    ('admin_user', 'admin@example.com', true),
    ('regular_user', 'user@example.com', false);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base.login_attempts;
-- +goose StatementEnd