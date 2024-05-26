package types

import (
	"time"

	"github.com/google/uuid"
)

var UUIDZERO = uuid.MustParse("00000000-0000-0000-0000-000000000000")

type Timestamps struct {
	CreatedAt time.Time `json:"created_at,omitempty" bun:",nullzero,notnull,default:current_timestamp`
	UpdatedAt time.Time `json:"updated_at,omitempty" bun:",nullzero,notnull,default:current_timestamp`
	DeletedAt time.Time `json:"deleted_at,omitempty" bun:",nullzero,notnull,default:current_timestamp`
}
