package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"uwe/types"
)

func (db DB) GetAccountByAPIKey(apiKey string) (types.Account, error) {
	var account types.Account

	err := db.NewSelect().
		Model(&account).
		Join("JOIN api_keys as keys on keys.account_id = account.id").
		Where("keys.active", true).
		Where("keys.key = ?", apiKey).
		Scan(context.Background())
	return account, err
}

func (db *DB) CreateAccount(account *types.Account) error {
	return db.RunInTx(context.Background(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		_, err := db.NewInsert().
			Model(account).
			Exec(context.Background())
		if err != nil {
			return err
		}

		key := types.NewAPIKey(account.ID, "default")
		_, err = db.NewInsert().
			Model(key).
			Exec(context.Background())
		if err != nil {
			return err
		}
		return nil
	})
}

func (db DB) CreateCustomer(c *types.Customer) error {
	_, err := db.NewInsert().Model(c).Exec(context.Background())
	return err
}

func (db DB) CreateFileUpload(f *types.FileUpload) error {
	_, err := db.NewInsert().Model(f).Exec(context.Background())

	return err
}

func (db DB) GetFileUploadById(id uuid.UUID) (types.FileUpload, error) {
	var fileUpload types.FileUpload
	err := db.NewSelect().Model(&fileUpload).Where("id = ?", id).Scan(context.Background())
	return fileUpload, err
}
