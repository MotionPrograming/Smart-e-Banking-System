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

	query := `
		INSERT INTO accounts
		(user_id, account_number, balance, currency, account_type, status)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
		userID,
		accountNumber,
		initialBalance,
		"BDT",
		accountType,
		"active",
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Return full object
	var acc domain.Account
	err = r.db.Get(
		&acc,
		`SELECT * FROM accounts WHERE id = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}
