-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_matrix (
    id SERIAL PRIMARY KEY,
    role TEXT NOT NULL,
    endpoint TEXT NOT NULL
);

INSERT INTO access_matrix (role, endpoint) VALUES
    ('ADMIN', '/user_v1.UserV1/Create'),
    ('ADMIN', '/user_v1.UserV1/Get'),
    ('ADMIN', '/user_v1.UserV1/Update'),
    ('ADMIN', '/user_v1.UserV1/Delete'),
    ('ADMIN', '/auth_v1.AuthV1/GetAccessToken'),
    ('ADMIN', '/auth_v1.AuthV1/GetRefreshToken'),
    ('ADMIN', '/auth_v1.AuthV1/Login'),
    ('ADMIN', '/access_v1.AccessV1/Check'),
    ('USER', '/user_v1.UserV1/Create'),
    ('USER', '/auth_v1.AuthV1/GetAccessToken'),
    ('USER', '/auth_v1.AuthV1/GetRefreshToken'),
    ('USER', '/access_v1.AccessV1/Check');
    ('UNKNOWN', '/user_v1.UserV1/Create'),
    ('USER', '/chat_v1.ChatV1/Create');
    ('USER', '/chat_v1.ChatV1/SendMessage');
    ('USER', '/chat_v1.ChatV1/Delete');
    ('ADMIN', '/chat_v1.ChatV1/Create');
    ('ADMIN', '/chat_v1.ChatV1/SendMessage');
    ('ADMIN', '/chat_v1.ChatV1/Delete');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE access_matrix;
-- +goose StatementEnd