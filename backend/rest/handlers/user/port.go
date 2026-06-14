package user

import "smart-e-banking/backend/domain"

type Service interface {
	CreateUser(user domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id int64) (*domain.User, error)
	EmailExists(email string) (bool, error)
	Authenticate(email, password string) (*domain.User, error)
}
