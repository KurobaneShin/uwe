package types

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
