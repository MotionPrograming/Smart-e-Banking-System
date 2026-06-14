package account

import (
	"errors"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

func (s *service) CreateAccount(userID int64, accountType string, initialBalance decimal.Decimal) (*domain.Account, error) {

	// 1. basic validation
	if initialBalance.IsNegative() {
		return nil, errors.New("initial balance cannot be negative")
	}

	// 2. business rule example
	if initialBalance.LessThan(decimal.NewFromInt(100)) {
		return nil, errors.New("minimum balance must be 100 BDT")
	}

	// 3. call repository
	return s.accountRepo.CreateAccount(userID, accountType, initialBalance)
}
