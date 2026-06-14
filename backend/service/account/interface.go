package account

import (
	"smart-e-banking/backend/domain"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

// Service layer
type Service interface {
	CreateAccount(userID int64, accountType string, initialBalance decimal.Decimal) (*domain.Account, error)

	GetAccountByID(id int64) (*domain.Account, error)
	GetAccountsByUserID(userID int64) ([]domain.Account, error)

	GetBalance(accountID int64) (decimal.Decimal, error)
	DeleteAccount(id int64) error

	Deposit(accountID int64, amount decimal.Decimal) error
	Withdraw(accountID int64, amount decimal.Decimal, reference string) error
	Transfer(fromAccountID, toAccountID int64, amount decimal.Decimal) error
}
type AccountRepo interface {
	CreateAccount(userID int64, accountType string, initialBalance decimal.Decimal) (*domain.Account, error)

	GetAccountByID(id int64) (*domain.Account, error)
	GetAccountsByUserID(userID int64) ([]domain.Account, error)

	UpdateAccount(acc *domain.Account) error
	DeleteAccount(id int64) error
	LockAccountForUpdate(tx *sqlx.Tx, accountID int64) (*domain.Account, error)
}

var _ Service = (*service)(nil)
