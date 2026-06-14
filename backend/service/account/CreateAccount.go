package account

import (
	"errors"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

func (s *service) CreateAccount(userID int64, accountType string, initialBalance decimal.Decimal) (*domain.Account, error) {
	if initialBalance.IsNegative() {
		return nil, errors.New("initial balance cannot be negative")
	}
	if initialBalance.LessThan(decimal.NewFromInt(100)) {
		return nil, errors.New("minimum initial balance is 100 BDT")
	}
	return s.accountRepo.CreateAccount(userID, accountType, initialBalance)
}
