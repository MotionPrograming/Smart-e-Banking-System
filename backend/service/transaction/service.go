package transaction

import (
	accountRepo "smart-e-banking/backend/repository/account"
	transactionRepo "smart-e-banking/backend/repository/transaction"

	"github.com/jmoiron/sqlx"
)

type service struct {
	db              *sqlx.DB
	accountRepo     accountRepo.Repository
	transactionRepo transactionRepo.Repository
}

func NewService(db *sqlx.DB, transactionRepo transactionRepo.Repository, accountRepo accountRepo.Repository) *service {
	return &service{
		db:              db,
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
	}
}
