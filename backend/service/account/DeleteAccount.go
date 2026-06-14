package account

import "errors"

func (s *service) DeleteAccount(id int64) error {
	acc, err := s.accountRepo.GetAccountByID(id)
	if err != nil {
		return err
	}
	if !acc.Balance.IsZero() {
		return errors.New("cannot delete account with a non-zero balance")
	}
	return s.accountRepo.DeleteAccount(id)
}
