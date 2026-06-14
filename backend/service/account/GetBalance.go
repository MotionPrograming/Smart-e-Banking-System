package account

import (
	"errors"

	"github.com/shopspring/decimal"
)

func (h *service) GetBalance(accountID int64) (decimal.Decimal, error) {

	account, err := h.GetAccountByID(accountID)
	if err != nil {
		return decimal.Zero, err
	}

	// optional safety check (if not already in GetAccountByID)
	if account.Status != "active" {
		return decimal.Zero, errors.New("account is not active")
	}

	return account.Balance, nil
}
