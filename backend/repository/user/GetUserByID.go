package user

import (
	"smart-e-banking/backend/domain"
)

func (r *userRepository) GetUserByID(id int64) (*domain.User, error) {
	query := `
		SELECT id, name, email, password_hash, phone, role, created_at, updated_at
		FROM users
		WHERE id = ?
		LIMIT 1
	`
	var u domain.User
	if err := r.db.Get(&u, query, id); err != nil {
		return nil, err
	}
	return &u, nil
}
