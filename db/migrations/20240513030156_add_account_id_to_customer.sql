-- +goose Up
-- +goose StatementBegin
ALTER TABLE customers add account_id serial;
ALTER TABLE customers
ADD CONSTRAINT fk_account
FOREIGN KEY(account_id)
REFERENCES accounts(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE customers DROP CONSTRAINT fk_account;
ALTER TABLE customers DROP COLUMN account_id;
-- +goose StatementEnd
