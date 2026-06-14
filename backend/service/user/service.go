package user

import "smart-e-banking/backend/domain"

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

type UserRepo interface {
	CreateUser(user domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id int64) (*domain.User, error)
	EmailExists(email string) (bool, error)
}
