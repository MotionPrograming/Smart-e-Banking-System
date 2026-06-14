package account

import "smart-e-banking/backend/domain"

func (s *service) GetAccountsByUserID(userID int64) ([]domain.Account, error) {
	return s.accountRepo.GetAccountsByUserID(userID)
}
