-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS journal (
    id SERIAL PRIMARY KEY,
    action TEXT NOT NULL,
    created_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE journal;
-- +goose StatementEnd