package banking

import (
	"errors"

	"github.com/shopspring/decimal"
)

func (s *service) Deposit(accountID int64, amount decimal.Decimal) error {

	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("invalid amount")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 🔒 lock account
	acc, err := s.accountRepo.LockAccountForUpdate(tx, accountID)
	if err != nil {
		return err
	}

	// 💰 update balance
	acc.Balance = acc.Balance.Add(amount)

	// 💾 save
	err = s.accountRepo.UpdateAccount(acc)
	if err != nil {
		return err
	}

	// 🧾 log transaction
	_, err = tx.Exec(`
		INSERT INTO transactions
		(from_account_id, to_account_id, amount, currency, type, status)
		VALUES (NULL, ?, ?, 'BDT', 'deposit', 'completed')
	`, accountID, amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}
