package transaction

import (
	"smart-e-banking/backend/domain"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateTransaction(r domain.Transaction) (*domain.Transaction, error)
	GetTransactionByID(transactionID int64) (*domain.Transaction, error)
	GetTransactionsByAccountID(accountID int64) ([]domain.Transaction, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

var _ Repository = (*transactionRepository)(nil)

func NewTransactionRepository(db *sqlx.DB) Repository {
	return &transactionRepository{
		db: db,
	}
}
