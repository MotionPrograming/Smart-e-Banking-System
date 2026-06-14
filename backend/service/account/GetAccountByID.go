package account

import "smart-e-banking/backend/domain"

func (s *service) GetAccountByID(id int64) (*domain.Account, error) {
	return s.accountRepo.GetAccountByID(id)
}
