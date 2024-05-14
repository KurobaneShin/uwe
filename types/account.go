package types

type Account struct {
	ID        int `bun:"id,pk,autoincrement"`
	AccountID int
	Timestamps
}
