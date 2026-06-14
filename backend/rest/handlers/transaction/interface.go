package transaction

import (
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

type Service interface {
	CreateTransaction(tx domain.Transaction) (*domain.Transaction, error)
	GetTransactionByID(id int64) (*domain.Transaction, error)
	GetTransactionsByAccountID(accountID int64) ([]domain.Transaction, error)

	Deposit(accountID int64, amount decimal.Decimal) error
	Withdraw(accountID int64, amount decimal.Decimal) error
	Transfer(fromAccountID, toAccountID int64, amount decimal.Decimal) error
}
