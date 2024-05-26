package types

import (
	"github.com/google/uuid"
)

type Customer struct {
	Id        uuid.UUID `json:"id,omitempty" bun:"id,pk"`
	Name      string    `json:"name,omitempty"`
	AccountId int       `json:"-"`
	Timestamps
}
