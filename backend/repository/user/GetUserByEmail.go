package user

import (
	"smart-e-banking/backend/domain"
)

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, name, email, password_hash, phone, role, created_at, updated_at
		FROM users
		WHERE email = ?
		LIMIT 1
	`
	var u domain.User
	if err := r.db.Get(&u, query, email); err != nil {
		return nil, err
	}
	return &u, nil
}
