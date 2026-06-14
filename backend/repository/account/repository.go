package account

import (
	"smart-e-banking/backend/domain"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Repository interface {
	CreateAccount(userID int64, accountType string, initialBalance decimal.Decimal) (*domain.Account, error)
	DeleteAccount(id int64) error
	GetAccountByID(accountID int64) (*domain.Account, error)
	GetAccountsByUserID(userID int64) ([]domain.Account, error)

	UpdateAccount(account *domain.Account) error
	LockAccountForUpdate(tx *sqlx.Tx, id int64) (*domain.Account, error)
}
type accountRepository struct {
	db *sqlx.DB
}

var _ Repository = (*accountRepository)(nil)

func NewAccountRepository(db *sqlx.DB) Repository {
	return &accountRepository{db: db}
}
