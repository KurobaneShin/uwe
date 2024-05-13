-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts (
  id serial NOT NULL primary key,
  created_at timestamp default 'now()',
  updated_at timestamp,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE accounts;
-- +goose StatementEnd
