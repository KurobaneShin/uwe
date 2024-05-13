package types

import "time"

type Timestamps struct {
	CreatedAt time.Time `json:"created_at,omitempty" bun:",nullzero,notnull,default:current_timestamp`
	UpdatedAt time.Time `json:"updated_at,omitempty" bun:",nullzero,notnull,default:current_timestamp`
	DeletedAt time.Time `json:"deleted_at,omitempty" bun:",nullzero,notnull,default:current_timestamp`
}
