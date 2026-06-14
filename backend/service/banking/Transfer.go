package banking

import (
	"errors"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

func (s *service) Transfer(fromID, toID int64, amount decimal.Decimal) error {

	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("invalid amount")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 🔒 deadlock-safe lock order
	var fromAcc, toAcc *domain.Account

	if fromID < toID {
		fromAcc, err = s.accountRepo.LockAccountForUpdate(tx, fromID)
		if err != nil {
			return err
		}

		toAcc, err = s.accountRepo.LockAccountForUpdate(tx, toID)
		if err != nil {
			return err
		}
	} else {
		toAcc, err = s.accountRepo.LockAccountForUpdate(tx, toID)
		if err != nil {
			return err
		}

		fromAcc, err = s.accountRepo.LockAccountForUpdate(tx, fromID)
		if err != nil {
			return err
		}
	}

	if fromAcc.Balance.LessThan(amount) {
		return errors.New("insufficient balance")
	}

	fromAcc.Balance = fromAcc.Balance.Sub(amount)
	toAcc.Balance = toAcc.Balance.Add(amount)

	err = s.accountRepo.UpdateAccount(fromAcc)
	if err != nil {
		return err
	}

	err = s.accountRepo.UpdateAccount(toAcc)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO transactions
		(from_account_id, to_account_id, amount, currency, type, status)
		VALUES (?, ?, ?, 'BDT', 'transfer', 'completed')
	`, fromID, toID, amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}
