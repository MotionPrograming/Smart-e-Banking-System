package user

import (
	userRepo "smart-e-banking/backend/repository/user"
)

type service struct {
	usrRepo userRepo.Repository
}

var _ Service = (*service)(nil)

func NewService(usrRepo userRepo.Repository) Service {
	return &service{usrRepo: usrRepo}
}
