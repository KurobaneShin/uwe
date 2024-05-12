package db

import (
	"context"

	"uwe/types"
)

func (db DB) CreateCustomer(c *types.Customer) error {
	_, err := db.NewInsert().Model(c).Exec(context.Background())
	return err
}

func (db DB) CreateFileUpload(f *types.FileUpload) error {
	_, err := db.NewInsert().Model(f).Exec(context.Background())

	return err
}
