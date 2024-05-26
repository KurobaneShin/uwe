
-- +goose Up
-- +goose StatementBegin
CREATE TABLE api_keys (
id serial not null primary key,
name text NOT NULL,
key text not null,
active boolean default true,
account_id int references accounts(id),
created_at timestamp default 'now()',
updated_at timestamp,
deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE api_keys;
-- +goose StatementEnd
