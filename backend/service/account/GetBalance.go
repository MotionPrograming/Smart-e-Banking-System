package account

import (
	"github.com/shopspring/decimal"
)

func (s *service) GetBalance(accountID int64) (decimal.Decimal, error) {
	acc, err := s.GetAccountByID(accountID)
	if err != nil {
		return decimal.Zero, err
	}
	return acc.Balance, nil
}
