package transaction

import (
	"smart-e-banking/backend/domain"
)

type Service interface {
	CreateTransaction(transaction domain.Transaction) (*domain.Transaction, error)
	GetTransactionByID(id int64) (*domain.Transaction, error)
	GetTransactionsByAccountID(accountID int64) ([]domain.Transaction, error)
}

type TransactionRepo interface {
	CreateTransaction(transaction domain.Transaction) (*domain.Transaction, error)
	GetTransactionByID(id int64) (*domain.Transaction, error)
	GetTransactionsByAccountID(accountID int64) ([]domain.Transaction, error)
}
