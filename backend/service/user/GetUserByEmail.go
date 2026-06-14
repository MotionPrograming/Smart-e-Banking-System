package user

import "smart-e-banking/backend/domain"

func (svc *service) GetUserByEmail(email string) (*domain.User, error) {
	return svc.usrRepo.GetUserByEmail(email)
}
