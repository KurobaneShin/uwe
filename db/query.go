package db

import (
	"context"
	"uwe/types"

	"github.com/google/uuid"
)

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
