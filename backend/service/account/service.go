package account

import (
	"smart-e-banking/backend/domain"
	accountRepo "smart-e-banking/backend/repository/account"
	transactionRepo "smart-e-banking/backend/repository/transaction"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Service interface {
	CreateAccount(userID int64, accountType string, initialBalance decimal.Decimal) (*domain.Account, error)
	GetAccountByID(id int64) (*domain.Account, error)
	GetAccountsByUserID(userID int64) ([]domain.Account, error)
	GetBalance(accountID int64) (decimal.Decimal, error)
	UpdateAccount(acc *domain.Account) error
	DeleteAccount(id int64) error

	Deposit(accountID int64, amount decimal.Decimal) error
	Withdraw(accountID int64, amount decimal.Decimal, reference string) error
	Transfer(fromAccountID, toAccountID int64, amount decimal.Decimal) error
}

type service struct {
	db              *sqlx.DB
	accountRepo     accountRepo.Repository
	transactionRepo transactionRepo.Repository
}

var _ Service = (*service)(nil)

func NewService(
	db *sqlx.DB,
	accountRepo accountRepo.Repository,
	transactionRepo transactionRepo.Repository,
) Service {
	return &service{
		db:              db,
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}
