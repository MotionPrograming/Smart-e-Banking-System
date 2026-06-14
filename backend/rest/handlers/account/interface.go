package account

import (
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

type Service interface {
	CreateAccount(userID int64, accountType string, initialBalance decimal.Decimal) (*domain.Account, error)

	DeleteAccount(id int64) error

	GetAccountByID(accountID int64) (*domain.Account, error)

	GetAccountsByUserID(userID int64) ([]domain.Account, error)

	UpdateAccount(acc *domain.Account) error

	Deposit(accountID int64, amount decimal.Decimal) error
	Withdraw(accountID int64, amount decimal.Decimal) error
	Transfer(fromID, toID int64, amount decimal.Decimal) error
}
