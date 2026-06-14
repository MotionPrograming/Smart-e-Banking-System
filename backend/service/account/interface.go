package account

import (
	"smart-e-banking/backend/domain"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type AccountRepo interface {
	CreateAccount(userID int64, accountType string, initialBalance decimal.Decimal) (*domain.Account, error)

	GetAccountByID(id int64) (*domain.Account, error)
	GetAccountsByUserID(userID int64) ([]domain.Account, error)

	UpdateAccount(acc *domain.Account) error
	DeleteAccount(id int64) error
	LockAccountForUpdate(tx *sqlx.Tx, accountID int64) (*domain.Account, error)
}

var _ Service = (*service)(nil)
