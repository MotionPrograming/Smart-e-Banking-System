package user

import (
	"smart-e-banking/backend/domain"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateUser(user domain.User) (*domain.User, error)
	GetUserByID(id int64) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	EmailExists(email string) (bool, error)
}

type userRepository struct {
	db *sqlx.DB
}

var _ Repository = (*userRepository)(nil)

func NewUserRepository(db *sqlx.DB) Repository {
	return &userRepository{
		db: db,
	}
}
