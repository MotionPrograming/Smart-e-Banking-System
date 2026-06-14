package user

import (
	"fmt"
	"smart-e-banking/backend/domain"
)

func (r *userRepository) CreateUser(user domain.User) (*domain.User, error) {
	query := `
	INSERT INTO users (
		name,
		email,
		password,
		phone,
		role,
		created_at,
		updated_at
	)
	VALUES (
		:name,
		:email,
		:password,
		:phone,
		:role,
		:created_at,
		:updated_at
	)
	RETURNING id
	`

	var userID int64

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
	}

	user.ID = userID
	return &user, nil
}
