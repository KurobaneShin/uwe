package types

import (
	"github.com/google/uuid"
)

type Customer struct {
	Id        uuid.UUID
	Name      string
	AccountId int
	Timestamps
}
