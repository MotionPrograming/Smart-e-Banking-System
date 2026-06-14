package user

import (
	"smart-e-banking/backend/domain"
)

func (r *userRepository) CreateUser(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (name, email, password_hash, phone, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Phone,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}
