package banking

import (
	"smart-e-banking/backend/repository/account"
	"smart-e-banking/backend/repository/transaction"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Service interface {
	Deposit(accountID int64, amount decimal.Decimal) error
	Withdraw(accountID int64, amount decimal.Decimal) error
	Transfer(fromAccountID, toAccountID int64, amount decimal.Decimal) error
}
type service struct {
	db              *sqlx.DB
	accountRepo     account.Repository
	transactionRepo transaction.Repository
}

func NewService(
	db *sqlx.DB,
	accountRepo account.Repository,
	transactionRepo transaction.Repository,
) Service {
	return &service{
		db:              db,
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}
