package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID            int64           `db:"id" json:"id"`
	FromAccountID int64           `db:"from_account_id" json:"from_account_id"`
	ToAccountID   int64           `db:"to_account_id" json:"to_account_id"`
	Amount        decimal.Decimal `db:"amount" json:"amount"`
	Currency      string          `db:"currency" json:"currency"`
	Type          string          `db:"type" json:"type"`
	Status        string          `db:"status" json:"status"`
	Reference     string          `db:"reference" json:"reference"`
	CreatedAt     time.Time       `db:"created_at" json:"created_at"`
}
