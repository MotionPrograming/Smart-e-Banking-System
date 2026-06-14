package account

import (
	"errors"
	"smart-e-banking/backend/domain"
)

func (s *service) UpdateAccount(acc *domain.Account) error {
	if acc == nil {
		return errors.New("account cannot be nil")
	}
	return s.accountRepo.UpdateAccount(acc)
}
