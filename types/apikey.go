package types

import (
	"crypto/rand"
	"encoding/hex"
)

type ApiKey struct {
	ID        int `bun:"id,pk,autoincrement"`
	Key       string
	Name      string
	AccountId int
	Disabled  bool
	Timestamps
}

func NewAPIKey(accountId int, name string) ApiKey {
	return ApiKey{
		AccountId: accountId,
		Name:      name,
		Key:       generateAPIKey(32),
	}
}

func generateAPIKey(length int) string {
	key := make([]byte, length/2) // Since hex encoding uses 2 characters per byte
	_, err := rand.Read(key)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(key)
}
