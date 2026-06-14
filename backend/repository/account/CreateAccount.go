package account

import (
	"smart-e-banking/backend/domain"
	"smart-e-banking/backend/util"

	"github.com/shopspring/decimal"
)

func (r *accountRepository) CreateAccount(
	userID int64,
	accountType string,
	initialBalance decimal.Decimal,
) (*domain.Account, error) {
	accountNumber := util.GenerateAccountNumber()

	result, err := r.db.Exec(
		`INSERT INTO accounts (user_id, account_number, balance, currency, account_type, status)
		 VALUES (?, ?, ?, 'BDT', ?, 'active')`,
		userID, accountNumber, initialBalance, accountType,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var acc domain.Account
	if err := r.db.Get(&acc, `SELECT * FROM accounts WHERE id = ?`, id); err != nil {
		return nil, err
	}
	return &acc, nil
}
