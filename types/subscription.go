package types

import "time"

type Currency int

const (
	EUR Currency = iota
	USD
	GBP
	BRL
)

func (c Currency) String() string {
	switch c {
	case EUR:
		return "EUR"
	case USD:
		return "USD"
	case GBP:
		return "GBP"
	case BRL:
		return "BRL"
	default:
		return "INVALID"
	}
}

type Subscription struct {
	Amount     int       `json:"amount,omitempty"`
	Currency   int       `json:"currency,omitempty"`
	Period     int       `json:"period,omitempty"`
	VAT        int       `json:"vat,omitempty"`
	StartedAt  time.Time `json:"started_at,omitempty"`
	CanceledAt time.Time `json:"canceled_at,omitempty"`
	ExternalId string    `json:"external_id,omitempty"`
}
