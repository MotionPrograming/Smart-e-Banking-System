package user

import "smart-e-banking/backend/domain"

func (svc *service) CreateUser(user domain.User) (*domain.User, error) {
	return svc.usrRepo.CreateUser(user)
}
