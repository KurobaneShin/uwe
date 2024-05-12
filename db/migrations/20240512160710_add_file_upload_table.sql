-- +goose Up
-- +goose StatementBegin
CREATE TABLE file_uploads(
  id UUID NOT NULL PRIMARY KEY,
  customer_id uuid references customers,
  "type" int NOT NULL,
  mapping jsonb default '{}',
  created_at timestamp default 'now',
  updated_at timestamp,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE file_uploads;
-- +goose StatementEnd
