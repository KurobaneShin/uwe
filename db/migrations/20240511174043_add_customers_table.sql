-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers (
id uuid NOT NULL PRIMARY KEY,
created_at timestamp,
updated_at timestamp,
deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customers;
-- +goose StatementEnd
