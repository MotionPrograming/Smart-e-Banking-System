package account

import (
	transRepoPkg "smart-e-banking/backend/repository/transaction"

	"github.com/jmoiron/sqlx"
)

type service struct {
	db              *sqlx.DB
	accountRepo     AccountRepo
	transactionRepo transRepoPkg.Repository
}

func NewService(
	db *sqlx.DB,
	accountRepo AccountRepo,
	transactionRepo transRepoPkg.Repository,

) Service {
	return &service{
		db:              db,
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}
