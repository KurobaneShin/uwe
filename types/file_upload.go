package types

import (
	"time"

	"github.com/google/uuid"
)

type FileType int

const (
	FileTypeSubscription FileType = iota
)

type FileUpload struct {
	ID         uuid.UUID      `json:"id,omitempty"`
	CustomerId uuid.UUID      `json:"customer_id,omitempty"`
	Type       FileType       `json:"type,omitempty"`
	Mapping    map[string]int `json:"mapping,omitempty" bun:"default:{}"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
}
