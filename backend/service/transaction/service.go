package transaction

import (
	"smart-e-banking/backend/domain"
	"smart-e-banking/backend/repository/account"
	"smart-e-banking/backend/repository/transaction"

	"github.com/jmoiron/sqlx"
)

type Service interface {
	CreateTransaction(tx domain.Transaction) (*domain.Transaction, error)
	GetTransactionByID(id int64) (*domain.Transaction, error)
	GetTransactionsByAccountID(accountID int64) ([]domain.Transaction, error)
}

type service struct {
	db              *sqlx.DB
	transactionRepo transaction.Repository // সঠিক ইন্টারফেস টাইপ
	accRepo         account.Repository
}

var _ Service = (*service)(nil)

func NewService(db *sqlx.DB, transRepo transaction.Repository, accRepo account.Repository) Service {
	return &service{
		db:              db,
		transactionRepo: transRepo,
		accRepo:         accRepo,
	}
}
