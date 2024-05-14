package db

import (
	"context"
	"database/sql"
	"uwe/types"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

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
